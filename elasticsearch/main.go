package main

import (
	"context"
	"fmt"

	"gopkg.in/olivere/elastic.v6"
)

type Tweet struct {
	User    string
	Message string
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	client, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL("http://10.10.5.230:9200"))
	if err != nil {
		fmt.Println("connect error", err)
	}
	fmt.Println("conn es succ")
	tweet := Tweet{"olivere", "take five"}
	// Index("twitter") 类似于数据库  Type("tweet") 类似于 表
	_, err = client.Index().Index("twitter").Type("tweet").Id("1").BodyJson(tweet).Do(ctx)
	if err != nil {
		panic(err)
		return
	}
	fmt.Println("insert succ")
}
