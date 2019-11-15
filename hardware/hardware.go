package hardware

/*
{
  "sysinfo": {
    "version": "0.9.1",
    "timestamp": "2016-09-24T13:30:28.369498856+02:00"
  },
  "node": {
    "hostname": "web12",
    "machineid": "04aa55927ebd390829367c1757b98cac",
    "timezone": "Europe/Zagreb"
  },



}
*/
type sys struct {
	version string
	timestamp string
}

type node struct {
	hostname string
	machine string
	timezone string
}

type os struct {
	name string
	version string
	vendor string
	serial string
	release string
	arch string
}

type kernel struct {
	release string
	version string
	arch string
}

type product struct {
	name string
	version string
	vendor string
	serial string
}

type board struct {
	product string
	version string
	vendor string
	serial string
}

type chassis struct {
	model string
	vendor string
}

type bios struct {
	vendor string
	version string
}

type cpu struct {
	vendor string
	model string
	speed int
	cache int
	cpus int
	cores int
	threads int
}

type memory struct {
	model string
	speed string
	size string
}

type disk struct {
	name string
	driver  string
	vendor  string
	model  string
	serial string
	size string
}

type network struct {
	name string
	driver string
	mac string
	port string
	speed string
}

type Hardware interface {
	Cpu() *cpu
	Mem() *memory
	Disk() []disk
	Network() []network
	System() *sys
	Kernel() *kernel
	OS() *os
	Bios() *bios
	Board() *board
	Product() *product
	Node() *node
	Chassis() *chassis
}


