package process

import (
	"chatroom/common/message"
	"chatroom/server/model"
	"chatroom/server/utils"
	"encoding/json"
	"fmt"
	"net"
)

type UserProcess struct {
	Conn net.Conn
	// onlineUsers map[int]*UserProcess
	UserId int
}

func (up *UserProcess) NotifyOthersOnlineUser(userId int) {
	for id, oup := range userMgr.onlineUsers {
		if id == userId {
			continue
		}
		oup.NotifyMeOnline(userId)
	}
}

func (up *UserProcess) NotifyMeOnline(userId int) {
	var notifyUserStatusMes message.NotifyUserStatusMes
	notifyUserStatusMes.UserId = userId
	notifyUserStatusMes.Status = message.UserOnline

	notifyUserStatusMesData, err := json.Marshal(notifyUserStatusMes)
	if err != nil {
		fmt.Println("json.Marshal(notifyUserStatusMes)错误：", err)
		return
	}

	var mes message.Message
	mes.Type = message.NotifyUserStatusMesType
	mes.Data = string(notifyUserStatusMesData)

	mesData, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal(mesData)错误：", err)
		return
	}

	tf := &utils.Transfer{Conn: up.Conn}
	err = tf.WritePkg(mesData)
	if err != nil {
		fmt.Println("tf.WritePkg(mesData)错误：", err)
		return
	}

}

func (up *UserProcess) ServerProcessRegister(mes *message.Message) (err error) {
	fmt.Println("到了ServerProcessRegister这里，mes为：", mes)
	var registerMes message.RegisterMes
	err = json.Unmarshal([]byte(mes.Data), &registerMes)
	if err != nil {
		fmt.Println("反序列化注册信息错误：", err)
		return
	}

	var registerResMes message.RegisterResMes

	err = model.MyUserDao.Register(&registerMes.User)
	if err != nil {
		if err == model.ERROR_USER_EXISTS {
			registerResMes.Code = 505
			registerResMes.Error = err.Error()
		} else {
			registerResMes.Code = 506
			registerResMes.Error = "注册发生未知错误"
		}
	} else {
		registerResMes.Code = 200
	}

	registerResMesData, err := json.Marshal(registerResMes)
	if err != nil {
		fmt.Println("序列化注册返回信息数据错误：", err)
	}

	var resMes message.Message
	resMes.Type = message.RegisterResMesType
	resMes.Data = string(registerResMesData)

	resMesData, err := json.Marshal(resMes)
	if err != nil {
		fmt.Println("序列化注册返回信息错误：", err)
		return
	}

	tf := &utils.Transfer{
		Conn: up.Conn,
	}

	err = tf.WritePkg(resMesData)

	return
}

func (up *UserProcess) ServerProcessLogin(mes *message.Message) (err error) {
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("json.Unmarshal error:", err)
		return
	}

	fmt.Println("到了userProcess这里，登录信息为：", loginMes)
	var loginResMes message.LoginResMes

	user, err := model.MyUserDao.Login(loginMes.UserId, loginMes.UserPwd)

	if err != nil {
		if err == model.ERROR_USER_NOTEXISTS {
			loginResMes.Code = 500
			loginResMes.Error = err.Error()
		} else if err == model.ERROR_USER_PWD {
			loginResMes.Code = 300
			loginResMes.Error = err.Error()
		} else {
			loginResMes.Code = 505
			loginResMes.Error = "服务器内部错误..."
		}
	} else {
		loginResMes.Code = 200

		up.UserId = loginMes.UserId
		userMgr.AddOnlineUser(up)

		up.NotifyOthersOnlineUser(loginMes.UserId)

		for id, _ := range userMgr.onlineUsers {
			loginResMes.UserIds = append(loginResMes.UserIds, id)
		}

		fmt.Println("登陆成功：", user)
	}

	//if loginMes.UserId == 100 && loginMes.UserPwd == "123456" {
	//	loginResMes.Code = 200
	//} else {
	//	loginResMes.Code = 500
	//	loginResMes.Error = "This user does not exist, please register before using ..."
	//}

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
		Conn: up.Conn,
	}

	err = tf.WritePkg(resMesData)
	return
}
