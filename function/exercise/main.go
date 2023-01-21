/*
使用函数
1. 打印金字塔：输入数字，打印对应层数金字塔
2. 打印乘法口诀表：输入数字，打印到对应数字的乘法口诀表
*/

package main

import "fmt"

func pyramid(layer int) {
	cnt := 1
	for temp := layer; temp > 0; temp-- {
		for i := temp - 1; i > 0; i-- {
			fmt.Printf(" ")
		}
		for i := 0; i < 2 * (cnt - 1) + 1; i++ {
			fmt.Printf("*")
		}
		fmt.Println()
		cnt++
	}
}

func MT(num int) {
	for i := 1; i <= num; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d×%d=%d\t", j, i, i * j)
		}
		fmt.Println()
	}
}

func main() {
	var layer int
	fmt.Scanln(&layer)

	pyramid(layer)

	var num int
	fmt.Scanln(&num)

	MT(num)
}