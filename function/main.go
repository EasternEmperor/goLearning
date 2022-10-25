package main

import (
	"fmt"
	ui "function/utils"
)



func main() {

	// 输入两个数字和运算符"+, -, *, /"，完成运算
	var n1 float64 = 3.2
	var n2 float64 = 2.1
	var operator byte = '+'
	var res float64 = ui.Cal(n1, n2, operator)
	// switch operator {
	// case '+':
	// 	res = n1 + n2
	// case '-':
	// 	res = n1 - n2
	// case '*':
	// 	res = n1 * n2
	// case '/':
	// 	res = n1 / n2
	// default:
	// 	fmt.Println("错误运算符...")
	// }
	
	fmt.Printf("%f %c %f = %f", n1, operator, n2, res)
}
// 为什么需要函数：如果要在多个地方使用到上面的功能？