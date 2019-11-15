package hardware

import "testing"

func TestWindows(t *testing.T) {
	win:=&windows{}
	win.Kernel()
	win.Bios()
	win.Board()
	win.Product()
	win.Cpu()
}
