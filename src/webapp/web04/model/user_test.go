package model

import (
	"fmt"
	"testing"
)

const flag = false

func TestMain(m *testing.M) {
	fmt.Println("测试开始！")
	m.Run()
}

func TestUser(t *testing.T) {
	fmt.Println("开始测试User中的相关方法！")
	if flag {
		t.Run("测试testAddUser", testAddUser)
		t.Run("测试testGetUserById", testGetUserByID)
	}
	t.Run("测试TestGetUsers", testGetUsers)
}

func testAddUser(t *testing.T) {
	if flag {
		fmt.Println("测试添加用户！")
		user := &User{}
		_ = user.AddUser()
		_ = user.AddUserAlpha()
	} else {
		fmt.Println("执行子测试函数！")
	}
}

func testGetUserByID(t *testing.T) {
	fmt.Println("测试查询一条记录！")
	u := User{ID: 1}
	user, err := u.GetUserByID()
	if err != nil {
		fmt.Println("u.GetUserById error:", err)
	} else {
		fmt.Println("得到user的信息是：", user)
	}
}

func testGetUsers(t *testing.T) {
	fmt.Println("测试查询所有记录！")
	u := &User{}
	users, err := u.GetUsers()
	if err != nil {
		fmt.Println("u.GetUsers error:", err)
	} else {
		fmt.Println("得到users的信息是：", users)
	}
	for i, user := range users {
		fmt.Printf("第%d个用户是%v\n", i+1, user)
	}
}
