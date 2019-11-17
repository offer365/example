package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// 标准库的sync/atomic包对原子操作提供了丰富的支持。
var total uint64

// atomic.AddUint64函数调用保证了total的读取、更新和保存是一个原子操作，因此在多线程中访问也是安全的。
// 原子操作配合互斥锁可以实现非常高效的单件模式。
// 互斥锁的代价比普通整数的原子读写高很多，在性能敏感的地方可以增加一个数字型的标志位，通过原子检测标志位状态降低互斥锁的使用次数来提高性能。

func worker(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i <= 100; i++ {
		atomic.AddUint64(&total, 1)
	}
}

type singleton struct{}

var (
	instance *singleton
	once     sync.Once
)

// once 的应用
func Instance() *singleton {
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}

// sync/atomic包对基本的数值类型及复杂对象的读写都提供了原子操作的支持。
// atomic.Value原子对象提供了Load和Store两个原子方法，分别用于加载和保存数据，返回值和参数都是interface{}类型，因此可以用于任意的自定义复杂类型
// 这是一个简化的生产者消费者模型：后台线程生成最新的配置信息；前台多个工作者线程获取最新的配置信息。所有线程共享配置信息资源。
var config atomic.Value // 保存当前配置信息
func example() {
	var config atomic.Value // 保存当前配置信息

	// 初始化配置信息
	config.Store(loadConfig())

	// 启动一个后台线程, 加载更新后的配置信息
	go func() {
		for {
			time.Sleep(time.Second)
			config.Store(loadConfig())
		}
	}()

	// 用于处理请求的工作者线程始终采用最新的配置信息
	for i := 0; i < 10; i++ {
		go func() {
			for r := range requests() {
				c := config.Load()
				fmt.Println(r, c)
				// ...
			}
		}()
	}
}

// 伪代码
func loadConfig() []string {
	return nil
}
func requests() []string {
	return nil
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go worker(&wg)
	go worker(&wg)
	wg.Wait()

}
