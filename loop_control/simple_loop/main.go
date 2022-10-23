package main

import "fmt"

func main() {

	// for循环基本语法
	for i := 0; i <= 10; i++ {
		fmt.Println("hello world!")
	}

	// 第二种写法：和C/C++一样，用封号隔开可以省略初始化、循环条件和迭代
	j := 0
	for ; j <= 10; j++ {
		fmt.Println("hello world!")
	}

	// 第三种写法：只跟循环条件，类似于while了
	for j <= 15 {
		fmt.Println("Hello world!")
		j++
	}

	// 第四种写法：for后不跟任何语句，相当于死循环，配合break
	for {
		fmt.Println("Hello World!")
		break;
	}

	// 第五种写法：for range:
	var str = "hello world!"
	for index, ch := range str {
		fmt.Printf("index = %d, char = %c\n", index, ch)
	}

	// 由于传统的下标遍历字符串以字节为单位，无法以下标输出汉字这种3字节单位
	// 可以将string转成切片类型，再按下标遍历即可
	str = "Hello 北京"
	str2 := []rune(str)
	for i := 0; i < len(str2); i++ {
		fmt.Printf("%c\n", str2[i])
	}
	// 但是for range方式可以直接输出汉字
	for index, value := range(str) {
		fmt.Printf("index = %d, value = %c\n", index, value)
	}
}