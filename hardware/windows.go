package hardware

import (
	"fmt"
	"runtime"
	"syscall"

	"github.com/StackExchange/wmi"
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

type windows struct {
}

func (w *windows) Cpu() *cpu {
	// Win32_Processor CPU 参数说明:  https://blog.csdn.net/yeyingss/article/details/49385421
	var dst []_cpu
	query := `SELECT Manufacturer,Name,MaxClockSpeed,ProcessorId FROM Win32_Processor WHERE (Name IS NOT NULL)`
	err := wmi.Query(query, &dst) // WHERE (BIOSVersion IS NOT NULL)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	fmt.Println(dst)
	if len(dst) > 0 {
		return &cpu{
			vendor:  dst[0].Manufacturer,
			model: dst[0].Name,
			speed:dst[0].MaxClockSpeed,
			cores:runtime.NumCPU(),
		}
	}

	return nil
}

func (w *windows) Mem() *memory {
	// Win32_PhysicalMemory 参数说明 https://blog.csdn.net/yeyingss/article/details/49357607
	var dst []_mem
	query := `SELECT Manufacturer,Speed,Capacity FROM Win32_PhysicalMemory WHERE (Name IS NOT NULL)`
	err := wmi.Query(query, &dst) // WHERE (BIOSVersion IS NOT NULL)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	fmt.Println(dst)
	if len(dst) > 0 {
		var size uint64
		for _,m :=range dst {
			size+=m.Capacity
		}
		return &memory{
			model:  dst[0].Manufacturer,
			speed:  dst[0].Speed,
			size:size,
		}
	}

	return nil
}

func (w *windows) Disk() []disk {
	panic("implement me")
}

func (w *windows) Network() []network {
	panic("implement me")
}

func (w *windows) System() *sys {
	panic("implement me")
}

func (w *windows) Kernel() *kernel {
	version, err := syscall.GetVersion()
	if err != nil {
		return nil
	}
	return &kernel{
		release: "",
		version: fmt.Sprintf("%d.%d (%d)", byte(version), uint8(version>>8), version>>16),
		arch:    runtime.GOARCH,
	}
}

func (w *windows) OS() *os {
	/*
	wmic:root\cli>ComputerSystem list full

	AdminPasswordStatus=0
	AutomaticResetBootOption=TRUE
	AutomaticResetCapability=TRUE
	BootOptionOnLimit=
	BootOptionOnWatchDog=
	BootROMSupported=TRUE
	BootupState=Normal boot
	Caption=PC-20181221MOYG
	ChassisBootupState=2
	CreationClassName=Win32_ComputerSystem
	CurrentTimeZone=480
	DaylightInEffect=
	Description=AT/AT COMPATIBLE
	Domain=WorkGroup
	DomainRole=0
	EnableDaylightSavingsTime=TRUE
	FrontPanelResetStatus=2
	InfraredSupported=FALSE
	InitialLoadInfo=
	InstallDate=
	KeyboardPasswordStatus=2
	LastLoadInfo=
	Manufacturer=LENOVO
	Model=20DCA09BCD
	Name=PC-20181221MOYG
	NameFormat=
	NetworkServerModeEnabled=TRUE
	NumberOfProcessors=1
	OEMStringArray=
	PartOfDomain=FALSE
	PauseAfterReset=-1
	PowerManagementCapabilities=
	PowerManagementSupported=
	PowerOnPasswordStatus=0
	PowerState=0
	PowerSupplyState=2
	PrimaryOwnerContact=
	PrimaryOwnerName=PC
	ResetCapability=1
	ResetCount=-1
	ResetLimit=-1
	Roles={"LM_Workstation","LM_Server","Print","NT"}
	Status=OK
	SupportContactDescription=
	SystemStartupDelay=
	SystemStartupOptions=
	SystemStartupSetting=
	SystemType=x64-based PC
	ThermalState=2
	TotalPhysicalMemory=8311197696
	UserName=PC-20181221MOYG\Administrator
	WakeUpType=6
	Workgroup=WorkGroup
	*/
	panic("implement me")
}

func (w *windows) Bios() *bios {
	// Win32_BIOS 参数的说明 https://blog.csdn.net/yeyingss/article/details/49383807
	var dst []_bios
	// query:=`SELECT Name,Manufacturer,SerialNumber,Version,SMBIOSBIOSVersion FROM Win32_BIOS WHERE (Name IS NOT NULL)`
	query := `SELECT Name,Manufacturer,SerialNumber,Version,SMBIOSBIOSVersion FROM Win32_BIOS WHERE (Name IS NOT NULL)`
	err := wmi.Query(query, &dst) // WHERE (BIOSVersion IS NOT NULL)
	if err != nil {
		return nil
	}
	if len(dst) > 0 {
		return &bios{
			vendor:  dst[0].Manufacturer,
			version: dst[0].SMBIOSBIOSVersion,
		}
	}

	return nil
}

func (w *windows) Board() *board {
	// Win32_baseboard 主板 参数说明  https://blog.csdn.net/yeyingss/article/details/49357639
	var dst []_board
	query:=`SELECT  Version,Product,SerialNumber,Manufacturer  FROM Win32_BaseBoard WHERE (Product IS NOT NULL)`
	err := wmi.Query(query, &dst)
	if err != nil {
		return nil
	}
	if len(dst) > 0 {
		return &board{
			product: dst[0].Product,
			version: dst[0].Version,
			vendor:  dst[0].Manufacturer,
			serial:  dst[0].SerialNumber,
		}
	}
	return nil
}

func (w *windows) Product() *product {
	/*
	wmic:root\cli> csproduct list full

	Description=计算机系统产品
	IdentifyingNumber=PF0K6984
	Name=20DCA09BCD
	SKUNumber=
	UUID=D6C3AF01-5534-11CB-9FCD-91A2E9FAD497
	Vendor=LENOVO
	Version=ThinkPad E450
	*/
	var dst []_product
	query:=`SELECT  Version,Name,Vendor,UUID  FROM Win32_ComputerSystemProduct WHERE (UUID IS NOT NULL)`
	err := wmi.Query(query, &dst)
	if err != nil {
		return nil
	}
	if len(dst) > 0 {
		return &product{
			name: dst[0].Name,
			version: dst[0].Version,
			vendor:  dst[0].Vendor,
			serial:  dst[0].UUID,
		}
	}

	return nil
}

func (w *windows) Node() *node {
	panic("implement me")
}

