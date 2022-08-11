package main

import (
	"dpCall/dp"
	"dpCall/probe"
	"dpCall/share"
	"dpCall/share/cluster"
	"dpCall/share/container"
	"dpCall/share/fsmon"
	"dpCall/share/global"
	"dpCall/share/utils"
	"dpCall/workerlet"
	"flag"
	log "github.com/sirupsen/logrus"
	"os"
	"sync/atomic"
)

const goroutineStackSize = 1024 * 1024

var containerTaskExitChan chan interface{} = make(chan interface{}, 1)
var errRestartChan chan interface{} = make(chan interface{}, 1)
var restartChan chan interface{} = make(chan interface{}, 1)
var monitorExitChan chan interface{} = make(chan interface{}, 1)

var Host share.CLUSHost = share.CLUSHost{
	Platform: share.PlatformDocker,
	Network:  share.NetworkDefault,
}
var Agent, parentAgent share.CLUSAgent
var agentEnv AgentEnvInfo

var evqueue cluster.ObjectQueueInterface
var messenger cluster.MessengerInterface
var agentTimerWheel *utils.TimerWheel
var prober *probe.Probe
var grpcServer *cluster.GRPCServer

//var scanUtil *scanUtils.ScanUtil
var fileWatcher *fsmon.FileWatch

var connLog *log.Logger = log.New()
var nvSvcPort, nvSvcBrPort string
var driver string
var exitingFlag int32
var exitingTaskFlag int32

var walkerTask *workerlet.Tasker

func shouldExit() bool {
	return (atomic.LoadInt32(&exitingFlag) != 0)
}

func isAgentContainer(id string) bool {
	return id == Agent.ID || id == parentAgent.ID
}

func getHostIPs() {
	addrs := getHostAddrs()
	Host.Ifaces, gInfo.hostIPs, gInfo.jumboFrameMTU = parseHostAddrs(addrs, Host.Platform, Host.Network)
	if tun := global.ORCH.GetHostTunnelIP(addrs); tun != nil {
		Host.TunnelIP = tun
	}

	if global.ORCH.ConsiderHostsAsInternal() {
		addHostSubnets(Host.Ifaces, gInfo.localSubnetMap)
	}
	mergeLocalSubnets(gInfo.internalSubnets)
}

// Sort existing containers, move containers share network ns to other containers to the front.
// Only need to consider containers in the set, not those already exist.
func sortContainerByNetMode(ids utils.Set) []*container.ContainerMetaExtra {
	sorted := make([]*container.ContainerMetaExtra, 0, ids.Cardinality())
	for id := range ids.Iter() {
		if info, err := global.RT.GetContainer(id.(string)); err == nil {
			sorted = append(sorted, info)
		}
	}

	return container.SortContainers(sorted)
}

// Sort existing containers, move containers share network ns to other containers to the front.
// Only for Container Start from Probe channel
func sortProbeContainerByNetMode(starts utils.Set) []*container.ContainerMetaExtra {
	sorted := make([]*container.ContainerMetaExtra, 0, starts.Cardinality())
	for start := range starts.Iter() {
		s := start.(*share.ProbeContainerStart)
		if info, err := global.RT.GetContainer(s.Id); err == nil {
			if info.Running && info.Pid == 0 { // cri-o: fault-tolerent for http channel errors
				info.Pid = s.RootPid_alt
				log.WithFields(log.Fields{"id": s.Id, "rootPid": info.Pid}).Debug("PROC: Update")
			}
			sorted = append(sorted, info)
		}
	}

	return container.SortContainers(sorted)
}

func adjustContainerPod(selfID string, containers []*container.ContainerMeta) string {
	for _, c := range containers {
		if v, ok := c.Labels["io.kubernetes.sandbox.id"]; ok {
			if v == selfID {
				log.WithFields(log.Fields{"Pod": selfID, "ID": c.ID}).Debug("Update")
				return c.ID
			}
		}

		if c.Sandbox != c.ID { // a child
			if c.Sandbox == selfID {
				log.WithFields(log.Fields{"Pod": selfID, "ID": c.ID}).Debug("Update ")
				return c.ID
			}
		}
	}
	return selfID
}

func main() {
	// 初始化日志
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
	log.SetFormatter(&utils.LogFormatter{Module: "AGT"})

	connLog.Out = os.Stdout
	connLog.Level = log.InfoLevel
	connLog.Formatter = &utils.LogFormatter{Module: "AGT"}

	log.WithFields(log.Fields{"version": "1.0"}).Info("START")

	rtSock := flag.String("u", "", "Container socket URL")
	skip_nvProtect := flag.Bool("s", false, "Skip NV Protect")

	// 获取编排类型，network,容器类型
	platform, flavor, network, containers, err := global.SetGlobalObjects(*rtSock, nil)
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error("Failed to initialize")
		os.Exit(-2)
	}
	// 获取容器编排引擎的版本信息
	log.WithFields(log.Fields{"endpoint": *rtSock, "runtime": global.RT.String()}).Info("Container socket connected")
	if platform == share.PlatformKubernetes {
		k8sVer, ocVer := global.ORCH.GetVersion(false, false)
		log.WithFields(log.Fields{"k8s": k8sVer, "oc": ocVer}).Info()
	}
	log.WithFields(log.Fields{"flavor": flavor, "network": network, "containers": containers}).Info()

	//容器ID
	var selfID string
	// 判断该进程是否运行在容器中
	agentEnv.runInContainer = global.SYS.IsRunningInContainer()
	if agentEnv.runInContainer {
		selfID, agentEnv.containerInContainer, err = global.SYS.GetSelfContainerID()
		if selfID == "" { // it is a POD ID in the k8s cgroup v2; otherwise, a real container ID
			log.WithFields(log.Fields{"error": err}).Error("Unsupported system. Exit!")
			os.Exit(-2)
		}
		// 将自身容器设置为保护模式
		agentEnv.containerShieldMode = (!*skip_nvProtect)
		log.WithFields(log.Fields{"shield": agentEnv.containerShieldMode}).Info("PROC:")
	} else {
		log.Info("Not running in container.")
	}
	//log.Info("selfID is " + selfID)
	//if platform == share.PlatformKubernetes {
	//	selfID = adjustContainerPod(selfID, containers)
	//}

	Host.Platform = platform
	Host.Flavor = flavor
	Host.Network = network
	Agent.Domain = global.ORCH.GetDomain(Agent.Labels)
	parentAgent.Domain = global.ORCH.GetDomain(parentAgent.Labels)

	//policyInit()

	dpStatusChan := make(chan bool, 2)
	dp.Open(dpTaskCallback, dpStatusChan, errRestartChan)
	//维持线程，测试代码
	//for {
	//	time.Sleep(time.Second)
	//}
	// Probe
	probeTaskChan := make(chan *probe.ProbeMessage, 256) // increase to avoid underflow
	fsmonTaskChan := make(chan *fsmon.MonitorMessage, 8)
	// 启动container监控线程
	//eventMonitorLoop(probeTaskChan, fsmonTaskChan, dpStatusChan)
}
