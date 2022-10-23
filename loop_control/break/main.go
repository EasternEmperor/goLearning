package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// 生成0-100之间的随机数，当生成99时停止并报告次数
	rand.Seed(time.Now().UnixNano())

	var times int = 0
	for {
		rd := rand.Intn(100)
		times++
		fmt.Println(rd)
		if rd == 99 {
			break
		}
	}
	fmt.Printf("共用%d次生成99", times)

	// 使用label
	fmt.Println()
	label1:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			fmt.Printf("i = %v, j = %v\n", i, j)
			if j == 2 {
				break label1	// 当j为2时终止label1标签的for循环
			}
		}
	}
}