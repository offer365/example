// Copyright © 2016 Zlatko Čalušić
//
// Use of this source code is governed by an MIT-style license that can be found in the LICENSE file.

package winsysinfo

import (
	"fmt"
	"regexp"

	"github.com/StackExchange/wmi"
)

// BIOS information.
type BIOS struct {
	Vendor  string `json:"vendor,omitempty"`
	Version string `json:"version,omitempty"`
	Date    string `json:"date,omitempty"`
}

type _bios struct {
	Name              string
	Manufacturer      string
	SerialNumber      string
	Version              string
}

func (si *SysInfo) getBIOSInfo() {
	// Win32_BIOS 参数的说明 https://blog.csdn.net/yeyingss/article/details/49383807
	var dst []_bios
	query := `SELECT Name,Manufacturer,SerialNumber,Version    FROM Win32_BIOS WHERE (Name IS NOT NULL)`
	err := wmi.Query(query, &dst) // WHERE (BIOSVersion IS NOT NULL)
	fmt.Println(err)
	if err != nil {
		return
	}
	if len(dst) > 0 {
		si.BIOS.Vendor = dst[0].Manufacturer
		si.BIOS.Version = dst[0].Version
		name := dst[0].Name
		if r, err := regexp.Compile(`\d+/\d+/\d+\s+\d+:\d+:\d+`); err == nil {
			si.BIOS.Date = r.FindString(name)
		}
	}
}
