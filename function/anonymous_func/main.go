package main

import "fmt"

var (
	// tips：全局变量不能用:=语法糖来初始化，应该直接使用等号=
	Func1 = func (n1, n2 int) int {
		return n1 * n2
	}
)

func main() {

	// 定义匿名函数时就调用，这种方式匿名函数只能调用一次
	res := func (n1, n2 int) int {
		return n1 + n2
	}(10, 20)
	fmt.Println("res =", res)

	// 匿名函数赋给变量，通过函数变量调用
	a := func (n1, n2 int) int {
		return n1 - n2
	}
	res = a(10, 20)
	fmt.Println("res =", res)
	res = a(20, 10)
	fmt.Println("res =", res)

	// 全局匿名函数
	res = Func1(10, 20)
	fmt.Println("res =", res)

}