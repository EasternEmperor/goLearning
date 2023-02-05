package main

import (
	"fmt"
)

type Student struct {

}

func TypeJudge(items ...interface{}) {
	for idx, v := range items {
		switch v.(type) {
		case bool:
			fmt.Printf("第%d个元素的类型：bool\t值为：%v\t", idx + 1, v)
		case float64:
			fmt.Printf("第%d个元素的类型：float64\t值为：%v\t", idx + 1, v)
		case int, int64:
			fmt.Printf("第%d个元素的类型：int\t值为：%v\t", idx + 1, v)
		case nil:
			fmt.Printf("第%d个元素的类型：nil\t值为：%v\t", idx + 1, v)
		case string:
			fmt.Printf("第%d个元素的类型：string\t值为：%v\t", idx + 1, v)
		case Student:
			fmt.Printf("第%d个元素的类型：Student\t值为：%v\t", idx + 1, v)
		case *Student:
			fmt.Printf("第%d个元素的类型：*Student\t值为：%v\t", idx + 1, v)
		default:
			fmt.Printf("第%d个元素的类型：未知\t值为：%v\t", idx + 1, v)
		}
		// fmt.Printf("\n%T\n", v)		// %T也可以打印出准确类型
		fmt.Println()
	}
}

func main() {

	var i int = 10
	var b bool = false
	var f float64 = 3.14
	var s string = "China"
	var stu1 Student = Student{}
	var stu2 *Student = &Student{}
	var f2 float32 = 3.14

	TypeJudge(i, b, f, s, stu1, stu2, f2)

}