package win

import (
	"fmt"
	"runtime"
	"syscall"

	"github.com/StackExchange/wmi"
	"github.com/offer365/example/hardware/hd"
)



var (
	_advapi = syscall.NewLazyDLL("Advapi32.dll")
	_kernel = syscall.NewLazyDLL("Kernel32.dll")
)


/*
开启 wmi 服务： https://jingyan.baidu.com/article/fd8044faebc6a85030137a4e.html
wql语法： https://www.cnblogs.com/railgunman/articles/1810485.html
wmi 使用：   https://www.jianshu.com/p/57e10b3313fb   https://www.jb51.net/article/56231.htm
powershell使用： Get-WmiObject -Class Win32_Product
WMI使用的WIN32_类库名： https://blog.csdn.net/liuxingbin/article/details/6790124
*/


type _host struct {

}

type _bios struct {
	Name              string
	Manufacturer      string
	SerialNumber      string
	SMBIOSBIOSVersion string
	Version           string
}

type _board struct {
	Manufacturer string
	SerialNumber string
	Product      string
	Version      string
}

type _product struct {
	Name string
	Version string
	Vendor string
	UUID string
}

type _cpu struct {
	Manufacturer string
	Name string
	MaxClockSpeed int
	ProcessorId string
}

type _mem struct {
	Manufacturer string
	Speed uint32
	Capacity uint64
}

type _network struct {
	Name string
	ServiceName string
	MACAddress string
	Speed string
}

type _os struct {
	Caption string
	Organization string
	BuildNumber string
	SerialNumber string
	Version string
	OSArchitecture string
}

type Windows struct {
}

func (w *Windows) Cpu() *hd.Cpu {
	// Win32_Processor CPU 参数说明:  https://blog.csdn.net/yeyingss/article/details/49385421
	var dst []_cpu
	query := `SELECT Manufacturer,Name,MaxClockSpeed,ProcessorId FROM Win32_Processor WHERE (Name IS NOT NULL)`
	err := wmi.Query(query, &dst) // WHERE (BIOSVersion IS NOT NULL)
	if err != nil {
		return nil
	}
	if len(dst) > 0 {
		return &hd.Cpu{
			Vendor:  dst[0].Manufacturer,
			Model: dst[0].Name,
			Speed:dst[0].MaxClockSpeed,
			Cores:runtime.NumCPU(),
		}
	}

	return nil
}

func (w *Windows) Mem() *hd.Memory {
	// Win32_PhysicalMemory 参数说明 https://blog.csdn.net/yeyingss/article/details/49357607
	var dst []_mem
	query := `SELECT Manufacturer,Speed,Capacity FROM Win32_PhysicalMemory WHERE (Name IS NOT NULL)`
	err := wmi.Query(query, &dst) // WHERE (BIOSVersion IS NOT NULL)
	if err != nil {
		return nil
	}

	if len(dst) > 0 {
		var size uint64
		for _,m :=range dst {
			size+=m.Capacity
		}
		return &hd.Memory{
			Model:  dst[0].Manufacturer,
			Speed:  dst[0].Speed,
			Size:size,
		}
	}

	return nil
}

func (w *Windows) Network() []hd.Network {
	var dst []_network
	query := `SELECT Name,ServiceName,MACAddress,Speed FROM Win32_NetworkAdapter WHERE (Speed IS NOT NULL) AND (Speed < 9223372036854775807) AND (NOT (PNPDeviceID LIKE 'ROOT%'))`
	err := wmi.Query(query, &dst) // WHERE (BIOSVersion IS NOT NULL)
	if err != nil {
		return nil
	}

	var nts []hd.Network

	for _,m :=range dst {
		nt:=hd.Network{
			Name:   m.ServiceName,
			Driver: m.Name,
			Mac:    m.MACAddress,
			Speed:  m.Speed,
		}
		nts = append(nts, nt)
	}
	return nts
}

func (w *Windows) OS() *hd.OS {
	var dst []_os
	query:=`SELECT  Caption,Organization,BuildNumber,SerialNumber,OSArchitecture,Version  FROM Win32_OperatingSystem WHERE (Name IS NOT NULL)`
	err := wmi.Query(query, &dst)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	if len(dst) > 0 {
		return &hd.OS{
			Name: dst[0].Caption,
			Version: dst[0].Version,
			Vendor:  dst[0].Organization,
			Serial:  dst[0].SerialNumber,
			Release:dst[0].BuildNumber,
			Arch:dst[0].OSArchitecture,
		}
	}

	return nil
}

func (w *Windows) Bios() *hd.Bios {
	// Win32_BIOS 参数的说明 https://blog.csdn.net/yeyingss/article/details/49383807
	var dst []_bios
	query := `SELECT Name,Manufacturer,SerialNumber,Version,SMBIOSBIOSVersion FROM Win32_BIOS WHERE (Name IS NOT NULL)`
	err := wmi.Query(query, &dst) // WHERE (BIOSVersion IS NOT NULL)
	if err != nil {
		return nil
	}
	if len(dst) > 0 {
		return &hd.Bios{
			Vendor:  dst[0].Manufacturer,
			Version: dst[0].SMBIOSBIOSVersion,
		}
	}

	return nil
}

func (w *Windows) Board() *hd.Board {
	// Win32_baseboard 主板 参数说明  https://blog.csdn.net/yeyingss/article/details/49357639
	var dst []_board
	query:=`SELECT  Version,Product,SerialNumber,Manufacturer  FROM Win32_BaseBoard WHERE (Product IS NOT NULL)`
	err := wmi.Query(query, &dst)
	if err != nil {
		return nil
	}
	if len(dst) > 0 {
		return &hd.Board{
			Name: dst[0].Product,
			Version: dst[0].Version,
			Vendor:  dst[0].Manufacturer,
			Serial:  dst[0].SerialNumber,
		}
	}
	return nil
}

func (w *Windows) Product() *hd.Product {
	var dst []_product
	query:=`SELECT  Version,Name,Vendor,UUID  FROM Win32_ComputerSystemProduct WHERE (UUID IS NOT NULL)`
	err := wmi.Query(query, &dst)
	if err != nil {
		return nil
	}
	if len(dst) > 0 {
		return &hd.Product{
			Model: dst[0].Name,
			Version: dst[0].Version,
			Vendor:  dst[0].Vendor,
			Uuid:dst[0].UUID,
		}
	}

	return nil
}


