package hardware

import (
	"fmt"
	"testing"
)

func TestWindows(t *testing.T) {
	hd:=GetHardware()
	fmt.Println(hd.Cpu())
	fmt.Println(hd.Mem())
	fmt.Println(hd.Bios())
	fmt.Println(hd.Board())
	fmt.Println(hd.Product())
	fmt.Println(hd.Network())
	fmt.Println(hd.OS())
}
