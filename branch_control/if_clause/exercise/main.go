package main

import "fmt"

func main() {
	// 1. 声明两个int32变量并赋值，两数和大于50打印"hello world!"
	var num1, num2 float32
	fmt.Println("请输入两个数：")
	fmt.Scanln(&num1, &num2)
	if num1 + num2 > 50 {
		fmt.Println("hello world!")
	}

	// 2. 若第一个数大于10，第二个数小于20，打印两数和
	if num1 > 10 && num2 < 20 {
		fmt.Println("两数和为：", num1 + num2)
	}

	// 3. 如果两数和能同时被3和5整除，打印相应信息
	if sum := int32(num1 + num2); sum % 3 == 0 && sum % 5 == 0 {
		fmt.Println("该两数和既能被3整除，也能被5整除！")
	} 

	// 4. 判断给定年份是否为闰年
	var year int
	fmt.Println("请输入一个年份：")
	fmt.Scanln(&year)
	if year % 400 == 0 || (year % 4 == 0 && year % 100 != 0) {
		fmt.Println(year, "是闰年！")
	} else {
		fmt.Println(year, "不是闰年！")
	}
}