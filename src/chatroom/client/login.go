package main

import (
	"chatroom/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
)

func login(userId int, userPwd string) (err error) {
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
	fmt.Printf("The length of the message sent by the client is %d bytes\n", pkgLen)
	fmt.Printf("The content of the message sent by the client is '%v'\n", dataMes)

	fmt.Printf("userId = %d\n", userId)
	fmt.Printf("userId = %s\n", userPwd)

	return nil
}
