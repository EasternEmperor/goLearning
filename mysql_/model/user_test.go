package model

import (
	"fmt"
	"testing"
)

// TestMain()可以在测试函数执行前做一些其他操作
func TestMain(m *testing.M) {
	// 这里可以做前置操作
	fmt.Println("测试开始：")

	// 调用测试函数
	m.Run()
}

// 测试函数：测试User结构体的方法
func TestUser(t *testing.T) {
	fmt.Println("开始测试User中的方法：")

	// 通过Run()执行测试子函数
	// t.Run("测试添加用户：", testAddUser)
	t.Run("测试按Id查询用户：", testGetUserById)
	t.Run("测试查询所有用户：", testGetUsers)
}

// 测试子函数：测试添加用户
func testAddUser(t *testing.T) {
	fmt.Println("AddUser()子测试执行...")
	user1 := &User{1, "ming", "123456", "ming@qq.com"}
	user2 := &User{2, "ning", "654321", "ning@qq.com"}
	// 调用添加用户的方法
	err := user1.AddUser()
	if err != nil {
		t.Fatalf("AddUser()方法错误！err = %v", err)
	}
	user2.AddUser2()
	if err != nil {
		t.Fatalf("AddUser2()方法错误！err = %v", err)
	}
	t.Logf("测试成功！")
}

// 测试子函数：测试查询用户
func testGetUserById(t *testing.T) {
	fmt.Println("GetUserById()子测试执行...")
	user := &User{
		Id : 1,
	}
	// 调用GetUserById()
	err := user.GetUserById()
	if err != nil {
		t.Fatalf("未能成功获取用户！err = %v", err)
	}
	fmt.Println("获取的User: ", user)
}

// 测试子函数：测试查询所有用户
func testGetUsers(t *testing.T) {
	fmt.Println("GetUsers()自测试执行...")
	user := &User{}
	// 调用获取所有错误的方法
	us, err := user.GetUsers()
	if err != nil {
		t.Fatalf("未能成功获取用户！err = %v", err)
	}

	fmt.Println("获取到的所有用户：")
	for idx, u := range us {
		fmt.Println("第", idx, "个用户是：")
		fmt.Println(u)
	}
}