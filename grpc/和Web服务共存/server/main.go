package main

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	cs "github.com/offer365/example/grpc/core/server"
	pb "github.com/offer365/example/grpc/和Web服务共存/proto"
)

// 参考 https://segmentfault.com/a/1190000016601836
// https://github.com/EDDYCJY/go-grpc-example

type HelloServiceImpl struct{}

// 实现HelloServiceServer接口
func (p *HelloServiceImpl) Hello(
	ctx context.Context, args *pb.String,
) (*pb.String, error) {
	reply := &pb.String{Value: "hello:" + args.GetValue()}
	return reply, nil
}

const PORT = "9003"

func main() {
	var err error
	helloServer := new(HelloServiceImpl)
	mux := GetHTTPServeMux()

	server, err := cs.NewRpcServer(
		cs.WithCertFile(`C:\Users\Administrator\go\src\github.com\offer365\example\grpc\core\cert\server.crt`),
		cs.WithKeyFile(`C:\Users\Administrator\go\src\github.com\offer365\example\grpc\core\cert\server.key`),
		cs.WithCaFile(`C:\Users\Administrator\go\src\github.com\offer365\example\grpc\core\cert\ca.crt`), )
	fmt.Println(err)
	pb.RegisterHelloServiceServer(server, helloServer)

	http.ListenAndServeTLS(":"+PORT,
		`C:\Users\Administrator\go\src\github.com\offer365\example\grpc\core\cert\server.crt`,
		`C:\Users\Administrator\go\src\github.com\offer365\example\grpc\core\cert\server.key`,
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
				server.ServeHTTP(w, r)
			} else {
				mux.ServeHTTP(w, r)
			}

			return
		}),
	)
}

func GetHTTPServeMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello"))
	})
	return mux
}
