// Copyright © 2016 Zlatko Čalušić
//
// Use of this source code is governed by an MIT-style license that can be found in the LICENSE file.

package winsysinfo

import (
	"github.com/StackExchange/wmi"
)

// Chassis information.
type Chassis struct {
	Type     uint   `json:"type,omitempty"`
	Vendor   string `json:"vendor,omitempty"`
	Version  string `json:"version,omitempty"`
	Serial   string `json:"serial,omitempty"`
	AssetTag string `json:"assettag,omitempty"`
}

type _chassis struct {
	ChassisTypes   uint16
	Manufacturer   string
	SerialNumber   string
	Version        string
	SMBIOSAssetTag string
}

func (si *SysInfo) getChassisInfo() {
	// Win32_baseboard 主板 参数说明  https://blog.csdn.net/yeyingss/article/details/49357639
	var dst []_chassis
	query := `SELECT  ChassisTypes,Manufacturer,SerialNumber,Version,SMBIOSAssetTag  FROM Win32_SystemEnclosure WHERE (Name IS NOT NULL)`
	err := wmi.Query(query, &dst)
	if err != nil {
		return
	}
	if len(dst) > 0 {
		si.Chassis.Type = 3 // dst[0].ChassisTypes
		si.Chassis.Version = dst[0].Version
		si.Chassis.Vendor = dst[0].Manufacturer
		si.Chassis.Serial = dst[0].SerialNumber
		si.Chassis.AssetTag = dst[0].SMBIOSAssetTag
	}
	return
}
