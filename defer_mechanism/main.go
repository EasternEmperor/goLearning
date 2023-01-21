package main

import (
	"fmt"
)

func sum(n1, n2 int) int {
	defer fmt.Println("n1 = ", n1)	// defer将该句压入一个栈中，暂时不执行，当函数执行完毕时弹出执行
	defer fmt.Println("n2 = ", n2)	// 所以会先输出"n2 = "
	// 增加一句话
	n1++	// 11
	n2++	// 21
	res := n1 + n2		// 32
	fmt.Println("res = ", res)
	return res
}

func main() {
	res := sum(10, 20)
	fmt.Println("res = ", res)		// 32
}