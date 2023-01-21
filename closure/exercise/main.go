/*
1. 函数makeSuffix(suffix string)，接收文件后缀名，返回一个闭包
2. 闭包可传入一个文件名，如果改文件名没有指定的后缀，如jpg，则返回 文件名.jpg ，如果已有该后缀则返回原文件名
3. 使用闭包
4. strings.HasSuffix(s, suffix string) bool		判断指定字符串s是否有后缀suffix
*/

package main
import (
	"fmt"
	"strings"
)

func makeSuffix(suffix string) func(string) string {
	return func(filename string) string {
		if strings.HasSuffix(filename, suffix) == true {
			return filename
		} else {
			return filename + suffix
		}
	}
}

func main() {
	f := makeSuffix(".jpg")
	fmt.Println(f("a.jpg"))
	fmt.Println(f("b"))
}