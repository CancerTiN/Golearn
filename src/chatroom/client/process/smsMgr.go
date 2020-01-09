package process

import (
	"chatroom/common/message"
	"encoding/json"
	"fmt"
)

func outputGroupMes(mes *message.Message) {
	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("json.Unmarshal([]byte(mes.Data), &smsMes)错误：", err)
		return
	}
	fmt.Printf("用户（编号%d）对大家说：%v\n", smsMes.UserId, smsMes.Content)
}
