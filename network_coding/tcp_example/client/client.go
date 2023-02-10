package main

import (
	"fmt"
	"net"
	"os"
	"bufio"
	"strings"
)

func main() {
	// 连接服务器
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("连接出错！err =", err)
		return 
	}
	// 最后要关闭连接
	defer conn.Close()

	// 从键盘向服务器发消息
	for {
		// 读取键盘输入
		reader := bufio.NewReader(os.Stdin)
		data, err := reader.ReadString('\n')
		data = data[ : len(data) - 2]
		// 读取出错
		if err != nil {
			fmt.Println("读取键盘输入出错！err =", err)
		}
		// 输入exit则退出客户端
		if strings.ToLower(data) == "exit" {
			break
		}
		// 传输该串
		n, err2 := conn.Write([]byte(data))
		if err != nil {
			fmt.Println("传输数据出错！err =", err2)
		}
		fmt.Printf("传输了%d个字节：%s\n", n, data)
	}

	fmt.Println("你已退出客户端！下次见~")

}