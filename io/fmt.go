package io

import (
	"fmt"
	"os"
)

var (
	a string
	b int
	c float32
)

// 从文件读取
func scanf() {
	// os.Stdin 是底层操作的是一个系统文件 这里也可以换成普通的文件
	fmt.Fscanf(os.Stdin, "%s%d%f", &a, &b, &c)
	fmt.Println(a, b, c)
}

// 从文件读取
func scan() {
	fmt.Fscan(os.Stdin, &a, &b, &c)
	fmt.Println(a, b, c)
}

// 从文件读取
func scanln() {
	fmt.Fscanln(os.Stdin, &a, &b, &c)
	fmt.Println(a, b, c)
}

// 写入到文件
func fprintln() {
	a, b, c = "hehe", 33, 33.33333
	// 以空格分割，结尾有换行符
	fmt.Fprintln(os.Stdout, a, b, c)
}

func fprintf() {
	a, b, c = "hehe", 33, 33.33333
	// 格式化输出到文件
	fmt.Fprintf(os.Stdout, "%s-%d-%f\n", a, b, c)
}

func fprint() {
	a, b, c = "hehe", 33, 33.33333
	// 没有分割符
	fmt.Fprint(os.Stdout, a)
	fmt.Fprint(os.Stdout, b)
	fmt.Fprint(os.Stdout, c)
}
