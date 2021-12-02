package netSpeed

import (
	"github.com/Martin-Winfred/DoouSyncer/pkg/monitor"
	"time"
)

func Speed(InterfaceName string) (recvSpeed, sendSpeed uint64, err error) {
	for {
		_, RecvOld, SendOld, _ := monitor.GetNetInfo(InterfaceName)
		time.Sleep(time.Second)
		_, Recv, Send, _ := monitor.GetNetInfo(InterfaceName)

		recvSpeed = (Recv - RecvOld)
		sendSpeed = (Send - SendOld)
		return
	}
	return
}
