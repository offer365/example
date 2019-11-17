package io

import (
	"fmt"
	"testing"
)

func TestReadFileWithLength(t *testing.T) {
	fmt.Println(ReadWithLength("test.txt", 10))
}

func TestLoopReadWithLength(t *testing.T) {
	fmt.Println(LoopReadWithLength("test.txt", 100))
}

func TestReadWithLine(t *testing.T) {
	fmt.Println(ReadWithLine("test.txt"))
}

func TestIoutil(t *testing.T) {
	fmt.Println(Ioutil("test.txt"))
}
