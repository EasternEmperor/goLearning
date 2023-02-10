package main

import (
	"fmt"
	"net"
)

// 接收客户端发送的消息
func process(conn net.Conn) {
	// 最后要关闭连接
	defer conn.Close()

	for {
		var data = make([]byte, 1024)
		n, err := conn.Read(data)
		// 读取出错
		if err != nil {
			fmt.Println("===>客户端退出！，err =", err)
			break
		}
		// 打印data
		fmt.Printf("读取数据%d字节: %s\n", n, string(data[ : n]))	// 注意data是1024大小，但是数据未必有那么多
	}
}

func main() {
	fmt.Println("服务器开始监听...")
	// 监听
	listen, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Println("监听出错，err =", err)
		return 
	}
	// 最后要关闭连接
	defer listen.Close()

	// 监听成功则用listen循环接收
	for {
		// 等待客户端返回确认连接的消息（第三次握手）
		fmt.Println("等待客户端连接...")
		// Accept()会阻塞直到客户端回信
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("连接失败！err =", err)
		} else {
			fmt.Println("连接成功，conn =", conn, "客户端识别: ", conn.RemoteAddr().String())
		}
		// 这里可以写一个协程，为客户端服务
		go process(conn)
	}
}