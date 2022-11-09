package main

import "fmt"

// 斐波那契数列
func feibo(n int) int {
	if n == 0 || n == 1 {
		return n
	} else {
		return feibo(n - 1) + feibo(n - 2)
	}
}

// 猴子吃桃子：一堆桃子，猴子每天吃一半+1，到第十天只剩一个桃子。问最初多少个桃子
func peaches(n int) int {
	if n > 1 {
		n--
		return 2 * (peaches(n) + 1)
	}
	return 1
}

func main() {

	var n int
	fmt.Scanf("%d", &n)
	res := feibo(n)
	fmt.Printf("feibo(%d) = %d\n", n, res)

	sum := peaches(10)
	fmt.Println("sum of peaches:", sum)
}