package main

import (
	"fmt"
	"encoding/json"
)

type Monster struct {
	Name string		`json:"name"`		// tag：使序列化时字段名变为指定的
	Age int			`json:"age"`		// 注意！！！json:后不能有空格！
	Sal int			`json:"sal"`
	Skill string	`json:"skill"`
}

// 对结构体序列化
func TestStruct() {
	// 创建结构体变量
	var m Monster = Monster{"牛魔王", 5000, 8000, "炎拳"}
	val, err := json.Marshal(m)
	if err != nil {
		fmt.Println("转换失败！err =", err)
	}
	fmt.Printf("转换后的json为：%v\n", string(val))
}
// 对map序列化
func TestMap() {
	// 创建map
	var m map[string]string = make(map[string]string)
	m["name"] = "Tom"
	m["age"] = "12"
	m["address"] = "Mexico"

	// 序列化
	val, err := json.Marshal(m)
	if err != nil {
		fmt.Println("序列化失败！err =", err)
	}
	fmt.Println("map序列化后：", string(val))
}
// 对切片序列化
func TestSlice() {
	// 创建切片
	var s []map[string]string
	m := make(map[string]string)
	m["name"] = "Mary"
	m["age"] = "20"

	m2 := make(map[string]string)
	m2["name"] = "Ultraman"
	m2["age"] = "20000"
	s = append(s, m, m2)

	// 序列化
	val, err := json.Marshal(s)
	if err != nil {
		fmt.Println("序列化失败！err =", err)
	}
	fmt.Println("slice序列化后：", string(val))

	// 普通切片
	var s2 []int
	s2 = make([]int, 10)
	s2[0] = 100
	s2[1] = 101

	// 序列化
	val2, err2 := json.Marshal(s2)
	if err != nil {
		fmt.Println("序列化失败！err =", err2)
	}
	fmt.Println("slice2序列化后：", string(val2))	// [100,101,0,0,0,0,0,0,0,0]
}

func main() {
	// 结构体序列化
	TestStruct()
	// map序列化
	TestMap()
	// slice序列化
	TestSlice()
}