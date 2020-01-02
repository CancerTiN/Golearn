package main

import (
	"fmt"
	"os"
	"strings"
)

var userId int
var userPwd string

func main() {
	var key int
	var loop bool = true
	for loop {
		fmt.Println(strings.Repeat("-", 12), "Welcome to multi-player chat system", strings.Repeat("-", 12))
		fmt.Println(strings.Repeat("\t", 3), "1 Login to chat room")
		fmt.Println(strings.Repeat("\t", 3), "2 User registration")
		fmt.Println(strings.Repeat("\t", 3), "3 Exit system")
		fmt.Println("please choose (1-3):")

		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			fmt.Println("Login to chat room")
			loop = false
		case 2:
			fmt.Println("User registration")
			loop = false
		case 3:
			fmt.Println("Exit system")
			os.Exit(0)
			loop = false
		default:
			fmt.Println("Your input is wrong, please try again")
		}
	}
	if key == 1 {
		fmt.Println("Please enter user ID:")
		// fmt.Scanln(&userId)
		fmt.Scanf("%d\n", &userId)
		fmt.Println("Please enter user password:")
		// fmt.Scanln(&userPwd)
		fmt.Scanf("%s\n", &userPwd)
		err := login(userId, userPwd)
		if err != nil {
			fmt.Println("Login failed")
		}
	} else if key == 2 {
		fmt.Println("Logic for user registration ...")
	}

}
