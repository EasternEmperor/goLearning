package main

import "fmt"

func main() {
	// 位运算演示
	fmt.Println(2 & 3)	// 2
	fmt.Println(2 | 3)	// 3
	fmt.Println(2 ^ 3)	// 1
	fmt.Println(2 ^ -2)	// -4

	var num int8 = 64
	fmt.Println(num << 1)	// -128, 左移溢出
}