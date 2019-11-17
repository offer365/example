package ssh

import (
	"fmt"
	"testing"
)

func TestNewCli(t *testing.T) {
	cli := NewCli("root", "666666", "10.0.0.250", "22")
	output, err := cli.Run("ls")
	fmt.Println(output, err)
	output, err = cli.Run("date")
	fmt.Println(output, err)
	cli.Close()
}
