package redis

import (
	"strings"
	"time"

	"github.com/gomodule/redigo/redis"
)

type Options struct {
	Addr      string
	Password  string
	Database  string
	Timeout   time.Duration
	MaxIdle   int
	MaxActive int
}

func DefaultOpts() *Options {
	return &Options{
		Addr:      "127.0.0.1:6379",
		Password:  "",
		Timeout:   20 * time.Second,
		MaxIdle:   60,
		MaxActive: 1000,
	}
}

type Option func(opts *Options)

func WithAddr(addr string) Option {
	return func(opts *Options) {
		opts.Addr = addr
	}
}

func WithDB(db string) Option {
	return func(opts *Options) {
		opts.Database = db
	}
}

func WithTimeout(timeout time.Duration) Option {
	return func(opts *Options) {
		opts.Timeout = timeout
	}
}

func WithPwd(pwd string) Option {
	return func(opts *Options) {
		opts.Password = pwd
	}
}

func WithMaxIdle(idle int) Option {
	return func(opts *Options) {
		opts.MaxIdle = idle
	}
}

func WithMaxActive(active int) Option {
	return func(opts *Options) {
		opts.MaxActive = active
	}
}

type Cache interface {
	Init(...Option)
	Conn() (pool redis.Conn)
}

func NewCache(driver string) Cache {
	switch strings.ToLower(driver) {
	case "redis":
		return new(RedisCli)
	}
	return nil
}
