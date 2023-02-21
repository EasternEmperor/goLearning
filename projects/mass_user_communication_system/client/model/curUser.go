package model

import (
	"net"
	"projects/mass_user_communication_system/common/message"
)

// 表示当前客户端的用户
// 要作为全局变量，可以让客户端都能访问
type CurUser struct {
	Conn net.Conn
	message.User
}