package winsysinfo

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
	"testing"

	"github.com/StackExchange/wmi"
)

func TestBios(t *testing.T) {
	// Win32_BIOS 参数的说明 https://blog.csdn.net/yeyingss/article/details/49383807
	var dst []_bios
	query := `SELECT Name,Manufacturer,SerialNumber,Version,SMBIOSBIOSVersion FROM Win32_BIOS WHERE (Name IS NOT NULL)`
	err := wmi.Query(query, &dst) // WHERE (BIOSVersion IS NOT NULL)

	if err != nil {
		return
	}
	var Date string
	if len(dst) > 0 {
		Vendor := dst[0].Manufacturer
		Version := dst[0].Version
		name := dst[0].Name
		if r, err := regexp.Compile(`\d+/\d+/\d+\s+\d+:\d+:\d+`); err == nil {
			Date = r.FindString(name)
		}
		fmt.Println(Vendor, Version, Date)
	}
}

func TestBoard(t *testing.T) {
	var dst []_board
	query := `SELECT  Version,Name,Tag,SerialNumber,Manufacturer  FROM Win32_BaseBoard WHERE (Name IS NOT NULL)`
	err := wmi.Query(query, &dst)
	fmt.Println(err)
	if err != nil {
		return
	}
	if len(dst) > 0 {
		Name := dst[0].Name
		Version := dst[0].Version
		Vendor := dst[0].Manufacturer
		Serial := dst[0].SerialNumber
		AssetTag := dst[0].Tag
		fmt.Println(Name, Version, Vendor, Serial, AssetTag)
	}
	return
}

func TestChassis(t *testing.T) {
	var dst []_chassis
	query := `SELECT  ChassisTypes,Manufacturer,SerialNumber,Version,SMBIOSAssetTag  FROM Win32_SystemEnclosure WHERE (Name IS NOT NULL)`
	err := wmi.Query(query, &dst)
	fmt.Println(err)
	if err != nil {
		return
	}
	if len(dst) > 0 {
		// Type:= dst[0].ChassisTypes
		Version := dst[0].Version
		Vendor := dst[0].Manufacturer
		Serial := dst[0].SerialNumber
		AssetTag := dst[0].SMBIOSAssetTag
		fmt.Println(Version, Vendor, Serial, AssetTag)
	}
	return
}

func TestCpu(t *testing.T) {
	// Win32_Processor CPU 参数说明:  https://blog.csdn.net/yeyingss/article/details/49385421
	var dst []_cpu
	query := `SELECT Manufacturer,Name,MaxClockSpeed,ProcessorId,L3CacheSize,NumberOfCores,NumberOfLogicalProcessors FROM Win32_Processor WHERE (Name IS NOT NULL)`
	err := wmi.Query(query, &dst) // WHERE (BIOSVersion IS NOT NULL)
	fmt.Println(err)
	if err != nil {
		return
	}
	if len(dst) > 0 {
		Vendor := dst[0].Manufacturer
		Model := dst[0].Name
		Speed := uint(dst[0].MaxClockSpeed)
		Cores := dst[0].NumberOfCores
		Cache := uint(dst[0].L3CacheSize)
		Threads := dst[0].NumberOfLogicalProcessors
		ProcessorId := dst[0].ProcessorId
		fmt.Println(Vendor, Model, Speed, Cores, Cache, Threads, ProcessorId)
	}

	return
}

func TestOS(t *testing.T) {
	var dst []_os
	query := `SELECT  Name,Manufacturer,BuildNumber,SerialNumber,OSArchitecture,Version  FROM Win32_OperatingSystem WHERE (Name IS NOT NULL)`
	err := wmi.Query(query, &dst)
	if err != nil {
		return
	}
	if len(dst) > 0 {
		Name := dst[0].Name
		Version := dst[0].Version
		Vendor := dst[0].Manufacturer
		Serial := dst[0].SerialNumber
		Release := dst[0].BuildNumber
		Architecture := dst[0].OSArchitecture
		fmt.Println(Name, Version, Vendor, Serial, Release, Architecture)
	}
}

func TestMem(t *testing.T) {
	// Win32_PhysicalMemory 参数说明 https://blog.csdn.net/yeyingss/article/details/49357607
	var dst []_mem
	query := `SELECT Manufacturer,Speed,Capacity,MemoryType FROM Win32_PhysicalMemory WHERE (Name IS NOT NULL)`
	err := wmi.Query(query, &dst) // WHERE (BIOSVersion IS NOT NULL)
	if err != nil {
		return
	}

	if len(dst) > 0 {
		var size uint
		for _, m := range dst {
			size += m.Capacity
		}

		Type := strconv.Itoa(dst[0].MemoryType)
		Vendor := dst[0].Manufacturer
		Speed := dst[0].Speed
		Size := size
		fmt.Println(Type, Vendor, Speed, Size)
	}

	return
}

func TestNetwork(t *testing.T) {
	var dst []_network
	query := `SELECT Name,Description,MACAddress,Speed,Manufacturer FROM Win32_NetworkAdapter WHERE (PhysicalAdapter = True) AND (NOT (PNPDeviceID LIKE 'ROOT%'))`
	err := wmi.Query(query, &dst) // WHERE (BIOSVersion IS NOT NULL)
	fmt.Println(err)
	if err != nil {
		return
	}

	var nts = []NetworkDevice{}
	for _, m := range dst {
		nt := NetworkDevice{
			Name:       m.Name,
			Driver:     m.Description,
			MACAddress: m.MACAddress,
			Speed:      m.Speed,
			Vendor:     m.Manufacturer,
		}
		nts = append(nts, nt)
	}
	fmt.Println(nts)
	return
}

func TestProduct(t *testing.T) {
	var dst []_product
	query := `SELECT  Version,Name,Vendor,UUID  FROM Win32_ComputerSystemProduct WHERE (UUID IS NOT NULL)`
	err := wmi.Query(query, &dst)
	if err != nil {
		return
	}
	if len(dst) > 0 {
		Name := dst[0].Name
		Version := dst[0].Version
		Vendor := dst[0].Vendor
		Serial := dst[0].UUID
		fmt.Println(Name, Version, Vendor, Serial)
	}

	return
}

func TestSysInfo_GetSysInfo(t *testing.T) {
	si := SysInfo{}
	si.GetSysInfo()
	byt, err := json.Marshal(si)
	fmt.Println(err)
	fmt.Println(string(byt))
}
