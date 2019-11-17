package io

import "os"

func termial() {
	var buf [16]byte
	os.Stdin.Read(buf[:]) // 从终端读取 相当于 fmt.Scan
	// fmt.Println(string(buf[:]))
	os.Stdout.WriteString(string(buf[:])) // 像终端写入相对于 fmt.Print
}
