package dao

import (
	"context"
	"time"

	"go.etcd.io/etcd/clientv3"
)

type Options struct {
	Addr     string
	Timeout  time.Duration
	Username string
	Password string
}
type Option func(opts *Options)

type EventFunc func(event *clientv3.Event) error

func WithAddr(addr string) Option {
	return func(opts *Options) {
		opts.Addr = addr
	}
}

func WithTimeout(timeout time.Duration) Option {
	return func(opts *Options) {
		opts.Timeout = timeout
	}
}

func WithUsername(username string) Option {
	return func(opts *Options) {
		opts.Username = username
	}
}

func WithPassword(pwd string) Option {
	return func(opts *Options) {
		opts.Password = pwd
	}
}

type Store interface {
	Init(ctx context.Context, option ...Option) (err error)
	Get(key string, lock bool) (resp *clientv3.GetResponse, err error)
	GetAll(prefix string, lock bool) (resp *clientv3.GetResponse, err error)
	Count(prefix string, lock bool) (resp *clientv3.GetResponse, err error)
	Put(key, val string, lock bool) (resp *clientv3.PutResponse, err error)
	Del(key string, lock bool) (resp *clientv3.DeleteResponse, err error)
	DelAll(prefix string, lock bool) (resp *clientv3.DeleteResponse, err error)
	Lease(key string, ttl int64) (resp *clientv3.LeaseGrantResponse, err error)
	PutWithLease(key, val string, leaseId clientv3.LeaseID, lock bool) (resp *clientv3.PutResponse, err error)
	DelWithLease(key string, leaseId clientv3.LeaseID, lock bool) (resp *clientv3.DeleteResponse, err error)
	KeepOnce(leaseId clientv3.LeaseID) (resp *clientv3.LeaseKeepAliveResponse, err error)
	KeepAlive(ctx context.Context, leaseId clientv3.LeaseID) (err error)
	Watch(ctx context.Context, key string, putFunc EventFunc, delFunc EventFunc)
	MoveLeader(transfereeID uint64) (resp *clientv3.MoveLeaderResponse, err error)
	MemberList() (resp *clientv3.MemberListResponse, err error)
	Close()
}

func NewStore() Store {
	return new(etcdStore)
}
