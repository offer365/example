package main

import (
	"fmt"
	"math"
	"time"
)

type BucketLimit struct {
	rate    float64 // 漏桶 露出的速率
	size    float64 // 漏桶大小
	unix    int64   // 时间戳
	current float64 // 当前桶的大小
}

func NewBucketLimit(size int64, rate float64) *BucketLimit {
	return &BucketLimit{
		rate:    rate,
		size:    float64(size),
		unix:    time.Now().UnixNano(),
		current: 0,
	}
}

func (bl *BucketLimit) refresh() {
	now := time.Now().UnixNano()
	diff := float64(now-bl.unix) / 1000 / 1000 / 1000
	bl.current = math.Max(0, bl.current-diff*bl.rate)
	bl.unix = now
	return
}

func (bl *BucketLimit) Allow() bool {
	bl.refresh()
	if bl.current < bl.size {
		bl.current += 1
		return true
	}
	return false
}

func main() {
	// 限速50qps
	limit := NewBucketLimit(100, 50)
	for i := 0; i < 1000; i++ {
		time.Sleep(10 * time.Millisecond)
		if limit.Allow() {
			fmt.Printf("%d,allow\n", i)
			continue
		}
		fmt.Printf("%d,not allow\n", i)
	}
}
