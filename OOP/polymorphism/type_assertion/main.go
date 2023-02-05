package main

import (
	"fmt"
)

type Point struct {
	x float64
	y float64
}

func main() {

	// 类型断言
	var p Point = Point{1, 2}
	var a interface{}		// 空接口
	a = p		// 空接口可以指向任何变量
	var p2 Point
	// 要将a赋值给p2，要经过类型断言
	p2 = a.(Point)
	fmt.Println(p2)
	// 输出类型
	// 注意：a输出的也是Point类型，说明与声明时的类型无关，由指向的变量类型决定
	fmt.Printf("p的类型：%T\ta的类型：%T\tp2的类型：%T\n", p, a, p2)	

	// 类型断言（带检测的）
	var f1 float64 = 3.14
	var b interface{}
	b = f1
	// 加上一个检测结果用于判断转换成功与否
	f2, ok := b.(float32)
	if ok {
		fmt.Println("转换成功~~")
		fmt.Printf("f2类型：%T\n", f2)
	} else {
		fmt.Println("转换失败！！！")
	}
	fmt.Println("程序继续执行......")

}