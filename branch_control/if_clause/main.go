package main

import "fmt"

func main() {
	
	// 1. 单分支

	// 输入年龄，如果大于18，则输出“你已成年，要对自己的行为负责！”
	var age int8
	fmt.Println("请输入你的年龄：")
	fmt.Scanln(&age)

	// 注意：Go中，if后面的花括号不能省略，不管从句中有多少条语句！
	// 注意：Go中，if后面的条件判断不能为赋值表达式！
	if age > 18 {
		fmt.Println("你已成年，要对自己的行为负责！")
	}
	// Go的特性：可以在if后面的条件判断中声明一个变量，该变量作用域仅在if从句中生效
	/* eg:
	if age := 20; age > 18 {
		fmt.Println("你已成年，要对自己的行为负责！")
	}
	*/

	// 2. 双分支

	// 注意Go中双分支的格式：else一定要在上一个if从句的花括号后写，否则报格式错误
	if age > 18 {
		fmt.Println("你已成年，要对自己的行为负责！")
	} else {
		fmt.Println("你未成年，未成年人保护法救你狗命！")
	}

	// 3. 多分支
	
	// if {} else if {} ... else {}
	var score int8
	fmt.Println("请输入你的分数：")
	fmt.Scanln(&score)
	if score == 100 {
		fmt.Println("奖励宝马一台！")
	} else if score > 80 {
		fmt.Println("奖励Iphone一台！")
	} else if score >= 60 {
		fmt.Println("奖励Ipad一台！")
	} else {
		fmt.Println("没有奖励！")
	}
}