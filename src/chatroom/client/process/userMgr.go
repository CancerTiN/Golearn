package process

import (
	"chatroom/common/message"
	"fmt"
)

var onlineUserMap map[int]*message.User = make(map[int]*message.User, 10)

func outputOnlineUser() {
	fmt.Println("当前在线用户列表：")
	for id, _ := range onlineUserMap {
		fmt.Println("用户编号：\t", id)
	}
}

func updateUserStatus(notifyUserStatusMes message.NotifyUserStatusMes) {
	user, ok := onlineUserMap[notifyUserStatusMes.UserId]
	if !ok {
		user = &message.User{UserId: notifyUserStatusMes.UserId}
	}
	user.UserStatus = notifyUserStatusMes.Status
	onlineUserMap[notifyUserStatusMes.UserId] = user

	outputOnlineUser()
}
