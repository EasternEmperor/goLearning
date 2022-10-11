package main

import "fmt"
import "unsafe"

func main() {
	// 整数类型后面的数字表示其包括的二进制数，从而可推定其范围
	var n int8 = 127
	fmt.Println("n =", n)

	// int根据操作系统位数，32位则默认32，64位默认64
	var a int = 8900
	var b uint = 0
	var c byte = 255		// byte为8位无符号，一般用于存放字符
	fmt.Println("a =", a, "b =", b, "c =", c)

	// 直接赋值为指定类型时，类型推导为int
	var num = 100
	// fmt.Printf("%T", num)		%T 可以输出对应变量的类型
	fmt.Printf("type of num: %T\n", num)
	fmt.Printf("type of c: %T", c)		// c声明为byte，但输出类型为uint8

	// 查看变量占用字节数：unsafe.Sizeof()
	fmt.Printf("\ntype of a: %T, memory usage of a: %d", a, unsafe.Sizeof(a))
}

// 使用时尽量使用占用空间小的适配数据类型。（保小不保大）