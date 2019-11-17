package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

type CounterLimit struct {
	counter  int64 // 计数器
	limit    int64 // 指定时间窗口允许的最大请求数
	interval int64 // 时间窗口 纳秒
	unix     int64 // 时间戳 纳秒
}

func NewCounterLimit(interval time.Duration, limit int64) *CounterLimit {
	return &CounterLimit{
		counter:  0,
		limit:    limit,
		interval: int64(interval),
		unix:     time.Now().UnixNano(),
	}
}

func (cl *CounterLimit) Allow() bool {
	now := time.Now().UnixNano()
	if now-cl.unix > cl.interval { // 如果大于时间窗口重新进行计数
		atomic.StoreInt64(&cl.counter, 0)
		atomic.StoreInt64(&cl.unix, now)
		return true
	}
	atomic.AddInt64(&cl.counter, 1)
	return cl.counter < cl.limit // 判断是否进行限流
}

func main() {
	limit := NewCounterLimit(time.Second, 100)
	for i := 0; i < 1000; i++ {
		if limit.Allow() {
			fmt.Printf("%d,allow\n", i)
			continue
		}
		fmt.Printf("%d,not allow\n", i)
	}
}
