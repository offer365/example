package mysql

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type mysqlCli struct {
	options *Options
	host    string
	db      *sqlx.DB
}

func (m *mysqlCli) Init(opts ...Option) (db *sqlx.DB, err error) {
	m.options = DefaultOpts()
	for _, opt := range opts {
		opt(m.options)
	}
	// dsn:="root:666666@tcp(127.0.0.1:3306)/golang_db?charset=utf8?parseTime=true"
	dsn := m.options.Username + ":" + m.options.Password + "@tcp(" + m.options.Addr + ")/" + m.options.Database
	if m.options.CharSet != "" {
		dsn += "?charset=" + m.options.CharSet
	}
	// 将mysql 里面的时间字段转成go 里面的时间字段 time.Time
	if m.options.ParseTime {
		dsn += "?parseTime=true"
	}
	m.db, err = sqlx.Open("mysql", dsn)
	m.db.SetMaxOpenConns(m.options.MaxConn)     // 设置最大打开的连接数
	m.db.SetMaxIdleConns(m.options.MaxIdleConn) //  设置最大空闲连接数
	db = m.db
	return
}
