package simple_factory

// 简单工厂模式
// go 语言没有构造函数一说，所以一般会定义NewXXX函数来初始化相关类。 NewXXX 函数返回接口时就是简单工厂模式，也就是说Golang的一般推荐做法就是简单工厂。
//
// 在这个simplefactory包中只有API 接口和NewAPI函数为包外可见，封装了实现细节。

import "fmt"

type API interface {
	Say(name string)
}

func NewAPI(mode int) API {
	if mode == 1 {
		return &dog{}
	} else {
		return &cat{}
	}
}

type dog struct{}

func (dog) Say(name string) {
	fmt.Println("wang!!!", name)
}

type cat struct{}

func (cat) Say(name string) {
	fmt.Println("miao!!!", name)
}
