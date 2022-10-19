package main

import "fmt"

func main() {
	// 1. 二进制：使用%b将十进制转为二进制输出
	var bnum int = 5
	fmt.Printf("binary num = %b\n", bnum)

	// 2. 八进制：在go中，以0开头的整型即为八进制
	var onum int = 011
	fmt.Printf("octonary num = %d\n", onum)

	// 3. 十六进制：以0x/0X开头的整型即为十六进制
	var hnum int = 0x11
	fmt.Printf("hexadecimal num = %d\n", hnum)
}