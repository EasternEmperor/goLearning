package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num int = 100
	var f float32 = 3.14
	var b bool = true
	var c byte = 'a'
	var str string

	// 1. 用fmt.Sprintf()进行转换
	str = fmt.Sprintf("%d", num)
	fmt.Printf("str type %T, str = %s\n", str, str)
	str = fmt.Sprintf("%f", f)
	fmt.Printf("str type %T, str = %s\n", str, str)
	str = fmt.Sprintf("%t", b)
	fmt.Printf("str type %T, str = %s\n", str, str)
	str = fmt.Sprintf("%c", c)
	fmt.Printf("str type %T, str = %s\n", str, str)

	// 2. 用 strconv() 函数
	str = strconv.FormatInt(int64(num), 10)
	fmt.Printf("str type %T, str = %s\n", str, str)
	str = strconv.Itoa(int(num))
	fmt.Printf("str type %T, str = %s\n", str, str)
	str = strconv.FormatFloat(float64(f), 'f', -1, 32)
	fmt.Printf("str type %T, str = %s\n", str, str)
	str = strconv.FormatBool(b)
	fmt.Printf("str type %T, str = %s\n", str, str)
}