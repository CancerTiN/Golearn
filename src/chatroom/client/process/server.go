package process

import (
	"chatroom/client/utils"
	"chatroom/common/message"
	"encoding/json"
	"fmt"
	"net"
	"os"
)

func ShowMenu(userId int) {
	fmt.Printf("--------恭喜%v登录成功--------\n", userId)
	fmt.Println("\t1. 在线用户")
	fmt.Println("\t2. 发送信息")
	fmt.Println("\t3. 信息列表")
	fmt.Println("\t4. 退出系统")
	fmt.Print("请选择（1-4）：")
	var key int
	fmt.Scanf("%d\n", &key)

	switch key {
	case 1:
		outputOnlineUser()
	case 2:
		fmt.Println("发送信息")
	case 3:
		fmt.Println("信息列表")
	case 4:
		fmt.Println("你选择了退出系统...")
		os.Exit(0)
	default:
		fmt.Print("你输入的选项不正确，请重新输入（1-4）：")
	}
}

func serverProcessMes(conn net.Conn) {
	tf := &utils.Transfer{Conn: conn}
	for {
		fmt.Println("客户端正在等待读取服务器发送的信息")
		mes, err := tf.ReadPkg()
		if err != nil {
			fmt.Println("客户端读取服务器发送的信息错误：", err)
			return
		}
		fmt.Printf("客户端读取到服务器发送的信息为：%v\n", mes)
		switch mes.Type {
		case message.NotifyUserStatusMesType:
			var notifyUserStatusMes message.NotifyUserStatusMes
			json.Unmarshal([]byte(mes.Data), &notifyUserStatusMes)
			updateUserStatus(notifyUserStatusMes)
		default:
			fmt.Println("服务器返回了未知的消息类型")
		}
	}
}
