package dao

import (
	"context"
	"io"
	"strings"
	"time"
)

// 数据库
type DB interface {
	Init(ctx context.Context, option ...Option) (err error)
	Insert(table string, instance interface{}) (id string, err error)
	Find(table string, filter interface{}, queryF interface{}, skip int64, limit int64, sort int) error
	FindOne(coll string, filter interface{}, result interface{}) (err error)
	Delete(coll string, filter interface{}) (err error)
	Update(coll string, filter, update interface{}) (err error)
	Count(coll string, filter interface{}) (num int64, err error)
	Aggregation(coll string, pipe interface{}, cursorF interface{}) (err error) // 聚合查询
}

// 存储
type Store interface {
	Init(ctx context.Context, option ...Option) (err error)
	Upload(name string, source io.Reader) (id string, err error)
	Download(id interface{}, stream io.Writer) (size int64, err error)
	FindFile(filter interface{}, cursorF interface{}, skip, limit int32) (err error)
	DeleteFile(id interface{}) (err error)
}

func NewDB(driver string) (db DB) {
	switch strings.ToLower(driver) {
	case "mongodb", "mongo":
		return new(MongoCli)
	}
	return nil
}

func NewStore(driver string) (store Store) {
	switch strings.ToLower(driver) {
	case "mongodb", "mongo":
		return new(MongoCli)
	}
	return nil
}

type Options struct {
	Host      string
	Port      string
	Username  string
	Password  string
	Timeout   time.Duration
	Database  string
	CollIndex map[string]string
}

func DefaultOpts() *Options {
	return &Options{
		Host:      "127.0.0.1",
		Port:      "27017",
		Username:  "admin",
		Password:  "",
		Timeout:   2 * time.Second,
		Database:  "",
		CollIndex: nil,
	}
}

type Option func(opts *Options)

func WithHost(host string) Option {
	return func(opts *Options) {
		opts.Host = host
	}
}

func WithPort(port string) Option {
	return func(opts *Options) {
		opts.Port = port
	}
}

func WithTimeout(timeout time.Duration) Option {
	return func(opts *Options) {
		opts.Timeout = timeout
	}
}

func WithUsername(user string) Option {
	return func(opts *Options) {
		opts.Username = user
	}
}

func WithPwd(pwd string) Option {
	return func(opts *Options) {
		opts.Password = pwd
	}
}

func WithDatabase(db string) Option {
	return func(opts *Options) {
		opts.Database = db
	}
}

// 集合索引
func WithCollIndex(ci map[string]string) Option {
	return func(opts *Options) {
		opts.CollIndex = ci
	}
}
