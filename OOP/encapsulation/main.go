package main

import (
	"fmt"
	"OOP/encapsulation/model"
)

func main() {

	// 创建Account变量
	var a model.Account
	
	// 设置账号
	// 失败: 报错导致程序终止
	// err := a.SetAccountNo("8888")
	// if err != nil {
	// 	panic(err)
	// } else {
	// 	fmt.Println("账号设置成功！")
	// }
	// 成功
	err := a.SetAccountNo("8888888")
	if err != nil {
		panic(err)
	} else {
		fmt.Println("账号设置成功！")
	}
	// 设置余额
	err = a.SetMoney(80)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("余额设置成功！")
	}
	// 设置密码
	err = a.SetPwd("123456")
	if err != nil {
		panic(err)
	} else {
		fmt.Println("密码设置成功！")
	}

	fmt.Println(&a)

}