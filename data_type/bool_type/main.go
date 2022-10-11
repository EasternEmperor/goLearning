// bool: true false
// 1个字节，逻辑运算

package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var a = false
	fmt.Println("a =", a)

	// 占用一个字节
	fmt.Println("size of a: ", unsafe.Sizeof(a))
}