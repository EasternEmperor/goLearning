package main
import (
	"github.com/garyburd/redigo/redis"
	"time"
)

// 定义一个全局的pool
var pool *redis.Pool

// 初始化操作
func initPool(maxIdle int, maxActive int, idleTimeout time.Duration, address string) {
	pool = &redis.Pool{
		MaxIdle: maxIdle,         	// 最大空闲连接数
		MaxActive: maxActive,       // 表示和数据库的最大连接数，0表示没有限制（则看MaxIdle）
		IdleTimeout: idleTimeout,   // 最大空闲时间，当一个连接在IdleTimeout内没有被使用，就会重新变为Idle
		// 初始化连接的代码，连接哪个ip的redis
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", address)
		},
	}
}