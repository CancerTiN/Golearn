package main

import (
	"chatroom/client/process"
	"fmt"
	"os"
	"strings"
)

var userId int
var userPwd string
var userName string

func main() {
	var key int
	var loop bool = true
	for loop {
		fmt.Println(strings.Repeat("-", 8), "Welcome to multi-player chat system", strings.Repeat("-", 8))
		fmt.Println("\t1 Login to chat room")
		fmt.Println("\t2 User registration")
		fmt.Println("\t3 Exit system")
		fmt.Println("please choose (1-3):")

		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			fmt.Println("登录聊天室")
			fmt.Print("请输入用户编号：")
			fmt.Scanf("%d\n", &userId)
			fmt.Print("请输入用户密码：")
			fmt.Scanf("%s\n", &userPwd)
			up := process.UserProcess{}
			_ = up.Login(userId, userPwd)
		case 2:
			fmt.Println("注册用户")
			fmt.Print("请输入用户编号：")
			fmt.Scanf("%d\n", &userId)
			fmt.Print("请输入用户密码：")
			fmt.Scanf("%s\n", &userPwd)
			fmt.Print("请输入用户名字：")
			fmt.Scanf("%s\n", &userName)
			up := process.UserProcess{}
			_ = up.Register(userId, userPwd, userName)
		case 3:
			fmt.Println("Exit system")
			os.Exit(0)
			loop = false
		default:
			fmt.Println("Your input is wrong, please try again")
		}
	}
	//if key == 1 {
	//	fmt.Println("Please enter user ID:")
	//	// fmt.Scanln(&userId)
	//	fmt.Scanf("%d\n", &userId)
	//	fmt.Println("Please enter user password:")
	//	// fmt.Scanln(&userPwd)
	//	fmt.Scanf("%s\n", &userPwd)
	//	err := login(userId, userPwd)
	//	if err != nil {
	//		fmt.Println("Login failed")
	//	}
	//} else if key == 2 {
	//	fmt.Println("Logic for user registration ...")
	//}

}
