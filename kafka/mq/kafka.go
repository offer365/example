package mq

import (
	"fmt"
	"log"
	"sync"

	"github.com/Shopify/sarama"
)

type Producer struct {
	config       *sarama.Config
	syncProducer sarama.SyncProducer
}

func SyncProducer(host, topic, val string) {
	config := sarama.NewConfig()
	// 等待服务器所有副本都保存成功后的响应
	config.Producer.RequiredAcks = sarama.WaitForAll
	// 随机的分区类型：返回一个分区器，该分区器每次选择一个随机分区
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	// 是否等待成功和失败后的响应
	config.Producer.Return.Successes = true

	// 使用给定代理地址和配置创建一个同步生产者
	// producer,err:= sarama.NewSyncProducer([]string{"10.0.0.55:9092"}, config)
	producer, err := sarama.NewSyncProducer([]string{host}, config)
	if err != nil {
		fmt.Println(err)
	}
	defer func() {
		if err := producer.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	// 构建发送的消息，
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.StringEncoder("key"), //
	}

	for {
		fmt.Println("topic = ", topic, ",value = ", val)
		// 将字符串转换为字节数组
		msg.Value = sarama.ByteEncoder(val)
		// fmt.Println(value)
		// SendMessage：该方法是生产者生产给定的消息
		// 生产成功的时候返回该消息的分区和所在的偏移量
		// 生产失败的时候返回error
		partition, offset, err := producer.SendMessage(msg)

		if err != nil {
			fmt.Println(err)
			fmt.Println("Send message Fail")
		}
		fmt.Printf("Partition = %d, offset=%d\n", partition, offset)
	}
}

func Consumer2(host, topic string) {
	var wg sync.WaitGroup
	// 根据给定的代理地址和配置创建一个消费者
	consumer, err := sarama.NewConsumer([]string{host}, nil)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := consumer.Close(); err != nil {
			log.Fatalln(err)
		}
	}()
	// Partitions(topic):该方法返回了该topic的所有分区id
	partitionList, err := consumer.Partitions(topic)
	if err != nil {
		panic(err)
	}

	for partition := range partitionList {
		// ConsumePartition方法根据主题，分区和给定的偏移量创建创建了相应的分区消费者
		// 如果该分区消费者已经消费了该信息将会返回error
		// sarama.OffsetNewest:表明了为最新消息
		pc, err := consumer.ConsumePartition("test", int32(partition), sarama.OffsetNewest)
		if err != nil {
			panic(err)
		}
		defer pc.AsyncClose()
		wg.Add(1)
		go func(sarama.PartitionConsumer) {
			defer wg.Done()
			// Messages()该方法返回一个消费消息类型的只读通道，由代理产生
			for msg := range pc.Messages() {
				fmt.Printf("%s---Partition:%d, Offset:%d, Key:%s, Value:%s\n", msg.Topic, msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
			}
		}(pc)
	}
	wg.Wait()
	consumer.Close()
}

func AsyncProducer(host, topic, val string) {
	fmt.Printf("producer_test\n")
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true

	producer, err := sarama.NewAsyncProducer([]string{host}, config)
	if err != nil {
		fmt.Printf("producer_test create producer error :%s\n", err.Error())
		return
	}

	defer producer.AsyncClose()

	// send message
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.StringEncoder("go_test"),
	}

	for {
		msg.Value = sarama.ByteEncoder(val)
		fmt.Printf("input [%s]\n", val)

		// send to chain
		producer.Input() <- msg

		select {
		case suc := <-producer.Successes():
			fmt.Printf("offset: %d,  timestamp: %s\n", suc.Offset, suc.Timestamp.String())
		case fail := <-producer.Errors():
			fmt.Printf("err: %s\n", fail.Err.Error())
		}
	}
}

func Consumer(host, topic string) {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	// consumer
	consumer, err := sarama.NewConsumer([]string{host}, config)
	if err != nil {
		fmt.Printf("consumer_test create consumer error %s\n", err.Error())
		return
	}

	defer consumer.Close()

	pc, err := consumer.ConsumePartition(topic, 0, sarama.OffsetOldest)
	if err != nil {
		fmt.Printf("try create partition_consumer error %s\n", err.Error())
		return
	}
	defer pc.Close()

	for {
		select {
		case msg := <-pc.Messages():
			fmt.Printf("msg offset: %d, partition: %d, timestamp: %s, value: %s\n",
				msg.Offset, msg.Partition, msg.Timestamp.String(), string(msg.Value))
		case err := <-pc.Errors():
			fmt.Printf("err :%s\n", err.Error())
		}
	}
}
