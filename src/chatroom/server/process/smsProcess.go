package process

import (
	"chatroom/common/message"
	"chatroom/server/utils"
	"encoding/json"
	"fmt"
	"net"
)

type SmsProcess struct {
}

func (sp *SmsProcess) SendGroupMes(mes *message.Message) (err error) {
	var smsMes message.SmsMes
	err = json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("json.Unmarshal([]byte(mes.Data), &smsMes)错误：", err)
		return
	}

	mesData, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal(mes)错误：", err)
		return
	}

	for id, up := range userMgr.onlineUsers {
		if id == smsMes.UserId {
			continue
		}
		sp.SendMesToEachOnlineUser(mesData, up.Conn)
	}

	return
}

func (sp *SmsProcess) SendMesToEachOnlineUser(data []byte, conn net.Conn) (err error) {
	tf := &utils.Transfer{Conn: conn}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("tf.WritePkg(data)错误：", err)
		return
	}

	return
}
