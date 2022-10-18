package main

import "fmt"

func main() {
	// '/'：只保留整数部分
	fmt.Println(5 / 2)		// 2

	// 由于go基本数据类型的严格性，只有浮点型参与运算才会保留小数
	var f float32 = 10 / 4
	fmt.Println(f)		// 2
	f = 10.0 / 4
	fmt.Println(f)		// 2.5

	// 演示%
	// 公式：a % b = a - a / b * b		// 注意这里的 a / b 是只保留整数部分的
	fmt.Println(10 % 3)		// 1
	fmt.Println(-10 % 3)	// -1
	fmt.Println(10 % -3)	// 1
	fmt.Println(-10 % -3)	// -1

	// go的自增/减++/--只能放在变量后面，不能放在前面
	// 同时：++/--只能独立使用！如if (i++ < 0) 是不行的！
}