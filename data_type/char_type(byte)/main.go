// go没有专门的字符类型，要存储单个字符，用byte保存：var ch byte = 'a'
// go的字符串是由单个字节连接起来的

package main

import "fmt"

func main() {

	// 定义单个字符
	var a byte = 'a'
	var b byte = '0'

	// 直接输出为字符的码值
	fmt.Println("a =", a, "b =", b)
	// 要输出字符需格式化输出
	fmt.Printf("a = %c, b = %c\n", a, b)

	// 汉字码值超过byte，可用int保存
	var c int = '钟'
	fmt.Printf("c = %c\n", c)

	// 给变量赋数字，%c格式化输出变为对应的unicode编码字符
	// 字符类型进行加减运算，实际为对应码值参与运算
}