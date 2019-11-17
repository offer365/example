package io

import (
	"bufio"
	"io/ioutil"
	"os"
)

func WriteWithByte(data []byte) {
	// os.O_CREATE 如果文件不存在就创建
	// os.O_TRUNC 清空文件
	// os.O_WRONLY 只写
	// os.O_APPEND 追加
	// os.O_RDONLY 只读
	// os.O_RDWR 读写
	// 0644 unix 中的文件权限
	file, err := os.OpenFile("write1.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err = check(err); err != nil {
		return
	}
	defer file.Close()
	// file.Write() 直接写入
	// file.WriteAt() 指定位置写入
	// file.WriteString()  写入一个字符串
	file.Write(data)
	file.Sync()

}

func WriteWithString(data string) {
	file, err := os.OpenFile("write1.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err = check(err); err != nil {
		return
	}
	defer file.Close()
	file.WriteString(data)
	file.Sync()
}

func WriteWithBuf(str string) {
	file, err := os.OpenFile("write2.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err = check(err); err != nil {
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	writer.WriteString(str)
	// 将缓存中内容刷入到文件中
	writer.Flush()
}

func WriteUtil(name string, data []byte) {
	err := ioutil.WriteFile(name, data, 0644)
	if err = check(err); err != nil {
		return
	}
}
