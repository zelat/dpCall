package main

import (
	"dpCall/dp"
	"dpCall/pipe"
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
	"runtime"
	"sync/atomic"
	"time"
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

var bench *Bench
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
	// 获取node上所有网卡信息
	addrs := getHostAddrs()
	log.WithFields(log.Fields{"addrs": addrs}).Info("")

	Host.Ifaces, gInfo.hostIPs, gInfo.jumboFrameMTU = parseHostAddrs(addrs, Host.Platform, Host.Network)
	// 获取node上tun设备的ip,mask
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

// 获取该节点的agent和host相关信息
func getLocalInfo(selfID string, pid2ID map[int]string) error {
	host, err := global.RT.GetHost()
	if err != nil {
		return err
	}
	Host = *host
	Host.CgroupVersion = global.SYS.GetCgroupVersion()

	//log.WithFields(log.Fields{"Host": Host}).Info("")

	getHostIPs()

	if networks, err := global.RT.ListNetworks(); err != nil {
		log.WithFields(log.Fields{"error": err}).Error("Error reading container networks")
	} else {
		gInfo.networks = networks
	}

	agentEnv.startsAt = time.Now().UTC()
	if agentEnv.runInContainer {
		dev, meta, err := global.RT.GetDevice(selfID)
		if err != nil {
			return err
		}
		Agent.CLUSDevice = *dev

		_, parent := global.RT.GetParent(meta, pid2ID)
		if parent != "" {
			dev, _, err := global.RT.GetDevice(parent)
			if err != nil {
				return err
			}
			parentAgent.CLUSDevice = *dev
			if parentAgent.PidMode == "host" {
				Agent.PidMode = "host"
			}
		}
	} else {
		Agent.ID = Host.ID
		Agent.Pid = os.Getpid()
		Agent.NetworkMode = "host"
		Agent.PidMode = "host"
		Agent.SelfHostname = Host.Name
		Agent.Ifaces = Host.Ifaces
	}
	Agent.HostName = Host.Name
	Agent.HostID = Host.ID
	Agent.Ver = Version

	agentEnv.cgroupMemory, _ = global.SYS.GetContainerCgroupPath(0, "memory")
	agentEnv.cgroupCPUAcct, _ = global.SYS.GetContainerCgroupPath(0, "cpuacct")
	return nil
}

func dumpGoroutineStack() {
	log.Info("Enforcer goroutine stack")
	buf := make([]byte, goroutineStackSize)
	bytes := runtime.Stack(buf, true)
	if bytes > 0 {
		log.Printf("%s", buf[:bytes])
	}
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

	// 设置debug
	debug := flag.Bool("d", false, "Enable control path debug")
	rtSock := flag.String("u", "", "Container socket URL")
	pipeType := flag.String("p", "", "Pipe driver")
	cnet_type := flag.String("n", "", "Container Network type")
	skip_nvProtect := flag.Bool("s", false, "Skip NV Protect")
	show_monitor_trace := flag.Bool("m", false, "Show process/file monitor traces")
	flag.Parse()

	// 全局设置debug mode
	if *debug {
		log.SetLevel(log.DebugLevel)
		gInfo.agentConfig.Debug = []string{"ctrl"}
	}
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

	// 获取node下所有容器的ID
	pid2ID := make(map[int]string)
	for _, meta := range containers {
		if meta.Pid != 0 {
			pid2ID[meta.Pid] = meta.ID
			log.WithFields(log.Fields{"ID": meta.ID}).Info("")
		}
	}

	for {
		// 获取该容器所在local host 和 dpCall进程的信息
		if err = getLocalInfo(selfID, pid2ID); err != nil {
			log.WithFields(log.Fields{"error": err}).Error("Failed to get local device information")
			os.Exit(-2)
		}

		// 等待agent所在容器的网卡同步
		if len(Agent.Ifaces) > 0 {
			break
		}
		log.Info("Wait for local interface ...")
		time.Sleep(time.Second * 4)
	}

	Host.Platform = platform
	Host.Flavor = flavor
	Host.Network = network
	Host.CapDockerBench = (global.RT.String() == container.RuntimeDocker)
	Host.CapKubeBench = global.ORCH.SupportKubeCISBench()
	Agent.Domain = global.ORCH.GetDomain(Agent.Labels)
	parentAgent.Domain = global.ORCH.GetDomain(parentAgent.Labels)

	//policyInit()
	Host.StorageDriver = global.RT.GetStorageDriver()
	log.WithFields(log.Fields{"hostIPs": gInfo.hostIPs}).Info("")
	log.WithFields(log.Fields{"host": Host}).Info("")
	log.WithFields(log.Fields{"agent": Agent}).Info("")

	//var driver string
	if *pipeType == "ovs" {
		driver = pipe.PIPE_OVS
	} else if *pipeType == "no_tc" {
		driver = pipe.PIPE_NOTC
	} else {
		driver = pipe.PIPE_TC
	}
	log.WithFields(log.Fields{"pipeType": driver, "jumboframe": gInfo.jumboFrameMTU}).Info("")
	// 增加veth pair
	if nvSvcPort, nvSvcBrPort, err = pipe.Open(driver, cnet_type, Agent.Pid, gInfo.jumboFrameMTU); err != nil {
		log.WithFields(log.Fields{"error": err}).Error("Failed to open pipe driver")
		os.Exit(-2)
	}
	//done := make(chan bool, 1)
	//c_sig := make(chan os.Signal, 1)
	//signal.Notify(c_sig, os.Interrupt, syscall.SIGTERM)

	// 读取已存在的容器
	existing, _ := global.RT.ListContainerIDs()
	if existing.Cardinality() > containerTaskChanSizeMin {
		ContainerTaskChan = make(chan *ContainerTask, existing.Cardinality())
	} else {
		ContainerTaskChan = make(chan *ContainerTask, containerTaskChanSizeMin)
	}
	// 读取容器存储类型
	rtStorageDriver = Host.StorageDriver
	log.WithFields(log.Fields{"name": rtStorageDriver}).Info("Runtime storage driver")

	// 创建一个dpStatus channel, 用于进程通信
	dpStatusChan := make(chan bool, 2)
	dp.Open(dpTaskCallback, dpStatusChan, errRestartChan)

	bPassiveContainerDetect := global.RT.String() == container.RuntimeCriO

	// Probe, 配置探针
	probeTaskChan := make(chan *probe.ProbeMessage, 256) // increase to avoid underflow
	fsmonTaskChan := make(chan *fsmon.MonitorMessage, 8)
	faEndChan := make(chan bool, 1)
	//fsmonEndChan := make(chan bool, 1)
	probeConfig := probe.ProbeConfig{
		Pid:                  Agent.Pid,
		PidMode:              Agent.PidMode,
		DpTaskCallback:       dpTaskCallback,
		NotifyTaskChan:       probeTaskChan,
		NotifyFsTaskChan:     fsmonTaskChan,
		PolicyLookupFunc:     hostPolicyLookup,
		ProcPolicyLookupFunc: processPolicyLookup,
		ReportLearnProc:      addLearnedProcess,
		ContainerInContainer: agentEnv.containerInContainer,
		GetContainerPid:      cbGetContainerPid,
		GetAllContainerList:  cbGetAllContainerList,
		//RerunKubeBench:       cbRerunKube,
		GetEstimateProcGroup: cbEstimateDeniedProcessdByGroup,
		GetServiceGroupName:  cbGetLearnedGroupName,
		FAEndChan:            faEndChan,
		DeferContStartRpt:    bPassiveContainerDetect,
		EnableTrace:          *show_monitor_trace,
		KubePlatform:         Host.Platform == share.PlatformKubernetes,
		WalkHelper:           walkerTask,
	}

	// 初始化prober
	if prober, err = probe.New(&probeConfig); err != nil {
		log.WithFields(log.Fields{"error": err}).Error("Failed to probe. Exit!")
		os.Exit(-2)
	}

	//// 文件系统监控
	//fmonConfig := fsmon.FileMonitorConfig{
	//	IsAufs:         global.RT.GetStorageDriver() == "aufs",
	//	EnableTrace:    *show_monitor_trace,
	//	EndChan:        fsmonEndChan,
	//	WalkerTask:     walkerTask,
	//	PidLookup:      prober.ProcessLookup,
	//	SendReport:     prober.SendAggregateFsMonReport,
	//	SendAccessRule: sendLearnedFileAccessRule,
	//	EstRule:        cbEstimateFileAlertByGroup,
	//}
	//
	//// 初始化fileWatcher
	//if fileWatcher, err = fsmon.NewFileWatcher(&fmonConfig); err != nil {
	//	log.WithFields(log.Fields{"error": err}).Error("Failed to open file monitor!")
	//	os.Exit(-2)
	//}
	//
	//prober.SetFileMonitor(fileWatcher)

	// 启动container监控线程
	eventMonitorLoop(probeTaskChan, fsmonTaskChan, dpStatusChan)

	clusterLoop(existing)
	existing = nil

	go statsLoop(bPassiveContainerDetect)

	// Wait for SIGTREM
	//go func() {
	//	<-c_sig
	//	done <- true
	//}()
	log.Info("Ready ...")

	var rc int
	select {
	//case <-done:
	//	rc = 0
	case <-monitorExitChan:
		rc = -2
	case <-restartChan:
		// Agent is kicked because of license limit.
		// Return -1 so that monitor will restart the agent,
		// and agent will reconnect after license update.
		rc = -1
	case <-errRestartChan:
		// Proactively restart agent to recover from error condition.
		// Return -1 so that monitor will restart the agent.
		rc = -1
		dumpGoroutineStack()
	}
	//维持线程，测试代码
	for {
		time.Sleep(time.Second)
	}

	log.Info("Exited")
	os.Exit(rc)
}
