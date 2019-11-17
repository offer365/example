package io

import (
	"fmt"
	"os"
)

func stat() {
	fileinfo, err := os.Stat(`litao-go/文件处理/test0/test.txt`)
	if err != nil && os.IsNotExist(err) {
		fmt.Println("err")
	}
	if !fileinfo.Mode().IsRegular() {
		return
	}
	fmt.Println(fileinfo.Name())
	fmt.Println(fileinfo.Size())
	fmt.Println(fileinfo.Mode())
	fmt.Println(fileinfo.IsDir())
	fmt.Println(fileinfo.ModTime().Format("2006-01-02"))
	fmt.Println(fileinfo.ModTime().Unix())
	fmt.Printf("%T\n", fileinfo.Sys())
	fmt.Printf("%#v\n", fileinfo.Sys())
	fmt.Printf("%+v\n", fileinfo.Sys())
	// os.Rename()

}
