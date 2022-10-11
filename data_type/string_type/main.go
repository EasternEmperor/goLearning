package main

import "fmt"

func main() {
	var s string = "长城 110 hello world"
	fmt.Println("s =", s)

	// 1、在go中string类型一旦赋值就不能修改了，是不可变的
	//s[0] = 'a'		// 报错：cannot assign to s[0] (value of type byte)

	// 2、字符串的两种表示形式：(1)双引号；(2)反引号
	// 反引号会将其中内容认为是文本，不会解析
	s2 := `
package main
import "fmt"
	`
	fmt.Println(s2)

	// 3、用+对字符串进行拼接操作
	str := "hello"
	str2 := " world"
	str3 := str + str2 + " haha!"
	fmt.Println(str3)
	// 多个相加可以换行，但是'+'需在上一行
	str4 := str +
	str2 +
	str3
	fmt.Println(str4)
}