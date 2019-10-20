package cache

import (
	"github.com/gomodule/redigo/redis"
	"time"
)

type RedisCli struct {
	host    string
	pool    *redis.Pool
	options *Options
}

func (rc *RedisCli) Init(opts ...Option) {
	rc.options = DefaultOpts()
	for _, opt := range opts {
		opt(rc.options)
	}
	rc.host = rc.options.Host + ":" + rc.options.Port
	rc.pool = &redis.Pool{
		MaxIdle:     64,                 // 连接池最大的空闲的数量
		MaxActive:   1000,               // 最大连接数
		IdleTimeout: rc.options.Timeout, // 连接超时时间
		Dial: func() (redis.Conn, error) { // 使用 dial 这个函数于redis 建立连接
			c, err := redis.Dial("tcp", rc.host)
			if err != nil {
				return nil, err
			}
			if rc.options.Password != "" {
				if _, err := c.Do("AUTH", rc.options.Password); err != nil {
					c.Close()
					return nil, err
				}
			}

			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error { // 验证连接的可用性
			// 一分钟 ping 一次
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}
}

func (rc *RedisCli) Conn() redis.Conn {
	return rc.pool.Get()
}
