package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

// 发布订阅（publish-and-subscribe）模型通常被简写为pub/sub模型。
// 在这个模型中，消息生产者成为发布者（publisher），而消息消费者则成为订阅者（sub），生产者和消费者是M:N的关系。
// 在传统生产者和消费者模型中，是将消息发送到一个队列中，而发布订阅模型则是将消息发布给一个主题。
// 为此，我们构建了一个名为pubsub的发布订阅模型支持包：
// Package pubsub implements a simple multi-topic pub-sub library.

type (
	sub   chan interface{}         // 订阅者为一个管道
	topic func(v interface{}) bool // 主题为一个过滤器
)

// 发布者对象
type Publisher struct {
	sync.RWMutex               // 读写锁
	buf          int           // 订阅队列的缓存大小
	timeout      time.Duration // 发布超时时间
	subs         map[sub]topic // 订阅者信息
}

// 构建一个发布者对象, 可以设置发布超时时间和缓存队列的长度
func NewPublisher(publishTimeout time.Duration, buffer int) *Publisher {
	return &Publisher{
		buf:     buffer,
		timeout: publishTimeout,
		subs:    make(map[sub]topic),
	}
}

// 添加一个新的订阅者，订阅全部主题
func (p *Publisher) Subscribe() chan interface{} {
	return p.SubscribeTopic(nil)
}

// 添加一个新的订阅者，订阅过滤器筛选后的主题
func (p *Publisher) SubscribeTopic(topic topic) chan interface{} {
	ch := make(chan interface{}, p.buf)
	p.Lock()
	p.subs[ch] = topic
	p.Unlock()
	return ch
}

// 退出订阅
func (p *Publisher) Evict(sub chan interface{}) {
	p.Lock()
	defer p.Unlock()

	delete(p.subs, sub)
	close(sub)
}

// 发布一个主题
func (p *Publisher) Publish(v interface{}) {
	p.RLock()
	defer p.RUnlock()

	var wg sync.WaitGroup
	for sub, topic := range p.subs { // 所有的订阅者
		wg.Add(1)
		go p.sendTopic(sub, topic, v, &wg)
	}
	wg.Wait()
}

// 关闭发布者对象，同时关闭所有的订阅者管道。
func (p *Publisher) Close() {
	p.Lock()
	defer p.Unlock()

	for sub := range p.subs {
		delete(p.subs, sub)
		close(sub) // 关闭管道
	}
}

// 发送主题，可以容忍一定的超时
func (p *Publisher) sendTopic(
	sub sub, topic topic, v interface{}, wg *sync.WaitGroup,
) {
	defer wg.Done()
	// 判断是否是全部订阅  是否是订阅的topic
	if topic != nil && !topic(v) {
		return
	}

	select {
	case sub <- v:
	case <-time.After(p.timeout):
	}
}

func main() {
	p := NewPublisher(100*time.Millisecond, 10)
	defer p.Close()

	all := p.Subscribe()
	golang := p.SubscribeTopic(func(v interface{}) bool {
		if s, ok := v.(string); ok {
			return strings.Contains(s, "golang")
		}
		return false
	})

	p.Publish("hello,  world!")
	p.Publish("hello, golang!")
	// 发布者
	go func() {
		for range time.NewTicker(time.Second).C {
			p.Publish("hehe")
		}
	}()

	// 订阅者
	go func() {
		for msg := range all {
			fmt.Println("all:", msg)
		}
	}()

	go func() {
		for msg := range golang {
			fmt.Println("golang:", msg)
		}
	}()

	// 运行一定时间后退出
	time.Sleep(9 * time.Second)
}
