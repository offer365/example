package io

import (
	"io"
	"os"
)

// func filecopy(dstname,srcname string) (int64,error)  {	// 如果形参的类型一样，可以这样简写
func Copy(dst string, src string) (written int64, err error) {
	srcF, err := os.Open(src)
	defer srcF.Close()
	if err = check(err); err != nil {
		return
	}

	// dstF,err:=os.OpenFile(dst,os.O_CREATE|os.O_WRONLY|os.O_TRUNC,0644)
	dstF, err := os.Create(dst) // Create 调用的是 OpenFile(name, O_RDWR|O_CREATE|O_TRUNC, 0666)
	// OpenFile(name, O_RDWR|O_CREATE|O_TRUNC, 0666)
	defer dstF.Close()
	if err = check(err); err != nil {
		return
	}
	// io.copy 返回两个值 一个是 copy 的长度 ，一个是err
	written, err = io.Copy(dstF, srcF)
	//  复制之后，文件内容并没有正在保存在文件中，需要使用sync 同步之后才能写入的硬盘。
	err = dstF.Sync()
	return
}
