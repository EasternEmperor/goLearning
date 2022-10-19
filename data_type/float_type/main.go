package main

import "fmt"

func main() {
	// float32单精度，float64双精度
	// 类型推导默认为float64
	var a = 1.1
	fmt.Printf("type of a: %T", a)

	// 科学计数法
	b := 5.1234e2
	fmt.Println("b =", b)
	fmt.Println("b =", b)		// shift alt ↓	复制行
	c := 512.34e-2
	fmt.Println("c =", c)

	var d float32 = 100
	fmt.Println(d)
}