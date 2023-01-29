package main

import (
	"fmt"
	"encoding/json"
)

type Cat struct {
	Name string
	Age int
	Fur string
}

type Dog struct {
	Name string
	Size []int			// 切片
	Ptr *int			// 指针
	M map[int]string	// map
}

func HasCat() {
	
	cats := make(map[string]Cat)
	cats["小白"] = Cat{"小白", 3, "白色"}
	cats["小花"] = Cat{"小花", 100, "花色"}
	
	var name string
	fmt.Scanln(&name)

	v, b := cats[name]
	if b {
		fmt.Println(v)
	} else {
		fmt.Println("未查到此猫信息！")
	}

}

type Monster struct {
	Name string `json:"name"`	// 后面这个就是tag，可以用于序列化为json串
	Age int `json:"age"`
	Skill string `json:"skill"`
}

func main() {

	// 问题：有两只猫，输入名字输出信息；没有则显示没有
	HasCat()

	// 结构体是值类型！
	tmp := 5
	var dog1 Dog
	var dog2 = dog1			// 赋值，注意是值类型！
	// 赋值
	dog1.Name = "White"
	dog1.Size = make([]int, 5)
	dog1.Size[0] = 100
	dog1.Ptr = &tmp
	dog1.M = make(map[int]string)
	dog1.M[10] = "You"
	// 输出
	fmt.Println(dog1)
	fmt.Println(dog2)		// 仍为空
	
	dog2 = dog1			// 再赋值，通过dog2来修改其中的切片、指针等引用类型字段
	fmt.Println(dog2)
	// 修改
	dog2.Name = "Black"
	tmp2 := 60
	dog2.Size[1] = 50
	dog2.Ptr = &tmp2
	dog2.M[7] = "Me"
	// 输出
	fmt.Println(dog1)		// dog1的Size, M也变化了
	fmt.Println(dog2)

	// tag举例
	var monster Monster
	monster.Name = "牛魔王"
	monster.Age = 1500
	monster.Skill = "吹牛"
	fmt.Println(monster)
	// 序列化为json串
	jsonStr, e := json.Marshal(monster)
	if e != nil {
		fmt.Println("转换出错！")
	} else {
		fmt.Println("jsonStr: ", string(jsonStr))	// 注意转换出来的为byte数组，要转为string再输出
	}


}