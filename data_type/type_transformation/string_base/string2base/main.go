package main

import (
	"fmt"
	"strconv"
)

func main() {
	var s1 string = "100"
	var s2 string = "3.14"
	var s3 string = "true"

	// 使用strconv包内函数来转换
	var n1 int64
	n1, _ = strconv.ParseInt(s1, 10, 8)
	fmt.Printf("n1 type %T, n1 = %v\n", n1, n1)
	var f float64
	f, _ = strconv.ParseFloat(s2, 32)
	fmt.Printf("f type %T, f = %v\n", f, f)
	var b bool
	b, _ = strconv.ParseBool(s3)
	fmt.Printf("b type %T, b = %v\n", b, b)
	var n2 int
	n2, _ = strconv.Atoi(s1)
	fmt.Printf("n2 type %T, n2 = %v\n", n2, n2)
	
	// 传入string不对应
	var n3 int64 = 1
	n3, _ = strconv.ParseInt(s2, 10, 8)
	fmt.Printf("n3 type %T, n3 = %v", n3, n3)
}