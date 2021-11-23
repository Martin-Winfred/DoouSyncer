package monitor

import (
	"errors"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	net2 "github.com/shirou/gopsutil/net"
	"os"
	"runtime"
	"time"
)

//Get Host Info
//Get bosic Info Of host"github.com/shirou/gopsutil/mem"
func GetHostInfo() (sysInfo, archInfo, name string, numCpu int, err error) {
	sysInfo = runtime.GOOS
	archInfo = runtime.GOARCH
	numCpu = runtime.NumCPU()
	name, err = os.Hostname()
	if err != nil {
		err = errors.New("can't detect HostInfo")
		return
	}
	return
}

func GetSYSInfo() (kernelVersion, version, platform, family string, err error) {
	kernelVersion, err = host.KernelVersion()
	if err != nil {
		err = errors.New("can't load kernel version")
		return
	}
	platform, family, version, err = host.PlatformInformation()
	if err != nil {
		err = errors.New("can't load platform version")
		return
	}
	return
}

// Get resource usage
func GetCpuPrefect() (cpuload float64, err error) {
	percent, err := cpu.Percent(time.Second, false)
	if err != nil {
		err = errors.New("can't get cpu load")
		return
	}
	cpuload = percent[0]
	return
}

func GetMemPercent() (memUsage float64, err error) {
	memInfo, err := mem.VirtualMemory()
	if err != nil {
		err = errors.New("can't get Memory usage")
		return
	}
	memUsage = memInfo.UsedPercent
	return
}

func GetDiskPercent() (diskInfo float64, err error) {
	parts, _ := disk.Partitions(true)
	if err != nil {
		err = errors.New("can't get disk Partitions")
		return
	}
	diskInf, _ := disk.Usage(parts[0].Mountpoint)
	if err != nil {
		err = errors.New("can't get disk info")
		return
	}
	diskInfo = diskInf.UsedPercent
	return
}


/* Unused function for get realtime bandwidth
func GetNetInfo(InterfaceName string) (name string, bytesRecv, bytesSend uint64, err error) {
	netCard, err := net2.IOCounters(true)
	var coun int
	if err != nil {
		errors.New("can't get Net speed")
		return
	} else {
		for i := 0; netCard[i].Name != InterfaceName; i++ {
			coun = i+1
		}
	}
	name = netCard[coun].Name
	bytesRecv = netCard[coun].BytesRecv
	bytesSend = netCard[coun].BytesSent

	for _, n := range netCard{
		if n.Name == InterfaceName {
			name = netCard[coun].Name
			bytesRecv = netCard[coun].BytesRecv
			bytesSend = netCard[coun].BytesSent
			break
		}
	}
	return
}
*/

func GetNetInfo(InterfaceName string) (name string, bytesRecv, bytesSend uint64, err error) {
	netCard, err := net2.IOCounters(true)
	var coun int
	if err != nil {
		errors.New("can't get Net speed")
		return
	} else {
		for i := 0; netCard[i].Name != InterfaceName; i++ {
			coun = i+1
		}
	}
	name = netCard[coun].Name
	bytesRecv = netCard[coun].BytesRecv
	bytesSend = netCard[coun].BytesSent
	return
}