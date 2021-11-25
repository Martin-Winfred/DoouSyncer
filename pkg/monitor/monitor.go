package monitor

import (
	"errors"
	"github.com/shirou/gopsutil/host"
	"os"
	"runtime"
)



import (
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	net3 "github.com/shirou/gopsutil/net"
	"log"
	"net"
	"time"
)
//The Function marked with checked means it is functioning


//Get Host Info
//Get bosic Info Of host
func GetHostInfo() (sysInfo, archInfo, name string, numCpu int, bootTime uint64,err error) {
	sysInfo = runtime.GOOS
	archInfo = runtime.GOARCH
	numCpu = runtime.NumCPU()
	name, err = os.Hostname()
	if err != nil {
		err = errors.New("can't detect HostInfo")
		return
	}
	bootTime,err = host.BootTime()
	if err != nil {
		err = errors.New("can't get boot time")
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

//Checked
func GetCpuPrefect() (cpuload float64, err error) {
	percent, err := cpu.Percent(time.Second, false)
	if err != nil {
		err = errors.New("can't get cpu load")
		return
	}
	cpuload = percent[0]
	return
}

//Checked
func GetMemPercent() (memUsage float64, err error) {
	memInfo, err := mem.VirtualMemory()
	if err != nil {
		err = errors.New("can't get Memory usage")
		return
	}
	memUsage = memInfo.UsedPercent
	return
}

/*
Checked
 */
func GetDiskPercent() (diskInfo float64, err error) {
	parts, _ := disk.Partitions(false)
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


//This function only shows all the bandwith used after boot
//Checked
func GetNetInfo(InterfaceName string) (name string, bytesRecv, bytesSend uint64, err error) {
	netCard, err := net3.IOCounters(true)
	var coun int
	if err != nil {
		errors.New("can't get Net Info")
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

//Checked
func GetOutboundIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)
	//fmt.Println(localAddr.String())
	return localAddr.IP.String()
}
