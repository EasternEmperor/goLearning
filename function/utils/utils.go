package utils

import "fmt"

// 注意若要能在其他地方使用该函数，则函数首字母要大写，表示为public
func Cal(n1 float64, n2 float64, operator byte) float64 {
	var res float64
	switch operator {
	case '+':
		res = n1 + n2
	case '-':
		res = n1 - n2
	case '*':
		res = n1 * n2
	case '/':
		res = n1 / n2
	default:
		fmt.Println("错误运算符...")
	}
	return res
}