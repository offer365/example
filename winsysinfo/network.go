// Copyright © 2016 Zlatko Čalušić
//
// Use of this source code is governed by an MIT-style license that can be found in the LICENSE file.

package winsysinfo

import (
	"github.com/StackExchange/wmi"
)

// NetworkDevice information.
type NetworkDevice struct {
	Name       string `json:"name,omitempty"`
	Driver     string `json:"driver,omitempty"`
	MACAddress string `json:"macaddress,omitempty"`
	Speed      uint   `json:"speed,omitempty"` // device max supported speed in Mbps
	Vendor     string `json:"vendor,omitempty"`
}

type _network struct {
	Name         string
	Description  string
	MACAddress   string
	Speed        uint
	Manufacturer string
}

func (si *SysInfo) getNetworkInfo() {
	var dst []_network

	// SELECT Name,Description,MACAddress,Speed,Manufacturer FROM Win32_NetworkAdapter WHERE (Speed IS NOT NULL) AND (Speed < 9223372036854775807) AND (NOT (PNPDeviceID LIKE 'ROOT%'))
	query := `SELECT Name,Description,MACAddress,Speed,Manufacturer FROM Win32_NetworkAdapter WHERE (PhysicalAdapter = True) AND (NOT (PNPDeviceID LIKE 'ROOT%'))`
	err := wmi.Query(query, &dst) // WHERE (BIOSVersion IS NOT NULL)

	if err != nil {
		return
	}

	si.Network = make([]NetworkDevice, 0)
	for _, m := range dst {
		nt := NetworkDevice{
			Name:       m.Name,
			Driver:     m.Description,
			MACAddress: m.MACAddress,
			Speed:      0, // TODO
			Vendor:     m.Manufacturer,
		}
		si.Network = append(si.Network, nt)
	}
	return
}
