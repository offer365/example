package dao

import (
	"github.com/jmoiron/sqlx"
	"strings"
)

type DB interface {
	Init(opts ...Option) (db *sqlx.DB, err error)
}

// 建库sql: CREATE DATABASE IF NOT EXISTS golang DEFAULT CHARSET utf8 COLLATE utf8_general_ci;
func NewDB(driver string) (db DB) {
	switch strings.ToLower(driver) {
	case "mysql":
		return new(MysqlCli)
	}
	return nil
}

//dsn:="root:bd5bd7e9e2c2c26fa8b4e791ff1428c1@tcp(140.143.244.118:3306)/golang_db"

type Options struct {
	Host        string
	Port        string
	Username    string
	Password    string
	Database    string
	MaxConn     int // 最大连接数
	MaxIdleConn int // 最小空闲连接
	CharSet     string
}

func DefaultOpts() *Options {
	return &Options{
		Host:        "127.0.0.1",
		Port:        "3306",
		Username:    "root",
		Password:    "",
		Database:    "",
		MaxConn:     100,
		MaxIdleConn: 16,
		CharSet:     "utf8",
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

func WithMaxConn(max int) Option {
	return func(opts *Options) {
		opts.MaxConn = max
	}
}

func WithMaxIdleConn(idle int) Option {
	return func(opts *Options) {
		opts.MaxIdleConn = idle
	}
}

func WithCharSet(char string) Option {
	return func(opts *Options) {
		opts.CharSet = char
	}
}
