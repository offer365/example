package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
	"time"

	"github.com/docker/docker/pkg/pubsub"
	"google.golang.org/grpc"
)

type PubSubService struct {
	pub *pubsub.Publisher
}

func NewPubSubService() *PubSubService {
	return &PubSubService{pub: pubsub.NewPublisher(1000*time.Millisecond, 10)}
}

func (p *PubSubService) Publish(ctx context.Context, arg *String) (*String, error) {
	p.pub.Publish(arg.GetValue())
	return &String{}, nil
}

func (p *PubSubService) Subscribe(arg *String, stream PubSubService_SubscribeServer) error {
	ch := p.pub.SubscribeTopic(func(v interface{}) bool {
		if key, ok := v.(string); ok {
			if strings.HasPrefix(key, arg.GetValue()) {
				return true
			}
		}
		return false
	})

	for v := range ch {
		if err := stream.Send(&String{Value: v.(string)}); err != nil {
			return err
		}
	}

	return nil
}

func main() {
	go server()
	go client1()
	go client2()
	go client3()
	select {}
}

func server() {
	grpcServer := grpc.NewServer()
	RegisterPubSubServiceServer(grpcServer, NewPubSubService())

	lis, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}
	grpcServer.Serve(lis)
}

// 从客户端向服务器发布信息
func client1() {
	// grpc.WithInsecure()选项跳过了对服务器证书的验证
	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := NewPubSubServiceClient(conn)

	for range time.NewTicker(time.Millisecond * 500).C {
		_, err = client.Publish(
			context.Background(), &String{Value: "golang: hello Go"},
		)
		if err != nil {
			log.Fatal(err)
		}
		_, err = client.Publish(
			context.Background(), &String{Value: "docker: hello Docker"},
		)
	}

	if err != nil {
		log.Fatal(err)
	}
}

// 在另一个客户端进行订阅信息
func client2() {
	// grpc.WithInsecure()选项跳过了对服务器证书的验证
	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := NewPubSubServiceClient(conn)
	stream, err := client.Subscribe(
		context.Background(), &String{Value: "golang:"},
	)
	if err != nil {
		log.Fatal(err)
	}

	for {
		reply, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}

		fmt.Println(reply.GetValue())
	}
}

func client3() {
	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := NewPubSubServiceClient(conn)
	stream, err := client.Subscribe(
		context.Background(), &String{Value: "docker:"},
	)
	if err != nil {
		log.Fatal(err)
	}

	for {
		reply, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}

		fmt.Println(reply.GetValue())
	}
}
