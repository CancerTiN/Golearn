package process

import (
	"chatroom/client/utils"
	"chatroom/common/message"
	"encoding/json"
	"fmt"
)

type SmsProcess struct {
}

func (sp *SmsProcess) SendGroupMes(content string) (err error) {
	var smsMes message.SmsMes
	smsMes.Content = content
	smsMes.UserId = CurUser.UserId
	smsMes.UserStatus = CurUser.UserStatus
	smsMesData, err := json.Marshal(smsMes)
	if err != nil {
		fmt.Println("json.Marshal(smsMes)错误：", err)
		return
	}

	var mes message.Message
	mes.Type = message.SmsMesType
	mes.Data = string(smsMesData)

	mesData, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal(mes)错误：", err)
		return
	}

	tf := &utils.Transfer{Conn: CurUser.Conn}
	err = tf.WritePkg(mesData)
	if err != nil {
		fmt.Println("tf.WritePkg(mesData)错误：", err)
		return
	}

	return
}
