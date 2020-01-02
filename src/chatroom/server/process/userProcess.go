package process

import (
	"chatroom/common/message"
	"chatroom/server/utils"
	"encoding/json"
	"fmt"
	"net"
)

type UserProcess struct {
	Conn net.Conn
}

func (u *UserProcess) ServerProcessLogin(mes *message.Message) (err error) {
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("json.Unmarshal error:", err)
		return
	}

	var loginResMes message.LoginResMes
	if loginMes.UserId == 100 && loginMes.UserPwd == "123456" {
		loginResMes.Code = 200
	} else {
		loginResMes.Code = 500
		loginResMes.Error = "This user does not exist, please register before using ..."
	}

	loginResMesData, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("json.Marshal error:", err)
		return
	}

	var resMes message.Message
	resMes.Type = message.LoginResMesType
	resMes.Data = string(loginResMesData)

	resMesData, err := json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal error:", err)
		return
	}

	tf := &utils.Transfer{
		Conn: u.Conn,
	}

	err = tf.WritePkg(resMesData)
	return
}
