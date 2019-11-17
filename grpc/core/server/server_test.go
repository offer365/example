package server

import (
	"context"
	"fmt"
	"net"
	"testing"

	pb "github.com/offer365/example/grpc/core/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type HelloServiceImpl struct{}

// 实现HelloServiceServer接口
func (p *HelloServiceImpl) Hello(
	ctx context.Context, args *pb.String,
) (*pb.String, error) {
	reply := &pb.String{Value: "hello:" + args.GetValue()}
	return reply, nil
}

func TestNewRpcServer(t *testing.T) {
	// grpcserver,listener,err:=NewRpcServer()
	// Token认证
	auth := func(ctx context.Context) error {
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return status.Errorf(codes.Unauthenticated, "missing credentials")
		}

		var user string
		var pwd string

		if val, ok := md["user"]; ok {
			user = val[0]
		}
		if val, ok := md["password"]; ok {
			pwd = val[0]
		}

		if user != "offer365" || pwd != "666666" {
			return status.Errorf(codes.Unauthenticated, "invalid token")
		}

		return nil
	}

	// 一元拦截器
	var interceptor grpc.UnaryServerInterceptor
	interceptor = func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		err = auth(ctx)
		if err != nil {
			return
		}
		// 继续处理请求
		return handler(ctx, req)
	}

	grpcserver, err := NewRpcServer(
		WithCa([]byte(Ca_crt)),
		WithKey([]byte(Server_key)),
		WithCert([]byte(Server_crt)),
		WithServerOption(grpc.UnaryInterceptor(interceptor)),
	)
	fmt.Println(err)
	helloServer := new(HelloServiceImpl)
	pb.RegisterHelloServiceServer(grpcserver, helloServer)
	listener, err := net.Listen("tcp", ":1234")
	grpcserver.Serve(listener)
}
