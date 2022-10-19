package main

import "fmt"

func main() {
	
	// 要求：从控制台接受用户信息，【姓名，年龄，薪水，是否通过考试】
	// 1. 使用Scanln()解决
	var name string
	var age int
	var salary float32
	var exam bool
	fmt.Println("请输入用户信息：")
	fmt.Scanln(&name, &age, &salary, &exam)
	fmt.Printf("用户姓名：%s\n年龄：%d\n薪水：%.2f\n是否通过考试：%t\n\n", name, age, salary, exam)

	// 2. 使用Scanf()解决
	var name2 string
	var age2 int
	var salary2 float32
	var exam2 bool
	fmt.Println("请输入用户信息：")
	fmt.Scanf("%s%d%f%t", &name2, &age2, &salary2, &exam2)
	fmt.Printf("用户姓名：%s\n年龄：%d\n薪水：%.2f\n是否通过考试：%t\n", name2, age2, salary2, exam2)
}