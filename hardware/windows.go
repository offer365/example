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

type windows struct {
}

func (w *windows) Cpu() *cpu {
	panic("implement me")
}

func (w *windows) Mem() *memory {
	panic("implement me")
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
	panic("implement me")
}

func (w *windows) Bios() *bios {
	/*
		BiosCharacteristics={7,9,11,12,15,16,32,33,40,42,43}
		BuildNumber=
		CodeSet=
		CurrentLanguage=en-US
		Description=Phoenix BIOS SC-T v2.1
		IdentificationCode=
		InstallableLanguages=1
		InstallDate=
		LanguageEdition=
		ListOfLanguages={"en-US"}
		Manufacturer=LENOVO
		Name=Phoenix BIOS SC-T v2.1
		OtherTargetOS=
		PrimaryBIOS=TRUE
		ReleaseDate=20161202000000.000000+000
		SerialNumber=PF0K6984
		SMBIOSBIOSVersion=J5ET56WW (1.27 )
		SMBIOSMajorVersion=2
		SMBIOSMinorVersion=7
		SMBIOSPresent=TRUE
		SoftwareElementID=Phoenix BIOS SC-T v2.1
		SoftwareElementState=3
		Status=OK
		TargetOperatingSystem=0
		Version=LENOVO - 1270
	*/
	var dst []_bios
	// query:=`SELECT Name,Manufacturer,SerialNumber,Version,SMBIOSBIOSVersion FROM Win32_BIOS WHERE (Name IS NOT NULL)`
	query := `SELECT Name,Manufacturer,SerialNumber,Version,SMBIOSBIOSVersion FROM Win32_BIOS WHERE (Name IS NOT NULL)`
	err := wmi.Query(query, &dst) // WHERE (BIOSVersion IS NOT NULL)
	if err != nil {
		fmt.Println(err)
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
	/*
		ConfigOptions={}
		Depth=
		Description=基板
		Height=
		HostingBoard=TRUE
		HotSwappable=FALSE
		InstallDate=
		Manufacturer=LENOVO
		Model=
		Name=基板
		OtherIdentifyingInfo=
		PartNumber=
		PoweredOn=TRUE
		Product=20DCA09BCD
		Removable=FALSE
		Replaceable=TRUE
		RequirementsDescription=
		RequiresDaughterBoard=FALSE
		SerialNumber=L1HF66601PE
		SKU=
		SlotLayout=
		SpecialRequirements=
		Status=OK
		Tag=Base Board
		Version=SDK0K09938 WIN
		Weight=
		Width=
	*/
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
	panic("implement me")
}

func (w *windows) Node() *node {
	panic("implement me")
}

func (w *windows) Chassis() *chassis {
	panic("implement me")
}
