// Copyright © 2016 Zlatko Čalušić
//
// Use of this source code is governed by an MIT-style license that can be found in the LICENSE file.

package winsysinfo

import "github.com/StackExchange/wmi"

// Board information.
type Board struct {
	Name     string `json:"name,omitempty"`
	Vendor   string `json:"vendor,omitempty"`
	Version  string `json:"version,omitempty"`
	Serial   string `json:"serial,omitempty"`
	AssetTag string `json:"assettag,omitempty"`
}

type _board struct {
	Manufacturer string
	SerialNumber string
	Name         string
	Version      string
	Tag          string
}

func (si *SysInfo) getBoardInfo() {
	// Win32_baseboard 主板 参数说明  https://blog.csdn.net/yeyingss/article/details/49357639
	var dst []_board
	query := `SELECT  Version,Name,Tag,SerialNumber,Manufacturer  FROM Win32_BaseBoard WHERE (Name IS NOT NULL)`
	err := wmi.Query(query, &dst)
	if err != nil {
		return
	}
	if len(dst) > 0 {
		si.Board.Name = dst[0].Name
		si.Board.Version = dst[0].Version
		si.Board.Vendor = dst[0].Manufacturer
		si.Board.Serial = dst[0].SerialNumber
		si.Board.AssetTag = dst[0].Tag
	}
	return
}
