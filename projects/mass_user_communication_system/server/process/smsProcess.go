package process

import (
	"fmt"
	"projects/mass_user_communication_system/common/message"
	"projects/mass_user_communication_system/server/utils"
	"encoding/json"
)

type SmsProcess struct {
	// 暂时不需要字段
}

// 服务端群发消息
func (this *SmsProcess) SendGroupMes(mes *message.Message) (err error) {
	// 取出mes的SmsMes实例获取要发送消息的用户的id
	var smsMes *message.SmsMes
	err = json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		err = fmt.Errorf("SendGroupMes获取smsMes失败！err = %v", err)
		return
	}

	// 遍历服务器UserMgr给所有在线用户（除了发消息的用户）发送
	for id, up := range userMgr.onlineUsers {
		// 过滤掉自己
		if id == smsMes.UserId {
			continue
		}
		err = this.SendMesToOneUser(mes, up)
	}
	return
}

// 给某个用户发消息
func (this *SmsProcess) SendMesToOneUser(mes *message.Message, up *UserProcess) (err error) {
	
	// 序列化mes，准备将其转发给各在线用户
	data, err := json.Marshal(mes)
	if err != nil {
		err = fmt.Errorf("SendGroupMes序列化mes失败！err = %v", err)
		return
	}

	// 创建Transfer实例传输mes
	tf := &utils.Transfer{
		Conn : up.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		err = fmt.Errorf("SendMesToOneUser发送mes失败！err = %v", err)
		return
	}
	return
}