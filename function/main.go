package main

import (
	"fmt"
	ui "function/utils"
)


// 自定义数据类型：给func(int, int) int取别名
type myFunc func(int, int) int

// 函数作为形参
func mySum(funcVar myFunc, n1 int, n2 int) int {
	return funcVar(n1, n2)
}

// 函数返回值命名
func getSum(n1 int, n2 int) (sum, sub int) {
	// 由于在返回值列表中定义过sum和sub的类型，这里可以直接使用这两个变量
	sub = n1 - n2
	sum = n1 + n2
	return	// no explicit return value
}

// 可变参数：可求出1到多个int的和
func sumMul(n1 int, args... int) (sum int) {
	sum = n1
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}
	return
}

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
// 为什么需要函数：如果要在多个地方使用到上面的功能？
	
	fmt.Printf("%f %c %f = %f", n1, operator, n2, res)

	var n3 int = 10
	var n4 int = 20
	sum, sub := ui.GetSumAndSub(n3, n4)
	fmt.Printf("\n%d + %d = %d\n%d - %d = %d\n", n3, n4, sum, n3, n4, sub)

	// 函数数据类型
	a := ui.GetSum
	fmt.Printf("a的数据类型为：%T, GetSum()的数据类型为：%T\n", a, ui.GetSum)
	ans := a(10, 30)
	fmt.Println("a(10, 30) =", ans)

	// 函数作为形参传入函数中调用
	ans = mySum(ui.GetSum, 10, 30)
	fmt.Println("mySum(10, 30) =", ans)

	// 返回值命名
	sum, sub = getSum(10, 20)
	fmt.Println("sum =", sum, "sub =", sub)
	
	// 可变参数
	fmt.Println(sumMul(1, 10, 100))

}