package main

import (
	"fmt"
	"reflect"
)

// 传入指针来修改值
func ValueSet(i interface{}) {
	// 转为reflect.Value
	v := reflect.ValueOf(i)
	fmt.Printf("v = %v, type(v) = %T\n", v, v)

	// v.Elem().SetXxx(新值)
	v.Elem().SetInt(20)
}

func main() {
	var num = 100
	ValueSet(&num)
	fmt.Println("num =", num)
	// var num2 = 3.14
	// ValueSet(&num2)	// num2是float64类型
}