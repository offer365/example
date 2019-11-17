package linux

import (
	"io/ioutil"
	"strings"

	"github.com/offer365/example/hardware/hd"
)

type Linux struct {}

func (l *Linux) Cpu() *hd.Cpu {
	return &hd.Cpu{}
}

func (l *Linux) Mem() *hd.Memory {
	panic("implement me")
}

func (l *Linux) Network() []hd.Network {
	panic("implement me")
}

func (l *Linux) OS() *hd.Os {
	panic("implement me")
}

func (l *Linux) Bios() *hd.Bios {
	return &hd.Bios{
		Vendor:  cat("/sys/class/dmi/id/bios_vendor"),
		Version: cat("/sys/class/dmi/id/bios_version"),
	}
}

func (l *Linux) Board() *hd.Board {
	return &hd.Board{
		Name:    cat("/sys/class/dmi/id/board_name"),
		Version: cat("/sys/class/dmi/id/board_version"),
		Vendor:  cat("/sys/class/dmi/id/board_vendor"),
		Serial:  cat("/sys/class/dmi/id/board_serial"),
	}
}

func (l *Linux) Product() *hd.Product {
	return &hd.Product{
		Model:   cat("/sys/class/dmi/id/product_name"),
		Version: cat("/sys/class/dmi/id/product_version"),
		Vendor:  cat("/sys/class/dmi/id/sys_vendor"),
		Uuid:    cat("/sys/class/dmi/id/product_serial"),
	}
}

func cat(path string) string  {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return ""
	}

	return strings.TrimSpace(string(data))
}

