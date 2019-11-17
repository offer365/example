package io

import (
	"bufio"
	"compress/gzip"
	"errors"
	"io"
	"io/ioutil"
	"os"
)

func ReadWithLength(name string, length int64) (str string, err error) {
	// 只读的方式打开
	file, err := os.Open(name)
	//  os.open 底层是调用的 OpenFile ,OpenFile(文件名，打开方式，权限)
	// file,err=os.OpenFile(filepath,os.O_RDONLY,0666)
	if err = check(err); err != nil {
		return
	}
	defer file.Close()
	buf := make([]byte, length)
	// 读取文件内容返回 读取的长度 和错误
	n, err := file.Read(buf)
	str = string(buf[0:n])
	return
}

func check(err error) error {
	switch {
	case err == nil:
		return nil
	case err == io.EOF:
		return err
	case os.IsNotExist(err):
		return errors.New("is not exist: " + err.Error())
	case os.IsExist(err):
		return errors.New("is already exist: " + err.Error())
	case os.IsTimeout(err):
		return errors.New("is timeout: " + err.Error())
	case os.IsPermission(err):
		return errors.New("is permission: " + err.Error())
	default:
		return errors.New("unknown error")
	}
}

// 循环读取指定的字节数
func LoopReadWithLength(name string, length int64) (str string, err error) {
	file, err := os.Open(name)
	if err = check(err); err != nil {
		return
	}
	defer file.Close()
	var content []byte
	buf := make([]byte, length)
	for {
		n, err := file.Read(buf)
		if err == io.EOF {
			break
		}
		if err = check(err); err != nil {
			return string(content), err
		}
		// 这个地方一定要记得... 展开  在一个切片中追加另一个切边 用...
		content = append(content, buf[:n]...)
	}
	str = string(content)
	return
}

// 使用bufio读取文件 一行一行读取文件
func ReadWithLine(name string) (lines []string, err error) {
	file, err := os.Open(name)
	if err = check(err); err != nil {
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		// 注意这里是字符 不是字符串 所以不能用双引号，要用单引号
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			return lines, err
		}
		lines = append(lines, line)
	}
}

// 使用 ioutil 方便的读取文件
func ReadUtil(name string) (str string, err error) {
	content, err := ioutil.ReadFile(name)
	if err = check(err); err != nil {
		return
	}
	str = string(content)
	return
}

// 读取压缩文件
func ReadCompressWithLength(name string, lenght int64) (str string, err error) {
	file, err := os.Open(name)
	if err = check(err); err != nil {
		return
	}
	gf, err := gzip.NewReader(file)
	if err = check(err); err != nil {
		return
	}
	var content []byte
	buf := make([]byte, lenght)
	for {
		n, err := gf.Read(buf)
		if err == io.EOF {
			break
		}
		if err = check(err); err != nil {
			return string(content), err
		}
		content = append(content, buf[:n]...)
	}
	str = string(content)
	return
}
