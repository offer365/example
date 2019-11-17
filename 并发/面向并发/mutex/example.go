package main

import "sync"

// 在worker的循环中，为了保证total.value += i的原子性，我们通过sync.Mutex加锁和解锁来保证该语句在同一时刻只被一个线程访问。
// 对于多线程模型的程序而言，进出临界区前后进行加锁和解锁都是必须的。
// 如果没有锁的保护，total的最终值将由于多线程之间的竞争而可能会不正确。
//
// 用互斥锁来保护一个数值型的共享资源，麻烦且效率低下

var total struct {
	sync.Mutex
	val int
}

// 必须是指针传递
func worker(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i <= 100; i++ {
		total.Lock()
		total.val += 1
		total.Unlock()
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go worker(&wg)
	go worker(&wg)
	wg.Wait()
}
