package hardware

import (
	"runtime"

	"github.com/offer365/example/hardware/hd"
	"github.com/offer365/example/hardware/linux"
	"github.com/offer365/example/hardware/win"
)

type Hardware interface {
	Cpu() *hd.Cpu
	Mem() *hd.Memory
	Network() []hd.Network
	OS() *hd.OS
	Bios() *hd.Bios
	Board() *hd.Board
	Product() *hd.Product
}

func GetHardware() Hardware {
	switch runtime.GOOS {
	case "windows":
		return &win.Windows{}
	case "linux":
		return &linux.Linux{}
	default:
		return nil
	}
}


