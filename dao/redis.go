package dao

import (
	"github.com/garyburd/redigo/redis"
	"log"
	"time"
)

var Pool *redis.Pool

func InitRedis(server, pwd string) {
	Pool = NewPool(server, pwd)
}

//redis 连接池
func NewPool(server, pwd string) *redis.Pool {

	return &redis.Pool{
		MaxIdle:     64,
		MaxActive:   1000,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			if pwd != "" {
				//密码
				if _, err := c.Do("AUTH", pwd); err != nil {
					_ = c.Close()
					return nil, err
				}
			}
			return c, nil
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				log.Println("Redis连接成功")
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}
}
