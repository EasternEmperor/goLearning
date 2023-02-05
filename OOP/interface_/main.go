package main

import (
	"fmt"
)

// 定义一个接口
type Usb interface {
	// 声明两个没有实现的方法
	Start()
	Stop()
}

// 实现接口，即结构体的方法中有与接口声明的方法同名的即为实现接口
type Phone struct {

}
func (phone Phone) Start() {
	fmt.Println("手机开始工作......")
}
func (phone Phone) Stop() {
	fmt.Println("手机停止工作......")
}

// 实现接口
type Camera struct {

}
func (c Camera) Start() {
	fmt.Println("相机开始工作......")
}
func (c Camera) Stop() {
	fmt.Println("相机停止工作......")
}

// 展示多态
type Computer struct {

}
func (computer *Computer) Working(usb Usb) {
	// 接口多态
	usb.Start()
	usb.Stop()
}


// 只要是自定义类型就可以实现接口
type AInterface interface {
	say()
}
type Integer int
func (i Integer) say() {
	fmt.Println("Integer =", i, "say hi")
}

func main() {

	// 创建Phone和Camera变量
	var phone Phone
	var camera Camera
	// 创建Computer变量，调用Working方法来展示多态
	var computer Computer
	// 使phone工作
	computer.Working(phone)
	// 使camera工作
	computer.Working(camera)


	// 非结构体的自定义类型实现接口
	var i Integer = 10
	i.say()

	// 空接口类型可以指向任何变量
	var i2 interface{}
	var num float64 = 50
	i2 = num
	fmt.Println(i2)		// 50

}