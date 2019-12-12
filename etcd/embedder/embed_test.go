package embedder

import (
	"context"
	"fmt"
	"testing"
	"time"

	"go.etcd.io/etcd/pkg/logutil"
	"go.uber.org/zap"
)

var sugar *zap.SugaredLogger

func init()  {
	lg, _ := zap.NewProduction()
	defer lg.Sync()
	cfg := logutil.DefaultZapLoggerConfig
	// cfg.Level = zap.NewAtomicLevelAt(zap.WarnLevel)
	cfg.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	lg, _ = cfg.Build()
	sugar = lg.Sugar()
}

func TestNewEmbed(t *testing.T) {
	embed := NewEmbed()
	err := embed.Init(context.Background(),
		WithName("default"),
		WithDir("../disk"),
		WithClientAddr("127.0.0.1:12379"),
		WithPeerAddr("127.0.0.1:12380"),
		WithCluster(map[string]string{"default": "127.0.0.1:12380"}),
		WithClusterState("new"),
		WithLogger(sugar),
		)
	// err := embed.Init(context.Background())
	fmt.Println(err)
	t.Error(err)
	ready := make(chan struct{})
	go func() {
		err := embed.Run(ready)
		fmt.Println(err)
		t.Error(err)
	}()
	select {
	case <-ready:
		err = embed.SetAuth("root", "613f#8d164df4ACPF49@93a510df49!66f98b*d6")
		fmt.Println(err)
		t.Error(err)
	}
	time.Sleep(10 * time.Second)
}
