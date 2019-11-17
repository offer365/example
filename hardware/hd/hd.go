package hd

type Os struct {
	Name string
	Version string
	Vendor string
	Serial string
	Release string
	Arch string
}

type Product struct {
	Model string
	Version string
	Vendor string
	Uuid  string
}

type Board struct {
	Name string
	Version string
	Vendor string
	Serial string
}

type Bios struct {
	Vendor string
	Version string
}

type Cpu struct {
	Vendor string
	Model string
	Speed int
	Cache int
	Cpus int
	Cores int
	Threads int
}

type Memory struct {
	Model string
	Speed uint32
	Size uint64
}


type Network struct {
	Name string
	Driver string
	Mac string
	Speed string
}
