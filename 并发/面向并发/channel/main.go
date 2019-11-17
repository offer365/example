package main

// 对于带缓冲的Channel，对于Channel的第K个接收完成操作发生在第K+C个发送操作完成之前，其中C是Channel的缓存大小。
// 如果将C设置为0自然就对应无缓存的Channel，也即使第K个接收完成在第K个发送完成之前。
// 因为无缓存的Channel只能同步发1个，也就简化为前面无缓存Channel的规则：对于从无缓冲Channel进行的接收，发生在对该Channel进行的发送完成之前。
// 我们可以根据控制Channel的缓存大小来控制并发执行的Goroutine的最大数目, 例如:

var limit = make(chan int, 3)

func main() {
	var work = []func(){func() {}}
	for _, w := range work {
		go func() {
			limit <- 1
			w()
			<-limit
		}()
	}
	select {}
}
