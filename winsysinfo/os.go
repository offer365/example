// Copyright © 2016 Zlatko Čalušić
//
// Use of this source code is governed by an MIT-style license that can be found in the LICENSE file.

package winsysinfo

import (
	"github.com/StackExchange/wmi"
)

// OS information.
type OS struct {
	Name         string `json:"name,omitempty"`
	Vendor       string `json:"vendor,omitempty"`
	Version      string `json:"version,omitempty"`
	Release      string `json:"release,omitempty"`
	Architecture string `json:"architecture,omitempty"`
	Serial       string `json:"serial,omitempty"`
}

type _os struct {
	Name           string
	Manufacturer   string
	BuildNumber    string
	SerialNumber   string
	Version        string
	OSArchitecture string
}

func (si *SysInfo) getOSInfo() {
	var dst []_os
	query := `SELECT  Name,Manufacturer,BuildNumber,SerialNumber,OSArchitecture,Version  FROM Win32_OperatingSystem WHERE (Name IS NOT NULL)`
	err := wmi.Query(query, &dst)
	if err != nil {
		return
	}
	if len(dst) > 0 {
		si.OS.Name = dst[0].Name
		si.OS.Version = dst[0].Version
		si.OS.Vendor = dst[0].Manufacturer
		si.OS.Serial = dst[0].SerialNumber
		si.OS.Release = dst[0].BuildNumber
		si.OS.Architecture = dst[0].OSArchitecture
	}
}
