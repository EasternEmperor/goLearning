package main

import (
	"fmt"
	"reflect"
)

// 演示反射
func testReflect(i interface{}) {

	// 获取reflect.Type
	rTyp := reflect.TypeOf(i)	
	fmt.Printf("rTyp = %v, type(rTyp) = %T\n", rTyp, rTyp)

	// 获取reflect.Value
	rVal := reflect.ValueOf(i)
	// Value可以通过方法来获取内部值，此时可进行运算
	tmp := rVal.Int()
	fmt.Println("tmp = rVal.Int() =", tmp, "\ntmp + 2 =", tmp +2 )
	fmt.Printf("rVal = %v, type(rVal) = %T\n", rVal, rVal)

	// 转为interface{}然后再转回int变量
	i2 := rVal.Interface()
	num := i2.(int)			// 类型断言
	fmt.Println("num =", num)

}

type Student struct {
	Name string
	ID string
	Grade int
}
type Monster struct {
	Name string
	Skill string
}
// 对结构体用反射
func reflectStruct(i interface{}) {
	// 获取Type
	rTyp := reflect.TypeOf(i)
	fmt.Printf("rTyp = %v, type(rTyp) = %T\n", rTyp, rTyp)

	// 获取Value
	rVal := reflect.ValueOf(i)
	fmt.Printf("rVal = %v, type(rVal) = %T\n", rVal, rVal)

	// 反射还原为变量
	switch i.(type) {
	case Student:
		// 学生变量
		stu := i.(Student)
		fmt.Printf("学生：\n姓名：%s\t成绩：%d\n", stu.Name, stu.Grade)
	case Monster:
		// 怪兽变量
		mon := i.(Monster)
		fmt.Printf("怪兽：\n学名：%s\t必杀技：%s\n", mon.Name, mon.Skill)
	default:
		fmt.Println("识别不出来的类型~")
	}

}

func main() {
	
	// 演示变量, interface{}, reflect.ValueOf()
	var num = 10
	testReflect(num)

	// 反射结构体
	var s = Student{"Tom", "114514", 19}
	var m = Monster{"牛魔王", "炎拳"}
	reflectStruct(s)
	reflectStruct(m)
	reflectStruct(num)
	
}