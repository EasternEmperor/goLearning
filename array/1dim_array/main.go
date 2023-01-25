package main

import (
	"fmt"
)

func test1(arr [3]int) {
	arr[0] = 100
}

func test2(arr *[3]int) {
	(*arr)[0] = 100
}

func main() {

	// 定义数组
	var hens [6]float64;
	// 给数组元素赋值，下标从0开始
	hens[0] = 3.0
	hens[1] = 5.0
	hens[2] = 1.0
	hens[3] = 2.0
	hens[4] = 12.5
	hens[5] = 4.2
	// 遍历数组求和
	totalweight := 0.0		// 注意数据类型
	for i := 0; i < len(hens); i++ {
		totalweight += hens[i]
	}
	fmt.Println(totalweight)

	// 初始化方法
	// 1.
	var arr1 [3]int = [3]int{1, 2, 3}
	fmt.Println("arr1 =", arr1)
	// 2.
	var arr2 = [3]int{4, 5, 6}
	fmt.Println("arr2 =", arr2)
	// 3.
	var arr3 = [...]int{7, 8, 9}
	fmt.Println("arr3 =", arr3)
	// 4. 指定元素下标
	var arr4 = [...]int{1: 10, 0: 11, 2: 12}
	fmt.Println("arr4 =", arr4)

	// for range 遍历
	for index, value := range arr1 {
		fmt.Printf("arr1[%d] = %d\n", index, value)
	}

	// 数组作为函数实参
	fmt.Println("arr1 =", arr1)
	test1(arr1)
	fmt.Println("arr1 =", arr1)
	// 传入数组指针
	fmt.Println("arr1 =", arr1)
	test2(&arr1)
	fmt.Println("arr1 =", arr1)

}