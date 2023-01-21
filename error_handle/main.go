package main

import (
	"fmt"
	"strings"
	"errors"
)

func div() {
	// defer后执行一个匿名函数来捕获错误
	defer func() {
		err := recover()	// 捕获error
		if err != nil {
			// 捕获到error，可以输出
			fmt.Println("err =", err)
			// 对错误进行处理/通知
			fmt.Println("告知管理员......")
		}
	}()
	n1 := 10
	n2 := 0
	fmt.Println(n1 / n2)
	//fmt.Println(10 / 0)	这样无法捕获error
}

// 读取配置文件错误
func ReadConf(filename string) (err error) {
	if strings.HasSuffix(filename, ".conf") == false {
		// 错误文件，读取失败返回error
		return errors.New("文件错误！")
	} else {
		return nil
	}
}

// 调用ReadConf()
func test() {
	err := ReadConf("a.jpg")
	if err != nil {
		panic(err)
	}
	fmt.Println("继续执行......")
}

func main() {

	// 错误执行机制
	div()

	// 自定义错误
	test()

	fmt.Println("程序继续执行......")

}