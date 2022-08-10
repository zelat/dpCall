package dp

// #include "../defs.h"
import "C"
import (
	"bytes"
	"encoding/binary"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net"
	"os"
	"unsafe"
)

func ParseDPMsgHeader(msg []byte) *C.DPMsgHdr {
	var hdr C.DPMsgHdr

	hdrLen := int(unsafe.Sizeof(hdr))
	if len(msg) < hdrLen {
		fmt.Print("Short header")
		return nil
	}

	r := bytes.NewReader(msg)
	binary.Read(r, binary.BigEndian, &hdr)
	if int(hdr.Length) != len(msg) {
		fmt.Println("kind : ", hdr.Kind, "expect : ", hdr.Length, "actual", len(msg), "Wrong message length.")
		return nil
	}

	return &hdr
}

func dpMessenger(msg []byte) {
	hdr := ParseDPMsgHeader(msg)
	if hdr == nil {
		return
	}
	//offset := int(unsafe.Sizeof(*hdr))
	//switch int(hdr.Kind) {
	//case C.DP_KIND_APP_UPDATE:
	//	dpMsgAppUpdate(msg[offset:])
	//case C.DP_KIND_THREAT_LOG:
	//	dpMsgThreatLog(msg[offset:])
	//case C.DP_KIND_CONNECTION:
	//	dpMsgConnection(msg[offset:])
	//case C.DP_KIND_FQDN_UPDATE:
	//	dpMsgFqdnIpUpdate(msg[offset:])
	//}
}

// 监听dp进程状态
func listenDP() {
	log.Debug("Listening to CTRL socket ...")
	os.Remove(ctrlServer)

	var conn *net.UnixConn
	kind := "unixgram"
	addr := net.UnixAddr{ctrlServer, kind}
	defer os.Remove(ctrlServer)
	conn, err := net.ListenUnixgram(kind, &addr)
	if err != nil {
		fmt.Println("error = %s", err.Error())
	}
	defer conn.Close()
	//读取C语言格式的struct
	for {
		var buf [C.DP_MSG_SIZE]byte
		n, err := conn.Read(buf[:])
		if err != nil {
			log.WithFields(log.Fields{"error": err}).Error("Read message error.")
		} else {
			dpAliveMsgCnt++
			dpMessenger(buf[:n])
		}
	}
}
