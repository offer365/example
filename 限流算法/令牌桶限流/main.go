package main

import (
	"fmt"
	"time"
)
// 谷歌实现的令牌桶算法
import "golang.org/x/time/rate"

func main() {
	limit := rate.NewLimiter(50, 100)
	for i := 0; i < 1000; i++ {
		time.Sleep(time.Millisecond * 10)
		if limit.Allow() {
			fmt.Printf("%d,allow\n", i)
			continue
		}
		fmt.Printf("%d,not allow\n", i)
	}
}
