package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// 1. 生成5个伪随机数，并将其反转打印
	// rand.Intn(n int): 生成[0, n)之间的随机数
	// rand.Seed(n): 给定随机数种子
	var arr [5]int
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(100)
	}
	// 反转
	fmt.Println("反转前：", arr)
	temp := 0		// 临时变量
	for i := 0; i < len(arr) / 2; i++ {
		temp = arr[len(arr) - i - 1]
		arr[len(arr) - i - 1] = arr[i]
		arr[i] = temp
	}
	fmt.Println("反转后：", arr)

}