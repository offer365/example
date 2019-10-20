package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"10.10.5.230:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println("connect failed")
	}
	fmt.Println("connect success")

	// cli.Watch 会一直监控这个key,并阻塞到这。直到这个key 发生变化，etcd 通知所有监听的链接。
	rch := cli.Watch(context.Background(), "test1")
	for wresp := range rch {
		for _, ev := range wresp.Events {
			fmt.Printf("%s %q :%q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
		}
	}
}
