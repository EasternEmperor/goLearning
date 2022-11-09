package main

import "fmt"

// 递归函数
func test(n int) {
	if n > 2 {
		n--
		test(n)
	}
	fmt.Println(n)
}
func test2(n int) {
	if n > 2 {
		n--
		test2(n)
	} else {
		fmt.Println(n)
	}
}

func main() {

	// test(4)输出
	test(4)		// 2 2 3

	fmt.Println()

	// test2(4)输出
	test2(4)	// 2

}