package main
import "fmt"

func main() {
	// 变量类型
	var i int = 1
	fmt.Printf("i的地址是：%v\n", &i)

	// 指针类型
	var ptr *int = &i
	fmt.Printf("ptr的值：%v\n", ptr)
	fmt.Printf("ptr的地址：%v\n", &ptr)
	fmt.Printf("ptr指向值：%d\n", *ptr)

	// 获取int变量地址，并显示；
	// 通过指针修改该变量值
	var num int = 10
	fmt.Printf("num = %d\n", num)
	var ptr2 *int = &num
	*ptr2 = 199
	fmt.Printf("num的地址：%v\n", ptr2)
	fmt.Printf("After modified, num = %d\n", *ptr2)
}