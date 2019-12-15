// Copyright © 2016 Zlatko Čalušić
//
// Use of this source code is governed by an MIT-style license that can be found in the LICENSE file.

// Package sysinfo is a pure Go library providing Linux OS / kernel / hardware system information.
package winsysinfo

/*
开启 wmi 服务： https://jingyan.baidu.com/article/fd8044faebc6a85030137a4e.html
wql语法： https://www.cnblogs.com/railgunman/articles/1810485.html
wmi 使用：   https://www.jianshu.com/p/57e10b3313fb   https://www.jb51.net/article/56231.htm
wql测试工具 http://blog.sina.com.cn/s/blog_189b318270102y77o.html
powershell使用： Get-WmiObject -Class Win32_Product
WMI使用的WIN32_类库名： https://blog.csdn.net/liuxingbin/article/details/6790124
*/

// SysInfo struct encapsulates all other information structs.
type SysInfo struct {
	Meta    Meta            `json:"sysinfo"`
	OS      OS              `json:"os"`
	Kernel  Kernel          `json:"kernel"`
	Product Product         `json:"product"`
	Board   Board           `json:"board"`
	Node    Node            `json:"node"`
	Chassis Chassis         `json:"chassis"`
	BIOS    BIOS            `json:"bios"`
	CPU     CPU             `json:"cpu"`
	Memory  Memory          `json:"memory"`
	Network []NetworkDevice `json:"network,omitempty"`
}

// GetSysInfo gathers all available system information.
func (si *SysInfo) GetSysInfo() {
	// Meta info
	si.getMetaInfo()

	// DMI info
	si.getProductInfo()
	si.getBoardInfo()
	si.getChassisInfo()
	si.getBIOSInfo()

	// SMBIOS info
	si.getMemoryInfo()

	// Hardware info
	si.getCPUInfo() // depends on Node info

	si.getNetworkInfo()

	// Software info
	si.getOSInfo()
	si.getKernelInfo()
	// Node info
	si.getNodeInfo()
}
