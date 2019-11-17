package main

import "sync"
import (
	"fmt"
)

// 在main函数所在线程中执行两次mu.Lock()，当第二次加锁时会因为锁已经被占用（不是递归锁）而阻塞，main函数的阻塞状态驱动后台线程继续向前执行。
// 当后台线程执行到mu.Unlock()时解锁，此时打印工作已经完成了，解锁会导致main函数中的第二个mu.Lock()阻塞状态取消，
// 此时后台线程和主线程再没有其它的同步事件参考，它们退出的事件将是并发的：在main函数退出导致程序退出时，后台线程可能已经退出了，也可能没有退出。
// 虽然无法确定两个线程退出的时间，但是打印工作是可以正确完成的。
//
// 使用sync.Mutex互斥锁同步是比较低级的做法。
func example1() {
	var mu sync.Mutex

	mu.Lock()
	go func() {
		fmt.Println("你好, 世界")
		mu.Unlock()
	}()

	mu.Lock()
}

// 对于从无缓冲Channel进行的接收，发生在对该Channel进行的发送完成之前。
// 因此，后台线程<-done接收操作完成之后，main线程的done <- 1发送操作才可能完成（从而退出main、退出程序），而此时打印工作已经完成了。
//
// 注释的代码虽然可以正确同步，但是对管道的缓存大小太敏感：如果管道有缓存的话，就无法保证main退出之前后台线程能正常打印了。
// 更好的做法是将管道的发送和接收方向调换一下，这样可以避免同步事件受管道缓存大小的影响：

func example2() {
	done := make(chan int, 1) // 带缓存的管道

	go func() {
		fmt.Println("你好, 世界")
		done <- 1
		// <- done
	}()

	<-done
	// done<-1
}

// wg.Add(1)用于增加等待事件的个数，必须确保在后台线程启动之前执行（如果放到后台线程之中执行则不能保证被正常执行到）。
// 当后台线程完成打印工作之后，调用wg.Done()表示完成一个事件。main函数的wg.Wait()是等待全部的事件完成。
func example3() {
	var wg sync.WaitGroup

	// 开N个后台打印线程
	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func() {
			fmt.Println("你好, 世界")
			wg.Done()
		}()
	}

	// 等待N个后台线程完成
	wg.Wait()
}

// 我们创建了一个带缓存的管道，管道的缓存数目要足够大，保证不会因为缓存的容量引起不必要的阻塞。
// 然后我们开启了多个后台线程，分别向不同的搜索引擎提交搜索请求。
// 当任意一个搜索引擎最先有结果之后，都会马上将结果发到管道中（因为管道带了足够的缓存，这个过程不会阻塞） 。
// 但是最终我们只从管道取第一个结果，也就是最先返回的结果。
//
// 通过适当开启一些冗余的线程，尝试用不同途径去解决同样的问题，最终以赢者为王的方式提升了程序的相应性能。
func example4() {
	ch := make(chan string, 32)

	go func() {
		ch <- searchByBing("golang")
	}()
	go func() {
		ch <- searchByGoogle("golang")
	}()
	go func() {
		ch <- searchByBaidu("golang")
	}()

	fmt.Println(<-ch)
}
