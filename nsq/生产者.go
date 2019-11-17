package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/nsqio/go-nsq"
)

var producer *nsq.Producer

//  初始化一个生产者
func InitProducer(addr string) (err error) {
	// 生成一个默认的配置
	config := nsq.NewConfig()
	// 传入 地址,初始化一个生产者
	producer, err = nsq.NewProducer(addr, config)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	nsqAddress := "140.143.244.118:4150"
	err := InitProducer(nsqAddress)
	if err != nil {
		fmt.Printf("init failed,%#v\n", err)
		return
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		data, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("read failed,%#v\n", err)
			continue
		}

		data = strings.TrimSpace(data)
		if data == "stop" {
			break
		}
		// 将数据发布到 list1 队列
		err = producer.Publish("a", []byte(data))
		if err != nil {
			fmt.Printf("publish failed,err: %#v\n", err)
			continue
		}
		fmt.Printf("publish success,%s\n", data)
	}
}
