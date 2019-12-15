package winsysinfo

type Node struct {
	Hostname   string `json:"hostname,omitempty"`
	MachineID  string `json:"machineid,omitempty"`
	Hypervisor string `json:"hypervisor,omitempty"`
	Timezone   string `json:"timezone,omitempty"`
}

func (si *SysInfo) getNodeInfo() {
	si.Node.Hostname = si.OS.HostName
	si.Node.MachineID = si.OS.Serial
	si.Node.Hypervisor = si.OS.Architecture
	si.Node.Timezone = si.OS.Timezone
}
