package main

import (
	"fmt"
)

func main() {

	var arr [5]int = [...]int{1, 2, 3, 4, 5}
	// 定义
	slice := arr[1 : 3]		// 和python类似，左闭右开
	fmt.Println("arr =", arr)
	fmt.Println("slice =", slice)
	fmt.Println("len(arr) =", len(arr))
	fmt.Println("len(slice) =", len(slice))
	fmt.Println("切片容量：cap(slice) =", cap(slice))

	// make
	var slice2 = make([]int, 4, 8)		// 默认元素0
	fmt.Println("slice2 =", slice2)
	fmt.Println("len(slice2) =", len(slice2))
	fmt.Println("cap(slice2) =", cap(slice2))
	slice2[0] = 100
	slice2[1] = 101
	fmt.Println("slice2 =", slice2)
	// 扩容？
	//slice2[4] = 102

	// 指定具体数组
	var slice3 = []string{"hello", "world", "!"}
	fmt.Println("slice3 =", slice3)
	fmt.Println("len(slice3) =", len(slice3))
	fmt.Println("cap(slice3) =", cap(slice3))

	fmt.Println("")

	// 内置函数append
	fmt.Printf("slice = %p\n", &slice)
	slice = append(slice, 6, 7, 8)
	fmt.Printf("slice = %p\n", &slice)		// 地址不变
	slice = append(slice, slice...)			// append切片注意三点...
	fmt.Printf("slice = %p\n", &slice)
	slice4 := append(slice, 9)
	fmt.Printf("slice4 = %p\n", &slice4)

	// 内置函数copy
	copy(slice2, slice)		// slice比slice2的len更长，会只将slice在slice2长度内的元素拷贝过去（截断）
	fmt.Println("slice2 =", slice2)
	fmt.Println("slice =", slice)

	// 修改字符串
	var str = "Hello World 中国！"
	strSlice1 := []rune(str)
	strSlice1[0] = '你'		// []rune按字符来处理字符串，所以中文可以用单引号
	str = string(strSlice1)
	fmt.Println(str)
	str = "Hello"
	strSlice2 := []byte(str)
	strSlice2[0] = 'H'
	str = string(strSlice2)
	fmt.Println(str)

}