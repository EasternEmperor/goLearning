package main

import (
	"fmt"
	"net"
	"projects/mass_user_communication_system/server/model"
	"time"
)

// 处理连接
func mainProcess(conn net.Conn) {
	// 处理连接
	// 最终要关闭连接
	defer conn.Close()

	// 创建Processor实例来处理连接
	var processor = &Processor{conn}
	err := processor.serverProcessRead()
	if err != nil {
		fmt.Println("\n====>客户端和服务端通讯协程出错！<====")
		fmt.Println("====>", err)
	}
}

// 初始化myUserDAO
func initMyUserDAO() {
	// 全局pool赋值
	model.MyUserDAO = model.NewUserDAO(pool)
}

func main() {

	// 启动时初始化pool
	initPool(16, 0, 100 * time.Second, "localhost:6379")
	// 启动时初始化myUserDAO
	initMyUserDAO()

	// 开始监听
	fmt.Println("服务器在8889端口开始监听...")
	listen, err := net.Listen("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("监听出错！err =", err)
		return 
	}
	// 程序结束停止监听
	defer listen.Close()

	// 循环等待客户端连接
	for {
		fmt.Println("等待客户端来连接...")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("等待出错！err =", err)
		}
		// 处理连接
		go mainProcess(conn)
	}
}