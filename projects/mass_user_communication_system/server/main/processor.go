package main

import (
	"fmt"
	"net"
	"projects/mass_user_communication_system/common/message"
	"projects/mass_user_communication_system/server/process"
	"projects/mass_user_communication_system/server/utils"
	"errors"
)

// 定义一个Processor结构体用于调用这个包的方法
type Processor struct {
	// 连接
	Conn net.Conn
}

// 根据客户端发送消息类型不同选择不同处理方法
func (this *Processor) serverProcessMes(mes *message.Message) (err error) {
	switch mes.Type {

	case message.LoginMesType:
		// 处理登录
		// 创建UserProcess实例处理登录请求
		var up = process.UserProcess{
			Conn : this.Conn,
		}
		err = up.ServerProcessLogin(mes)

	case message.RegisterMesType:
		// 处理注册
		var up = process.UserProcess{
			Conn : this.Conn,
		}
		err = up.ServerProcessRegister(mes)

	case message.SmsMesType:
		// 处理群发消息请求
		// 创建SmsProcess实例处理
		var smsProcess = &process.SmsProcess{}
		// 调用方法群发mes
		err = smsProcess.SendGroupMes(mes)

	case message.LogoutMesType:
		// 处理离线请求
		var up = process.UserProcess{
			Conn : this.Conn,
		}
		err = up.ServerProcessLogout(mes)

	default:
		err = errors.New("消息类型不存在，无法处理！")
	}
	return 
}

// 读取客户端发送消息
func (this *Processor) serverProcessRead() (err error) {
	// 循环读取
	for {
		// 读取传输的包
		// 创建Tranfer实例收取消息
		var tf = &utils.Transfer{
			Conn : this.Conn,
		}
		var mes message.Message
		mes, err = tf.ReadPkg()
		if err != nil {
			str := fmt.Sprintln("读取失败！err =", err)
			err = errors.New(str)
			return
		}
		// 处理输入
		err = this.serverProcessMes(&mes)
		if err != nil {
			return
		}
	}

	return 
}