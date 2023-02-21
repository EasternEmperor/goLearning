package process

import (
	"fmt"
	"projects/mass_user_communication_system/common/message"
	"encoding/json"
)

// 打印群发消息
func printGroupMes(mes *message.Message) (err error) {
	// 反序列化mes.Data为message.SmsMes
	var smsMes message.SmsMes
	err = json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		err = fmt.Errorf("printGroupMes反序列化mes.Data失败！err = %v", err)
		return 
	}

	// 打印信息
	info := fmt.Sprintf("用户%d群发消息:\n%s", smsMes.UserId, smsMes.Content)
	fmt.Println(info, "\n")
	return 
}