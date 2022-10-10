// 转义字符：'\'
// '\t'	制表符
// '\n'	换行符
// '\r'	回车

package main

import "fmt"

func main() {
	// \r：回车，从当前行的最前面开始输出，会覆盖之前的输出，如下：
	fmt.Println("Jack\rMa")

	// exercise
	fmt.Println("姓名\t年龄\t籍贯\t住址\nJohn\t12\t河北\t北京")
}