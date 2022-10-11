package main
import "fmt"


// go变量使用方法：
// 1、指定变量类型，不赋初值则为默认值(int 0 ...)
// 2、根据值自行判断变量类型（类型推导）
// 3、省略var直接声明并赋值变量，注意：:= 左边的变量不能是已声明过的，否则会导致编译错误

func main() {
	// 定义变量
	var num1 intj
	// 变量赋值
	num1 = 10
	// 使用变量
	fmt.Println("num1 = ", num1)

	var num2 = 10.1
	fmt.Println("num2 = ", num2)

	name := "Tom"
	fmt.Println("name = ", name)
}
