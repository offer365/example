package kafka

import (
	"fmt"

	"github.com/Shopify/sarama"
)

// 追加 advertised.host.name=kafka服务器ip 到 kafka 配置文件 config/server.properties
// 命令行开启kafka消费者客户端命令
// ./bin/kafka-console-consumer.sh --bootstrap-server 127.0.0.1:9092 --topic asr_log
// Note：在0.9版本指定的是zookeeper server，0.11变成了broker server
// ./bin/kafka-console-consumer.sh -zookeeper localhost:2181 --topic asr_log

func kafka() {
	// 实例化一个 配置对象
	config := sarama.NewConfig()
	// ack 表示命令正确接收
	config.Producer.RequiredAcks = sarama.WaitForAll
	// 随机分区
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true

	// new 一个 kafka 的生产者实例 同步的
	client, err := sarama.NewSyncProducer([]string{"10.10.5.230:9092"}, config)
	if err != nil {
		fmt.Println("producer close,err:", err)
		return
	}
	defer client.Close()
	for {
		msg := &sarama.ProducerMessage{}
		msg.Topic = "asr_log"
		msg.Value = sarama.StringEncoder("this is test.")
		// 发送消息，返回分区id,偏移量。
		pid, offset, err := client.SendMessage(msg)
		if err != nil {
			fmt.Println("send message failed,", err)
			return
		}
		fmt.Printf("pid:%v offset:%v\n", pid, offset)
	}
}
