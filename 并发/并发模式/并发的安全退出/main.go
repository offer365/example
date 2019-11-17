package main

import (
	"fmt"
	"sync"
	"time"
)

// 当有多个管道均可操作时，select会随机选择一个管道。基于该特性我们可以用select实现一个生成随机数序列的程序：
func example1() {
	ch := make(chan int)
	go func() {
		for {
			select {
			case ch <- 0:
			case ch <- 1:
			}
		}
	}()

	for v := range ch {
		fmt.Println(v)
	}
}

// 可以通过close关闭一个管道来实现广播的效果，所有从关闭管道接收的操作均会收到一个零值和一个可选的失败标志。
func worker1(cannel chan bool) {
	for {
		select {
		default:
			fmt.Println("hello")
			time.Sleep(time.Millisecond * 100)
			// 正常工作
		case <-cannel:
			return
			// 退出
		}
	}
}

// 通过close来关闭cancel管道向多个Goroutine广播退出的指令。
// 不过这个程序依然不够稳健：当每个Goroutine收到退出指令退出时一般会进行一定的清理工作，
// 但是退出的清理工作并不能保证被完成，因为main线程并没有等待各个工作Goroutine退出工作完成的机制。
// 我们可以结合sync.WaitGroup来改进:

func worker(wg *sync.WaitGroup, cannel chan bool) {
	defer wg.Done()

	for {
		select {
		default:
			fmt.Println("hello")
		case <-cannel:
			return
		}
	}
}

func main() {
	cancel := make(chan bool)

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go worker(&wg, cancel)
	}

	time.Sleep(time.Second)
	close(cancel)
	wg.Wait()
}
