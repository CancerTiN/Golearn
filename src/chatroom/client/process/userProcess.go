package process

import (
	"chatroom/client/utils"
	"chatroom/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
)

type UserProcess struct {
}

func (u *UserProcess) Register(userId int, userPwd string, userName string) (err error) {
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("net.Dial error:", err)
		return err
	}
	defer conn.Close()

	var registerMes message.RegisterMes
	registerMes.User.UserId = userId
	registerMes.User.UserPwd = userPwd
	registerMes.User.UserName = userName
	registerMesData, err := json.Marshal(registerMes)
	if err != nil {
		fmt.Println("json.Marshal(registerMes)错误：", err)
		return
	}

	var mes message.Message
	mes.Type = message.RegisterMesType
	mes.Data = string(registerMesData)

	dataMes, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal(mes)错误：", err)
		return err
	}

	tf := &utils.Transfer{Conn: conn}
	err = tf.WritePkg(dataMes)
	if err != nil {
		fmt.Println("发送注册信息错误：", err)
	}
	resMes, err := tf.ReadPkg()
	if err != nil {
		fmt.Println("收到注册信息错误：", err)
	}

	var registerResMes message.RegisterResMes
	err = json.Unmarshal([]byte(resMes.Data), &registerResMes)
	if registerResMes.Code == 200 {
		fmt.Println("注册成功，请重新登录...")
	} else {
		fmt.Println(registerResMes.Error)
	}
	return
}

func (u *UserProcess) Login(userId int, userPwd string) (err error) {
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("net.Dial error:", err)
		return err
	}
	defer conn.Close()

	var loginMes message.LoginMes
	loginMes.UserId = userId
	loginMes.UserPwd = userPwd
	loginMesData, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("json.Marshal error:", err)
		return err
	}

	var mes message.Message
	mes.Type = message.LoginMesType
	mes.Data = string(loginMesData)

	dataMes, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal error:", err)
		return err
	}

	var pkgLen = uint32(len(dataMes))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4], pkgLen)

	n, err := conn.Write(buf[:])
	if err != nil || n != 4 {
		fmt.Println("conn.Write error:", err)
		return err
	}

	_, err = conn.Write(dataMes)
	if err != nil {
		fmt.Println("conn.Write error:", err)
		return err
	}

	//fmt.Printf("The length of the message sent by the client is %d bytes\n", pkgLen)
	//fmt.Printf("The content of the message sent by the client is '%v'\n", dataMes)

	//fmt.Printf("userId = %d\n", userId)
	//fmt.Printf("userPwd = %s\n", userPwd)

	tf := &utils.Transfer{Conn: conn}
	resMes, err := tf.ReadPkg()
	if err != nil {
		fmt.Println("readPkg error:", err)
		return
	}
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(resMes.Data), &loginResMes)
	if loginResMes.Code == 200 {
		fmt.Println("login successful")
		go serverProcessMes(conn)
		for {
			ShowMenu()
		}
	} else {
		fmt.Println(loginResMes.Error)
	}

	return nil
}
