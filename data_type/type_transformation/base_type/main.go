package main

import "fmt"

func main() {

	// int -> float
	var i int32 = 100
	//var i2 int64 = i		// 同样不能直接转换！
	var f float32 = float32(i)
	fmt.Printf("i = %d, f = %f\n", i, f)
	// 结论：在go中，只要类型不同就要显示转换，包括int8、int32....
	// 同时：从高精度->低精度，float有精度丢失情况，int可能有溢出情况（按溢出处理）
	var i3 int64 = 10000
	var i4 int8 = int8(i3)		// i4 = 16
	fmt.Printf("i4 = %d\n", i4)

	// 经典错误：
	//var i5 int8 = int8(i) + 128		// 128 (untyped int constant) overflows int8
	// 注意有些constant数字实际上是暗含了其类型的，如128不可能是int8类型
}