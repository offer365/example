// Copyright © 2016 Zlatko Čalušić
//
// Use of this source code is governed by an MIT-style license that can be found in the LICENSE file.

package winsysinfo

import "github.com/StackExchange/wmi"

// Product information.
type Product struct {
	Name    string `json:"name,omitempty"`
	Vendor  string `json:"vendor,omitempty"`
	Version string `json:"version,omitempty"`
	Serial  string `json:"serial,omitempty"`
}

type _product struct {
	Name    string
	Version string
	Vendor  string
	UUID    string
}

func (si *SysInfo) getProductInfo() {
	var dst []_product
	query := `SELECT  Version,Name,Vendor,UUID  FROM Win32_ComputerSystemProduct WHERE (UUID IS NOT NULL)`
	err := wmi.Query(query, &dst)
	if err != nil {
		return
	}
	if len(dst) > 0 {
		si.Product.Name = dst[0].Name
		si.Product.Version = dst[0].Version
		si.Product.Vendor = dst[0].Vendor
		si.Product.Serial = dst[0].UUID

	}

	return
}
