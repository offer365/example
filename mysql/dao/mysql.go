package dao

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type MysqlCli struct {
	options *Options
	host    string
	db      *sqlx.DB
}

func (m *MysqlCli) Init(opts ...Option) (db *sqlx.DB, err error) {
	m.options = DefaultOpts()
	for _, opt := range opts {
		opt(m.options)
	}
	//dsn:="root:666666@tcp(127.0.0.1:3306)/golang_db?charset=utf8"
	dsn := m.options.Username + ":" + m.options.Password + "@tcp(" + m.options.Host + ":" + m.options.Port + ")/" + m.options.Database
	if m.options.CharSet != "" {
		dsn += "?charset=" + m.options.CharSet
	}
	m.db, err = sqlx.Open("mysql", dsn)
	m.db.SetMaxOpenConns(m.options.MaxConn)     // 设置最大打开的连接数
	m.db.SetMaxIdleConns(m.options.MaxIdleConn) //  设置最大空闲连接数
	db = m.db
	return
}
