package process

import (
	"fmt"
	"projects/mass_user_communication_system/common/message"
	"projects/mass_user_communication_system/client/model"
)

// 客户端要维护的map
var onlineUsers map[int]*message.User = make(map[int]*message.User, 10)
// 当前用户全局
var CurUser model.CurUser

// 根据收到的NotifyUserStatusMes更新onlineUsers
func updateUserStatus(notifyUserStatusMes *message.NotifyUserStatusMes) {
	// 对维护的全局用户状态map变量进行更新
	user, ok := onlineUsers[notifyUserStatusMes.UserId]
	// 还没有该用户，则新建一个
	if !ok {
		user = &message.User{
			UserId : notifyUserStatusMes.UserId,
		}
	}
	user.UserStatus = notifyUserStatusMes.Status
	onlineUsers[notifyUserStatusMes.UserId] = user

	printOnlineUser()
}

// 打印所有在线用户
func printOnlineUser() {
	// 遍历onlineUsers
	fmt.Println("当前在线用户列表：")
	for id, user := range onlineUsers {
		if user.UserStatus == message.UserOnline {
			fmt.Println("用户id:\t", id)
		}
	}
}