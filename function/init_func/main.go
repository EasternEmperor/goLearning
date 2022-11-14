package main

import (
	"fmt"
	// 如果在引入包中有init，则引入包的init会比本文件内的init更早执行
	ui "function/utils"
)

// 全局变量最先定义
var age = test()

func test() int {
	fmt.Println("test()...")
	return 90
}

// 通常可以在init函数中完成初始化工作
func init() {
	fmt.Println("init()...")
}

func main() {
	fmt.Println("main()...")
	fmt.Println("age = ", age)
	fmt.Println("Age =", ui.Age, "Name =", ui.Name)
}