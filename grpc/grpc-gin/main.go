package main

import (
	"context"
	"crypto/rand"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	corec "github.com/offer365/example/grpc/core/client"
	cores "github.com/offer365/example/grpc/core/server"
	"github.com/offer365/example/grpc/grpc-gin/proto"
	"github.com/offer365/odin/log"
	"golang.org/x/net/http2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

var token = "666666"
var port = ":4567"

func gRpcServer() *grpc.Server {
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

		if user != token || pwd != token {
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

	gs, err := cores.NewRpcServer(
		cores.WithServerOption(grpc.UnaryInterceptor(interceptor)),
		cores.WithCert([]byte(cores.Server_crt)),
		cores.WithKey([]byte(cores.Server_key)),
		cores.WithCa([]byte(cores.Ca_crt)),
	)
	fmt.Println(err)
	proto.RegisterStaterServer(gs, proto.NewNode())
	return gs

}

func ginServer() http.Handler {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.GET("/test", func(c *gin.Context) {
		c.String(200, "hello")
	})

	return r
}

func listener() net.Listener {
	certificate, err := tls.X509KeyPair([]byte(cores.Server_crt), []byte(cores.Server_key))
	fmt.Println("err:", err)
	certPool := x509.NewCertPool()

	if ok := certPool.AppendCertsFromPEM([]byte(cores.Ca_crt)); !ok {
		fmt.Println(ok)
	}
	tlsConfig := &tls.Config{
		Certificates:       []tls.Certificate{certificate},
		ClientAuth:         tls.NoClientCert, // NOTE: 这是可选的!
		ClientCAs:          certPool,
		InsecureSkipVerify: true,
		Rand:               rand.Reader,
		Time:               time.Now,
		NextProtos:         []string{"http/1.1", http2.NextProtoTLS},
	}
	lis, err := tls.Listen("tcp", port, tlsConfig)
	fmt.Println("err:", err)
	return lis
}

func gateway() {
	handle := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			gRpcServer().ServeHTTP(w, r) // grpc server
		} else {
			ginServer().ServeHTTP(w, r) // gin web server
		}
		return
	})
	lis := listener()
	http.Serve(lis, handle)
}

type Authentication struct {
	User     string
	Password string
}

func (a *Authentication) GetRequestMetadata(context.Context, ...string) (map[string]string, error) {
	return map[string]string{"user": a.User, "password": a.Password}, nil
}
func (a *Authentication) RequireTransportSecurity() bool {
	return true
}

func client() {
	auth := &Authentication{
		User:     token,
		Password: token,
	}

	Con, err := corec.NewRpcClient(
		corec.WithAddr("127.0.0.1"+port),
		// corec.WithTimeout(8000*time.Millisecond),
		corec.WithDialOption(grpc.WithPerRPCCredentials(auth)),
		corec.WithServerName("server.io"),
		corec.WithCert([]byte(cores.Client_crt)),
		corec.WithKey([]byte(cores.Client_key)),
		corec.WithCa([]byte(cores.Ca_crt)),
	)
	if err != nil {
		log.Sugar.Error(err)
		return
	}
	cli := proto.NewStaterClient(Con)
	for range time.NewTicker(1 * time.Second).C {
		node, err := cli.Status(context.Background(), &proto.Args{Name: "super", Addr: "127.0.0.1:1234"})
		fmt.Println(node, err)
	}

	return
}

func main() {
	go client()
	gateway()
}
