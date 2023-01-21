package main

import (
	"fmt"
)

// 累加器
func AddUpper() func(int) int {		// 返回值：一个函数func（闭包）
	var n int = 10
	return func(x int) int {
		n += x
		return n
	}
}

/**
这里返回了一个匿名函数，但是这个匿名函数引用了函数外的变量n，则该匿名函数和n就形成一个整体，构成闭包
可以将闭包理解成一个类，函数是操作，而引用到的变量是类的成员变量
*/

func main() {
	f := AddUpper()
	fmt.Println(f(1))	// 11
	fmt.Println(f(2))	// 13: 因为在闭包下，上面调用f(1)时改变了n的值
	fmt.Println(f(3))	// 16: 同理
}