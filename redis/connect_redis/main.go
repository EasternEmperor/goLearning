package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func Batch(conn redis.Conn) {
	// 批量写入Hash数据: HMSet
	_, err := conn.Do("HMSet", "user01", "name", "Tom", "age", "3")
	if err != nil {
		fmt.Println("HMSet err =", err)
	}

	// 批量读取: HMGet
	val, err := redis.Strings(conn.Do("HMGet", "user01", "name", "age"))
	if err != nil {
		fmt.Println("HMGet err =", err)
	}
	// 遍历输出
	for _, val := range val {
		fmt.Println(val)
	}
}

func main() {
	// 连接本地redis数据库
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("conn error =", err)
		return 
	}
	// 记得关闭
	defer conn.Close()

	// 写入String型数据
	_, err = conn.Do("set", "name", "TomJerry")
	if err != nil {
		fmt.Println("写入出错！err =", err)
	}
	// 读取
	val, err := redis.String(conn.Do("Get", "name"))
	if err != nil {
		fmt.Println("读取出错！err =", err)
	}

	fmt.Println("conn success, data =", val)

	// 批量写入和读出
	Batch(conn)
}