package process

import (
	"fmt"
	"net"
	"projects/mass_user_communication_system/client/utils"
	"projects/mass_user_communication_system/common/message"
	"encoding/json"
)

// 显示登录成功后的界面
func ShowLoginMenu() bool {
	fmt.Println("----------用户登录成功------------")
	fmt.Println("----------1. 显示在线用户列表-----")
	fmt.Println("----------2. 发送消息------------")
	fmt.Println("----------3. 消息列表------------")
	fmt.Println("----------4. 退出系统------------")
	fmt.Print("请选择(1 - 4):")

	var key int
	// 要发送的短消息
	var content string
	// 用于发送消息的SmsProcess实例
	var smsProcess = &SmsProcess{}
	fmt.Scanln(&key)
	switch key {
	case 1:
		//fmt.Println("显示在线用户列表")
		printOnlineUser()
		return true
	case 2:
		// fmt.Println("发送消息")
		fmt.Println("输入你要群发的消息：")
		fmt.Scanln(&content)
		smsProcess.SendGroupMes(content)
		return true
	case 3:
		fmt.Println("消息列表")
		return true
	case 4:
		fmt.Println("退出系统")
		// 处理离线请求
		var up = &UserProcess{}
		err := up.Logout(CurUser.UserId)
		if err != nil {
			fmt.Println(err.Error())
		}
		return false
	default:
		fmt.Println("输入不符合规范，请重新输入！")
		return true
	}
	
}

// 和服务器保持连接，读取服务端的消息
func processMes(conn net.Conn) {
	// 创建Transfer实例
	tf := &utils.Transfer{
		Conn : conn,
	}
	// 循环读取
	for {
		fmt.Println("\n客户端正在读取服务器发送的消息...")
		mes, err := tf.ReadPkg()
		if err != nil {
			fmt.Println("ProcessMes tf ReadPkg() err =", err)
			return
		}
		// 读取到消息输出，判断是什么类型，采取不同操作
		switch mes.Type {

		case message.NotifyUserStatusMesType:
			// 改变传输的userId的状态
			// 1. 取出mes中的NotifyUserStatusMes
			fmt.Println("\n修改用户状态...")
			var notifyUserStatusMes message.NotifyUserStatusMes
			err := json.Unmarshal([]byte(mes.Data), &notifyUserStatusMes)
			if err != nil {
				fmt.Println("反序列化mes.notifyUserStatusMes失败！err =", err)
				break
			}
			// 2. 把这个用户的信息，状态在onlineUsers中改变
			updateUserStatus(&notifyUserStatusMes)

		case message.SmsMesType:
			// 收到群发消息
			printGroupMes(&mes)

		default:
			fmt.Println("processMes获取消息状态未定义，无法解析！")
		}
	}
}