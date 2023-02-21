package main

import (
	"fmt"
	"projects/mass_user_communication_system/client/process"
)

// 定义全局变量保存用户id和密码
var (
	userId int
	userPwd string
	userName string
)

func main() {
	// 接受用户选择
	var key int
	// 判断是否继续显示菜单
	var loop = true
	// 登录和注册遇到的错误
	var err error

	for loop {
		// 打印菜单
		fmt.Println("------------------1. 用户登录------------------")
		fmt.Println("------------------2. 注册新用户-----------------")
		fmt.Println("------------------3. 退出系统------------------")
		// 读取输入的选项
		fmt.Scanln(&key)
		switch key {

		case 1:
			// 处理登录
			fmt.Printf("请输入用户Id：")
			fmt.Scanln(&userId)
			fmt.Printf("请输入密码：")
			fmt.Scanln(&userPwd)
			// 创建UserProcess实例处理登录请求
			up := &process.UserProcess{}
			err = up.Login(userId, userPwd)

		case 2:
			// 处理注册
			fmt.Printf("请输入要注册的用户Id: ")
			fmt.Scanln(&userId)
			fmt.Printf("请输入密码: ")
			fmt.Scanln(&userPwd)
			fmt.Println("请输入用户名(nickname): ")
			fmt.Scanln(&userName)
			// 创建UserProcess实例处理注册请求
			up := &process.UserProcess{}
			err = up.Register(userId, userPwd, userName)

		case 3:
			// 退出系统
			fmt.Println("~~~~~~您已退出系统，下次见~~~~~~")
			loop = false

		default:
			fmt.Println("您的输入不符合规范，请重新输入！")
			
		}

		// 打印错误
		if err != nil {
			fmt.Println(err.Error())
		}

	}
	
}