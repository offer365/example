package main

import (
	"context"
	"fmt"
	"log"

	cc "github.com/offer365/example/grpc/core/client"
	pb "github.com/offer365/example/grpc/和Web服务共存/proto"
)

const PORT = "9003"

func main() {

	conn, err := cc.NewRpcClient(
		cc.WithAddr("127.0.0.1:"+PORT),
		cc.WithServerName("server.io"),
		cc.WithCertFile(`C:\Users\Administrator\go\src\github.com\offer365\example\grpc\core\cert\client.crt`),
		cc.WithKeyFile(`C:\Users\Administrator\go\src\github.com\offer365\example\grpc\core\cert\client.key`),
		cc.WithCaFile(`C:\Users\Administrator\go\src\github.com\offer365\example\grpc\core\cert\ca.crt`),
	)

	if err != nil {
		log.Fatalf("grpc.Dial err: %v", err)
	}
	defer conn.Close()

	client := pb.NewHelloServiceClient(conn)
	resp, err := client.Hello(context.Background(), &pb.String{
		Value: "gRPC",
	})
	if err != nil {
		log.Fatalf("client.Search err: %v", err)
	}

	fmt.Printf("resp: %s", resp.Value)
}
