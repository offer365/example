package client

import (
	"context"
	"fmt"
	pb "github.com/offer365/example/grpc_example/core/proto"
	"log"
	"testing"
)

func TestNewRpcClient(t *testing.T) {
	conn,err:=NewRpcClient()
	fmt.Println(err)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	cli := pb.NewHelloServiceClient(conn)
	for i:=0;i<10;i++{
		reply, err := cli.Hello(context.Background(), &pb.String{Value: "hello"})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(reply.GetValue())
	}
}
