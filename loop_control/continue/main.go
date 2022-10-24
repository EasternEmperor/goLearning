package main

import "fmt"

func main() {

	// 使用标签
	here:
	for i := 0; i < 2; i++ {
		for j := 0; j < 4; j++ {
			if j == 2 {
				continue here;
			}
			fmt.Println("i =", i, "j =", j)
		}
	}

	// exec: 从键盘读入不确定个数的整数，以0结束，输出正数和负数的个数
	positive, negative := 0, 0
	for {
		fmt.Println("输入整数：")
		var num int
		fmt.Scanln(&num)
		if (num == 0) {
			break
		} else if num > 0 {
			positive++
		} else {
			negative++
		}
	}
	fmt.Printf("正数有%d个，负数有%d个", positive, negative)

	// 某人有cash元，每经过一次一次路口，需要缴费，规则如下：
	// 1. 现金>50000时，缴费5%
	// 2. 现金<=50000时，缴费1000
	// 输出此人能经过的路口数
	var cash float64 = 100000
	var times = 0
	for {
		if cash < 1000 {
			break
		}
		if cash > 50000 {
			cash -= cash * 0.05
			times++
		} else {
			cash -= 1000
			times++
		}
	}
	fmt.Println("\n\n此人能经过", times, "个路口")

}