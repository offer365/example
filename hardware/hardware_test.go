package hardware

import "testing"

func TestWindows_Bios(t *testing.T) {
	win:=&windows{}
	win.Kernel()
	win.Bios()
	win.Board()
}
