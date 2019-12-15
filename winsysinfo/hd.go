package winsysinfo

type Host struct {
	MachineID    string `json:"machineid,omitempty"`
	Architecture string `json:"architecture,omitempty"`
	Hypervisor   string `json:"hypervisor,omitempty"`
}
