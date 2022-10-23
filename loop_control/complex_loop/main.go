package main

import "fmt"

func main() {
	// 打印金字塔
	var layer int
	fmt.Println("输入金字塔层数：")
	fmt.Scanln(&layer)
	fmt.Println("简单金字塔：")
	for j := 0; j <= layer; j++ {
		for i := 0; i < j; i++ {
			fmt.Printf("*")
		}
		fmt.Println("")
	}
	fmt.Println("\n复杂金字塔：")
	for j := 0; j <= layer; j++ {
		// 打印空格
		for i := 0; i < layer - j; i++ {
			fmt.Printf(" ")
		}
		stars := 2 * (j - 1) + 1
		for k := 0; k < stars; k++ {
			fmt.Printf("*")
		}
		fmt.Println("")
	}

	// 九九乘法表
	fmt.Println("\n九九乘法表：")
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d\t", j, i, i * j)
		}
		fmt.Println("")
	}
}