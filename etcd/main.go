package main

import (
	"context"
	"fmt"
	etcdClient "go.etcd.io/etcd/clientv3"
	"time"
)

func example() {
	cli, err := etcdClient.New(etcdClient.Config{
		// etcd 的全部节点
		Endpoints:   []string{"10.10.5.230:2379"},
		DialTimeout: 3 * time.Second,
	})
	if err != nil {
		fmt.Println("connect failed,", err)
		return
	}
	fmt.Println("connect success")
	defer cli.Close()
}

func main() {
	example2()
}

func example2() {
	cli, err := etcdClient.New(
		etcdClient.Config{
			Endpoints:   []string{"10.10.5.230:2379"},
			DialTimeout: 3 * time.Second,
		})
	if err != nil {
		fmt.Println("connect failed,", err)
		return
	}
	fmt.Println("connect success")
	defer cli.Close()

	ctx, cancle := context.WithTimeout(context.Background(), time.Second)
	_, err = cli.Put(ctx, "test1", "1111")
	cancle()
	if err != nil {
		fmt.Println(err)
	}

	ctx, cancle = context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Get(ctx, "test1")
	cancle()
	if err != nil {
		fmt.Println(err)
	}
	for _, v := range resp.Kvs {
		fmt.Println(string(v.Key), string(v.Value))
	}

}
