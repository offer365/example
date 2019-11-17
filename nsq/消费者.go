package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/nsqio/go-nsq"
)

type Consumer struct{}

// 实现一个固定的接口 名字是固定的 HaandleMessage
func (*Consumer) HandleMessage(m *nsq.Message) (err error) {
	fmt.Println("receive", m.NSQDAddress, "message:", string(m.Body))
	return nil
}

// 初始化消费者
func InitConsumer(topic, channel, address string) (err error) {
	// 使用默认配置
	cfg := nsq.NewConfig()
	// 轮询发现 lookupd
	cfg.LookupdPollInterval = 15 * time.Second     // 设置服务发现的轮询时间
	c, err := nsq.NewConsumer(topic, channel, cfg) // 新建一个消费者
	if err != nil {
		return err
	}
	consumer := &Consumer{}
	// 如果接收到消息会回调 HandleMessage 这个接口
	c.AddHandler(consumer) // 添加消费者接口

	// 建立 nsqlookupd 连接
	if err := c.ConnectToNSQLookupd(address); err != nil {
		return err
	}
	return nil
}

func main() {
	// 传入 队列，管道，nsqlookupd 的地址
	err := InitConsumer("a", "b", "140.143.244.118:4161")
	if err != nil {
		fmt.Printf("init failed,%#v\n", err)
		return
	}
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT)
	<-c
}
