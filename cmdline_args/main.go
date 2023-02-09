package main

import (
	"fmt"
	"os"
	"flag"
)

func main() {

	fmt.Println("命令行参数有：")
	// 遍历os.Args，打印命令行参数
	for i, v := range os.Args {
		fmt.Printf("os.Args[%d] = %s\n", i, v)
	}

	// 以参数名接收命令行参数
	var user string
	var pwd string
	var port int
	var host string
	// 绑定参数名
	flag.StringVar(&user, "u", "", "用户名，默认为空")
	flag.StringVar(&pwd, "pwd", "", "密码，默认为空")
	flag.IntVar(&port, "p", 3366, "端口号，默认为3366")
	flag.StringVar(&host, "h", "localhost", "ip地址，默认为127.0.0.1")
	// flag.Parse()转换
	flag.Parse()
	// 打印参数
	fmt.Println("转换后：")
	fmt.Println("u =", user, "pwd =", pwd, "p =", port, "h =", host)

}