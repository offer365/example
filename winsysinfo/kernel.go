// Copyright © 2016 Zlatko Čalušić
//
// Use of this source code is governed by an MIT-style license that can be found in the LICENSE file.

package winsysinfo

// Kernel information.
type Kernel struct {
	Release      string `json:"release,omitempty"`
	Version      string `json:"version,omitempty"`
	Architecture string `json:"architecture,omitempty"`
}

func (si *SysInfo) getKernelInfo() {
	si.Kernel.Release = si.OS.Release
	si.Kernel.Version = si.OS.Version
	si.Kernel.Architecture = si.OS.Architecture
}
