package main

import (
	"net"
	"net/rpc"
	"time"
)

// 内网 服务端

type HelloService struct{}

func (p *HelloService) Hello(request string, reply *string) error {
	*reply = "hello:" + request
	return nil
}

// RPC是基于C/S结构，RPC的服务端对应网络的服务器，RPC的客户端也对应网络客户端。
// 但是对于一些特殊场景，比如在公司内网提供一个RPC服务，但是在外网无法链接到内网的服务器。
// 这种时候我们可以参考类似反向代理的技术，首先从内网主动链接到外网的TCP服务器，然后基于TCP链接向外网提供RPC服务。
// 反向RPC的内网服务将不再主动提供TCP监听服务，而是首先主动链接到对方的TCP服务器。然后基于每个建立的TCP链接向对方提供RPC服务。
func main() {
	rpc.Register(new(HelloService))

	for {
		conn, _ := net.Dial("tcp", "127.0.0.1:1234")
		if conn == nil {
			time.Sleep(time.Second)
			continue
		}

		rpc.ServeConn(conn)
		conn.Close()
	}
}
