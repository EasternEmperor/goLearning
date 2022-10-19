package main

import "fmt"

// execise
func main() {

	// 1、交换a、b，但不使用中间变量
	// 解：使用加法+减法来完成
	var a int = 10
	var b int = 20
	a = a + b
	b = a - b
	a = a - b
	fmt.Printf("a = %d, b = %d\n", a, b)
}