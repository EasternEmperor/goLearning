package process

import (
	"fmt"
	"projects/mass_user_communication_system/common/message"
	"projects/mass_user_communication_system/client/utils"
	"encoding/json"
)

type SmsProcess struct {
	// 暂时无字段
}

func (this *SmsProcess) SendGroupMes(content string) (err error) {

	// 1. 创建一个消息mes
	var mes message.Message
	mes.Type = message.SmsMesType

	// 2. 创建一个SmsMes实例
	var smsMes message.SmsMes
	smsMes.Content = content
	smsMes.UserId = CurUser.UserId
	smsMes.UserStatus = CurUser.UserStatus

	// 3. 序列化smsMes
	data, err := json.Marshal(smsMes)
	if err != nil {
		err = fmt.Errorf("SendGroupMes序列化smsMes出错！err = %v", err)
		return
	}
	
	// 4.组装mes并序列化
	mes.Data = string(data)
	data, err = json.Marshal(mes)
	if err != nil {
		err = fmt.Errorf("SendGroupMes序列化mes出错！err = %v", err)
		return
	}

	// 5. 发送mes序列化后的data
	tf := &utils.Transfer{
		Conn : CurUser.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		err = fmt.Errorf("SendGroupMes发送mes出错！err = %v", err)
		return
	}
	return

}