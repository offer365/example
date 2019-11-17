package main

import (
	"context"
	"fmt"
)

// 需要先生成最初的2, 3, 4, ...自然数序列（不包含开头的0、1）：
// 返回生成自然数序列的管道: 2, 3, 4, ...
func GenerateNatural(ctx context.Context) chan int {
	ch := make(chan int)
	go func() {
		for i := 2; ; i++ {
			select {
			case <-ctx.Done():
				return
			case ch <- i:
			}
		}
	}()
	return ch
}

// GenerateNatural函数内部启动一个Goroutine生产序列，返回对应的管道。
// 然后是为每个素数构造一个筛子：将输入序列中是素数倍数的数提出，并返回新的序列，是一个新的管道。

// 管道过滤器: 删除能被素数整除的数
func PrimeFilter(ctx context.Context, in <-chan int, prime int) chan int {
	out := make(chan int)
	go func() {
		for {
			if i := <-in; i%prime != 0 {
				select {
				case <-ctx.Done():
					return
				case out <- i:
				}
			}
		}
	}()
	return out
}

// primeFilter函数也是内部启动一个Goroutine生产序列，返回过滤后序列对应的管道。
// 现在我们可以在main函数中驱动这个并发的素数筛了：
func main() {
	// 通过 Context 控制后台Goroutine状态
	ctx, cancel := context.WithCancel(context.Background())

	ch := GenerateNatural(ctx) // 自然数序列: 2, 3, 4, ...
	for i := 0; i < 100; i++ {
		prime := <-ch // 新出现的素数
		fmt.Printf("%v: %v\n", i+1, prime)
		ch = PrimeFilter(ctx, ch, prime) // 基于新素数构造的过滤器
	}

	cancel()
}

// 先是调用GenerateNatural()生成最原始的从2开始的自然数序列。
// 然后开始一个100次迭代的循环，希望生成100个素数。
// 在每次循环迭代开始的时候，管道中的第一个数必定是素数，我们先读取并打印这个素数。
// 然后基于管道中剩余的数列，并以当前取出的素数为筛子过滤后面的素数。
// 不同的素数筛子对应的管道是串联在一起的。
//
// 素数筛展示了一种优雅的并发程序结构。
// 但是因为每个并发体处理的任务粒度太细微，程序整体的性能并不理想。
// 对于细粒度的并发程序，CSP模型中固有的消息传递的代价太高了（多线程并发模型同样要面临线程启动的代价）。
