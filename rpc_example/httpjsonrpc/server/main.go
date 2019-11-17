package main

import (
	"io"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type HelloService struct{}

func (p *HelloService) Hello(request string, reply *string) error {
	*reply = "hello:" + request
	return nil
}

// Go语言的RPC框架有两个比较有特色的设计：
// 一个是RPC数据打包时可以通过插件实现自定义的编码和解码；
// 另一个是RPC建立在抽象的io.ReadWriteCloser接口之上的，我们可以将RPC架设在不同的通讯协议之上。
// 这里我们将尝试通过官方自带的net/rpc/jsonrpc扩展实现一个跨语言的RPC。
func main() {
	rpc.RegisterName("HelloService", new(HelloService))

	http.HandleFunc("/rpc", func(w http.ResponseWriter, r *http.Request) {
		var conn io.ReadWriteCloser = struct {
			io.Writer // 嵌套
			io.ReadCloser
		}{
			ReadCloser: r.Body,
			Writer:     w,
		}

		rpc.ServeRequest(jsonrpc.NewServerCodec(conn))
	})

	http.ListenAndServe(":1234", nil)
}
