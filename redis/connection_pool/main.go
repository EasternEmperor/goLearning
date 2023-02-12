package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

// 声明为全局变量
var pool *redis.Pool

// 启动时对pool初始化
func init() {
	pool = &redis.Pool{
		MaxIdle: 8,         // 最大空闲连接数
		MaxActive: 0,       // 表示和数据库的最大连接数，0表示没有限制（则看MaxIdle）
		IdleTimeout: 100,   // 最大空闲时间，当一个连接在IdleTimeout内没有被使用，就会重新变为Idle
		// 初始化连接的代码，连接哪个ip的redis
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "localhost:6379")
		},
	}
}

func main() {

	// 从pool取出一个连接
	conn := pool.Get()
	// 记得关闭
	// defer会在return之后启动操作
	// 多个defer是先进后出启动的
	defer conn.Close()
	defer pool.Close()

	// 写入
	_, err := conn.Do("set", "name", "汤姆猫~")
	if err != nil {
		fmt.Println("写入失败！err =", err)
		return 
	}

	// 读取
	val, err := redis.String(conn.Do("get", "name"))
	if err != nil {
		fmt.Println("读取失败！err =", err)
		return
	}
	fmt.Println("读取到: ", val)
}