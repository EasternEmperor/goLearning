package main

import "fmt"

func main() {
	// switch入门案例：接收字符a-g，a输出星期一……g输出星期日。
	fmt.Println("输入a-g之间的字母：")
	var ch byte
	fmt.Scanf("%c", &ch)
	switch ch {
	case 'a':
		fmt.Println("星期一")
	case 'b':
		fmt.Println("星期二")
	case 'c':
		fmt.Println("星期三")
	case 'd':
		fmt.Println("星期四")
	case 'e':
		fmt.Println("星期五")
	case 'f':
		fmt.Println("星期六")
	case 'g':
		fmt.Println("星期日")
	default:
		fmt.Println("错误输入")
	}


	// switch后面可以不带表达式，将switch语句作为if语句使用
	var score int = 80
	switch {
	case score > 80 && score <= 100:
		fmt.Println("A")
	case score >= 60:
		fmt.Println("B")
	default:
		fmt.Println("C")
	}

	// switch 穿透：fallthrough
	var num int = 10
	switch num {
	case 10:
		fmt.Println("ok1")
		fallthrough
	case 20:
		fmt.Println("ok2")		// 虽然num不为20，由于fallthrough仍会输出ok2
	case 30:
		fmt.Println("ok3")
	default:
		fmt.Println("nope")
	}
}