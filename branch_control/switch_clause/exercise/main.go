package main

import "fmt"

func main() {
	// 1. 把用户输入字母a-e转换为大写，其他字符则输出"other"
	fmt.Println("输入一个a-e字母：")
	var ch byte
	fmt.Scan(&ch)
	switch ch {
	case 'a':
		fmt.Println("A")
	case 'b':
		fmt.Println("B")
	case 'c':
		fmt.Println("C")
	case 'd':
		fmt.Println("D")
	case 'e':
		fmt.Println("E")
	default:
		fmt.Println("other")
	}

	// 2. 对学生成绩大于60，输出“合格”；低于60，“不合格”；分数不能大于100
	var score int = 80
	switch {
	case score > 100 || score < 0:
		fmt.Println("错误输入")
	case score > 60:
		fmt.Println("合格")
	default:
		fmt.Println("不合格")
	}

	// 根据用户输入月份输出季节
	var month int
	fmt.Println("请输入月份：")
	fmt.Scan(&month)
	switch month {
	case 3, 4, 5:
		fmt.Println("春季")
	case 6, 7, 8:
		fmt.Println("夏季")
	case 9, 10, 11:
		fmt.Println("秋季")
	case 12, 1, 2:
		fmt.Println("冬季")
	default:
		fmt.Println("错误输入")
	}
}
// 由于windows换行为\n\r的形式，所以使用Scan读取输入会好一点，和C/C++中的scanf一样