package main

import (
	"fmt"
)

type Stu struct {
	Name string
	Age int
	Address string
}


func main() {

	// 使用方式
	// 1. 先声明再make分配空间
	var a map[string]string
	a = make(map[string]string)
	a["you"] = "spiderman"
	a["me"] = "ironman"
	a["he"] = "kola"
	fmt.Println(a)

	// 2. 声明时make分配空间
	m := make(map[string]string)
	m["she"] = "woman"
	fmt.Println(m)

	// 3. 声明时赋值
	m2 := map[string]string{
		"hero1" : "wuyong",
		"hero2" : "lujunyi",		// 注意：这里最后一个元素后也要有逗号！
	}
	fmt.Println(m2)

	// 增
	a["it"] = "hork"
	// 改
	a["you"] = "gundam"
	// 删：delete(m, key)，存在对应key则删除，否则不操作也不报错
	delete(a, "it")
	delete(a, "it")		// 不操作，不报错
	fmt.Println(a)
	// 查
	value, ok := a["it"]
	if ok {
		fmt.Println("a[it] =", value)
	} else {
		fmt.Println("a[it] 不存在")
	}

	fmt.Println()

	// 遍历for range
	for k, v := range a {
		fmt.Printf("a[%s] = %s\n", k, v)
	}

	// map切片
	monsters := make([]map[string]string, 2)	// 声明时同时分配空间给map切片
	// 添加元素
	if monsters[0] == nil {
		// 首先需要给map元素分配空间
		monsters[0] = make(map[string]string)
		monsters[0]["name"] = "牛魔王"
		monsters[0]["age"] = "1000"
	}
	if monsters[1] == nil {
		monsters[1] = make(map[string]string)
		monsters[1]["name"] = "红孩儿"
		monsters[1]["age"] = "500"
	}
	// 超过预分配的内存，则使用append来添加map元素
	NewMonster := make(map[string]string)
	NewMonster["name"] = "葫芦娃"
	NewMonster["age"] = "10"
	// append
	monsters = append(monsters, NewMonster)

	for i := 0; i < len(monsters); i++ {
		fmt.Printf("monsters[%d]:\n", i)
		for k, v := range monsters[i] {
			fmt.Printf("%s: %s\t", k, v)
		}
		fmt.Println()
	}

	// map值为struct
	var students = make(map[string]Stu, 10)
	stu1 := Stu{"Tom", 18, "北京"}
	stu2 := Stu{"Mary", 20, "上海"}
	students["no1"] = stu1
	students["no2"] = stu2
	
	fmt.Println(students)


}