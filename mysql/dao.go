package mysql

import (
	"strings"

	"github.com/jmoiron/sqlx"
)

type DB interface {
	Init(opts ...Option) (db *sqlx.DB, err error)
}

// 建库sql: CREATE DATABASE IF NOT EXISTS `giligili` DEFAULT CHARACTER SET 'utf8mb4' DEFAULT COLLATE 'utf8mb4_unicode_ci';
func NewDB(driver string) (db DB) {
	switch strings.ToLower(driver) {
	case "mysql":
		return new(mysqlCli)
	}
	return nil
}

// dsn:="root:bd5bd7e9e2c2c26fa8b4e791ff1428c1@tcp(140.143.244.118:3306)/golang_db"

type Options struct {
	Addr        string
	Username    string
	Password    string
	Database    string
	MaxConn     int // 最大连接数
	MaxIdleConn int // 最大空闲连接数
	CharSet     string
	ParseTime   bool
}

func DefaultOpts() *Options {
	return &Options{
		Addr:        "127.0.0.1:3306",
		Username:    "root",
		Password:    "",
		Database:    "",
		MaxConn:     100,
		MaxIdleConn: 16,
		CharSet:     "utf8",
		ParseTime:   false,
	}
}

type Option func(opts *Options)

func WithAddr(addr string) Option {
	return func(opts *Options) {
		opts.Addr = addr
	}
}

func WithUsername(user string) Option {
	return func(opts *Options) {
		opts.Username = user
	}
}

func WithPassword(pwd string) Option {
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

func WithParseTime(p bool) Option {
	return func(opts *Options) {
		opts.ParseTime = p
	}
}
