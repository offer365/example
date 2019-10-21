package embedder

import (
	"context"
	"strings"
)

type Options struct {
	name         string
	dir          string
	clientAddr   string
	peerAddr     string
	cluster      map[string]string
	clusterState string // "new" or "existing"
	metrics      string
	metricsUrl   string
}

type Option func(opts *Options)

func DefaultOpts() *Options {
	return &Options{
		name:         "default",
		dir:          "disk/default",
		clientAddr:   "127.0.0.1:12379",
		peerAddr:     "127.0.0.1:12380",
		cluster:      map[string]string{"default": "127.0.0.1"},
		clusterState: "new",
		metrics:      "",
		metricsUrl:   "",
	}
}

func WithName(name string) Option {
	return func(opts *Options) {
		opts.name = name
	}
}

func WithDir(dir string) Option {
	return func(opts *Options) {
		opts.dir = dir
	}
}

func WithClientAddr(addr string) Option {
	return func(opts *Options) {
		opts.clientAddr = addr
	}
}

func WithPeerAddr(addr string) Option {
	return func(opts *Options) {
		opts.peerAddr = addr
	}
}

func WithCluster(cluster map[string]string) Option {
	return func(opts *Options) {
		opts.cluster = cluster
	}
}

func WithClusterState(state string) Option {
	return func(opts *Options) {
		// "new" or "existing"
		if strings.HasPrefix(state, "exist") {
			opts.clusterState = "existing"
		} else {
			opts.clusterState = "new"
		}
	}
}

func WithMetrics(addr string, mode string) Option {
	return func(opts *Options) {
		switch {
		case strings.HasPrefix(mode, "b"):
			opts.metrics = "base"
		case strings.HasPrefix(mode, "e"):
			opts.metrics = "extensive"
		default:
			opts.metrics = "base"
		}
		if addr != "" && !strings.HasPrefix(addr, "http://") {
			opts.metricsUrl = "http://" + addr
			return
		}
		opts.metricsUrl = addr
	}
}

type Embed interface {
	Init(ctx context.Context, option ...Option) (err error)
	Run(ready chan struct{}) (err error)
	SetAuth(username, password string) (err error)
	IsLeader() bool
	Close()
}

func NewEmbed() Embed {
	return new(etcdEmbed)
}
