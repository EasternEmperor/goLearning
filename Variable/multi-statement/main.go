package main

import "fmt"

// 二、定义全局变量
var m1, m10 = 100, 200
var m2 = 200
var m3 = "Jack"
// 或者用括号
var (
	m4 = 200
	m5 = "Mary"
)

func main() {
	// 一、定义多个变量的三种方式：
	// 1、一次定义多个变量
	var n1, n2, n3 int
	fmt.Println("n1 =", n1, "n2 =", n2, "n3 =", n3)

	// 2、定义并赋值
	var n, name, num = 10.1, "Tom", 100
	fmt.Println("n =", n, "name =", name, "num =", num)

	// 3、类型推导，使用:=
	n4, name2, n5 := 10.1, "Tom", 100
	fmt.Println("n4 =", n4, "name2 =", name2, "n5 =", n5)

	// 输出全局变量
	fmt.Println("m1 =", m1, "m2 =", m2, "m3 =", m3, "m10 =", m10)
	fmt.Println("m4 =", m4, "m5 =", m5)
	
}