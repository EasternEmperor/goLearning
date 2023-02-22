package model

import (
	"fmt"
	"mysql_/utils"
)

// 创建User结构体
type User struct {
	Id			int
	Username	string
	Password	string
	Email		string
}

// 方法一：AddUser 添加User方法，使用Prepare
func (user *User) AddUser() (err error) {
	// 1. 写语句
	sqlStr := "insert into users(username, password, email) values(?,?,?)"
	// 2. 预编译
	inStmt, err := utils.Db.Prepare(sqlStr)
	if err != nil {
		fmt.Println("预编译异常！err =", err)
		return
	}

	// 3. 执行语句
	_, err = inStmt.Exec(user.Username, user.Password, user.Email)
	if err != nil {
		fmt.Println("执行出现异常！err =", err)
		return
	}
	return
}

// 方法二：直接使用DB的Exec
func (user *User) AddUser2() (err error) {
	// 1. 写语句
	sqlStr := "insert into users(username, password, email) values(?,?,?)"
	// 2. 执行语句
	_, err = utils.Db.Exec(sqlStr, user.Username, user.Password, user.Email)
	if err != nil {
		fmt.Println("执行出现异常！err =", err)
		return
	}
	return
}