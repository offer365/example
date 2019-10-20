# 安装

1. 安装[官方的protoc工具](https://github.com/protocolbuffers/protobuf/releases)
2. 安装针对Go语言的代码生成插件，可以通过 `go get github.com/golang/protobuf/protoc-gen-go`
3. 通过以下命令生成相应的Go代码： `protoc --go_out=. hello.proto`
4. protoc-gen-go内部已经集成了一个名字为grpc的插件，可以针对gRPC生成代码 `protoc --go_out=plugins=grpc:. hello.proto`


# 安装 gRpc

安装官方安装命令：

`go get google.golang.org/grpc`

是安装不起的，会报：

`package google.golang.org/grpc: unrecognized import path "google.golang.org/grpc"(https fetch: Get https://google.golang.org/grpc?go-get=1: dial tcp 216.239.37.1:443: i/o timeout)`

正确的安装方式：

```git clone https://github.com/grpc/grpc-go.git $GOPATH/src/google.golang.org/grpc

git clone https://github.com/golang/net.git $GOPATH/src/golang.org/x/net

git clone https://github.com/golang/text.git $GOPATH/src/golang.org/x/text

go get -u github.com/golang/protobuf/{proto,protoc-gen-go}

git clone https://github.com/google/go-genproto.git $GOPATH/src/google.golang.org/genproto

cd $GOPATH/src/

go install google.golang.org/grpc```