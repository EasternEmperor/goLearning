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

// 反序列化为结构体
func TestStruct() {
	// 待反序列化的字符串
	str := "{\"name\":\"牛魔王\",\"age\":5000,\"sal\":8000,\"skill\":\"炎拳\"}"

	// 序列化后保存到的变量
	var m Monster
	// 反序列化
	err := json.Unmarshal([]byte(str), &m)
	if err != nil {
		fmt.Println("反序列化失败！err =", err)
		return 
	}
	// 输出
	fmt.Println("反序列化得到：", m, "\n")
}

// 反序列化为map
func TestMap() {
	// 待反序列化的字符串
	str := "{\"name\":\"牛魔王\",\"age\":5000,\"sal\":8000,\"skill\":\"炎拳\"}"

	// 序列化后保存到的变量
	var m map[string]interface{}
	// 反序列化；m不需要make分配空间，make操作被封装到Unmarshal()内部了
	err := json.Unmarshal([]byte(str), &m)
	if err != nil {
		fmt.Println("反序列化失败！err =", err)
		return 
	}
	// 输出
	fmt.Println("反序列化得到：", m, "\n")
}

// 反序列化为slice
func TestSlice() {
	// 待反序列化的字符串
	str := "[{\"age\":\"20\",\"name\":\"Mary\"},{\"age\":\"20000\",\"name\":\"Ultraman\"}]"

	// 序列化后保存到的变量
	var s []map[string]interface{}
	// 反序列化；make操作被封装到Unmarshal()内部了
	err := json.Unmarshal([]byte(str), &s)
	if err != nil {
		fmt.Println("反序列化失败！err =", err)
		return 
	}
	// 输出
	fmt.Println("反序列化得到：", s, "\n")

	// 普通数组
	str2 := "[100,101,0,0,0,0,0,0,0,0]"
	var s2 []int
	// 反序列化；make操作被封装到Unmarshal()内部了
	err2 := json.Unmarshal([]byte(str2), &s2)
	if err != nil {
		fmt.Println("二次反序列化失败！err2 =", err2)
		return 
	}
	// 输出
	fmt.Println("二次反序列化得到：", s2, "\n")
}

func main() {
	// 反序列化为结构体
	TestStruct()
	// 反序列化为map
	TestMap()
	// 反序列化为slce
	TestSlice()
}