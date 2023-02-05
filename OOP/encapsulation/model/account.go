package model

import (
	"fmt"
	"errors"
)

type Account struct {
	AccountNo string
	money int
	pwd string
}
// Set
// 赋值账号，不符合条件则返回error
func (acnt *Account) SetAccountNo(no string) error {
	if len(no) < 6 || len(no) > 10 {
		return errors.New("账号错误！")
	}
	acnt.AccountNo = no
	return nil
}
// 赋值余额
func (acnt *Account) SetMoney(m int) error {
	if m < 20 {
		return errors.New("余额不符合规定！")
	}
	acnt.money = m
	return nil
}
// 赋值密码
func (acnt *Account) SetPwd(password string) error {
	if len(password) != 6 {
		return errors.New("密码错误！")
	}
	acnt.pwd = password
	return nil
}
// Get
// 获取账号
func (acnt *Account) GetAccountNo() string {
	return acnt.AccountNo
}
// 获取余额
func (acnt *Account) GetMoney() int {
	return acnt.money
}
// 获取密码
func (acnt *Account) GetPwd() string {
	return acnt.pwd
}
// string方法，用于打印
func (acnt *Account) String() string {
	str := fmt.Sprintf("账号：%s\n余额：%d\n密码：%s", acnt.AccountNo, acnt.money, acnt.pwd)
	return str
}