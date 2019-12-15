// Copyright © 2016 Zlatko Čalušić
//
// Use of this source code is governed by an MIT-style license that can be found in the LICENSE file.

package winsysinfo

import (
	"github.com/StackExchange/wmi"
)

// CPU information.
type CPU struct {
	Vendor      string `json:"vendor,omitempty"`
	Model       string `json:"model,omitempty"`
	Speed       uint   `json:"speed,omitempty"`   // CPU clock rate in MHz
	Cache       uint   `json:"cache,omitempty"`   // CPU cache size in KB
	Cpus        uint   `json:"cpus,omitempty"`    // number of physical CPUs
	Cores       uint   `json:"cores,omitempty"`   // number of physical CPU cores
	Threads     uint   `json:"threads,omitempty"` // number of logical (HT) CPU cores
	ProcessorId string `json:"processorid,omitempty"`
}

type _cpu struct {
	Manufacturer              string
	Name                      string
	MaxClockSpeed             uint
	L3CacheSize               uint
	NumberOfCores             uint
	NumberOfLogicalProcessors uint
	ProcessorId               string
}

func (si *SysInfo) getCPUInfo() {
	// Win32_Processor CPU 参数说明:  https://blog.csdn.net/yeyingss/article/details/49385421
	var dst []_cpu
	query := `SELECT Manufacturer,Name,MaxClockSpeed,ProcessorId,L3CacheSize,NumberOfCores,NumberOfLogicalProcessors FROM Win32_Processor WHERE (Name IS NOT NULL)`
	err := wmi.Query(query, &dst) // WHERE (BIOSVersion IS NOT NULL)
	if err != nil {
		return
	}
	if len(dst) > 0 {
		si.CPU.Vendor = dst[0].Manufacturer
		si.CPU.Model = dst[0].Name
		si.CPU.Speed = dst[0].MaxClockSpeed
		si.CPU.Cpus = uint(len(dst))
		si.CPU.Cache = dst[0].L3CacheSize
		si.CPU.Cores = dst[0].NumberOfCores
		si.CPU.Threads = dst[0].NumberOfLogicalProcessors
		si.CPU.ProcessorId = dst[0].ProcessorId
	}

	return
}
