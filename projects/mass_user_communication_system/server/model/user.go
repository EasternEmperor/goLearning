package model

// 定义用户的结构体
type User struct {
	// 为了能够序列化和反序列化成功，必须保证User字段对应的tag名
	// 和序列化的mes中的用户信息json串的名字一致！
	UserId int		`json:"userId"`
	UserPwd string	`json:"userPwd"`
	UserName string	`json:"userName"`
}