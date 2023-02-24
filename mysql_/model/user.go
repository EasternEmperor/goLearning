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
func (this *User) String() string {
	str := fmt.Sprintf("Id: %v\tusername: %v\tpassword: %v\temail: %v", this.Id, 
						this.Username, this.Password, this.Email)
	return str
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

// 查询用户：通过Id
func (this *User) GetUserById() (err error) {
	// 1. sql语句
	sqlStr := "select * from users where id = ?"
	// 2. 执行语句
	row := utils.Db.QueryRow(sqlStr, this.Id)

	// 3. 声明要接收的用户
	err = row.Scan(&this.Id, &this.Username, &this.Password, &this.Email)
	if err != nil {
		fmt.Println("GetUserById查询用户出错！err =", err)
		return
	}
	return
}

// 查询所有用户
func (this *User) GetUsers() (users []*User, err error) {
	// 1. sql语句
	sqlStr := "select * from users"
	// 2. 执行语句
	rows, err := utils.Db.Query(sqlStr)
	if err != nil {
		fmt.Println("GetUsers查询用户出错！err =", err)
		return
	}

	// 3. 声明User切片接收查询到的用户
	// 注意在用Scan()获取值之前要使用Next()
	for rows.Next() {
		var user = &User{}
		err = rows.Scan(&user.Id, &user.Username, &user.Password, &user.Email)
		if err != nil {
			fmt.Println("从rows读取用户失败！err =", err)
			return
		}
		users = append(users, user)
	}
	return
}