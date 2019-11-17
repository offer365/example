package proto

import (
	"context"
	"time"
)

func NewNode() *Node {
	return &Node{
		Name:  "server.io",
		Addr:  "127.0.0.1:1234",
		Peers: nil,
		Start: time.Now().UnixNano(),
		Hwmd5: "xxx",
		Now:   time.Now().UnixNano(),
	}
}

func (t *Node) Status(ctx context.Context, args *Args) (*Node, error) {
	t.Now = time.Now().UnixNano()
	return t, nil
}
