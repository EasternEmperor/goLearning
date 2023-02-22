package model

import (
	"fmt"
	"testing"
)

func TestAddUser(t *testing.T) {
	fmt.Println("测试添加用户...")
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