// Copyright © 2016 Zlatko Čalušić
//
// Use of this source code is governed by an MIT-style license that can be found in the LICENSE file.

package winsysinfo

import (
	"strconv"

	"github.com/StackExchange/wmi"
)

// Memory information.
type Memory struct {
	Type   string `json:"type,omitempty"`
	Speed  uint   `json:"speed,omitempty"` // RAM data rate in MT/s
	Size   uint   `json:"size,omitempty"`  // RAM size in MB
	Vendor string `json:"vendor,omitempty"`
}

type _mem struct {
	MemoryType   int
	Manufacturer string
	Speed        uint
	Capacity     uint
}

func (si *SysInfo) getMemoryInfo() {
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
		si.Memory.Type = strconv.Itoa(dst[0].MemoryType)
		si.Memory.Vendor = dst[0].Manufacturer
		si.Memory.Speed = dst[0].Speed
		si.Memory.Size = size
	}

	return
}
