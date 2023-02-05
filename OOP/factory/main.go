package main

import (
	"fmt"
	"OOP/factory/model"
)

func main() {

	// 通过工厂模式创建model包内的私有结构体student
	stu := model.NewStudent("Tom", 10)
	fmt.Println(stu)
	// 通过工厂模式来访问结构体student的私有字段age
	fmt.Println("stu.Name =", stu.Name, "\nstu.age =", stu.GetAge())

}