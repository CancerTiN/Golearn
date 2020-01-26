package dao

import (
	"fmt"
	"testing"
)

func TestUser(t *testing.T) {
	fmt.Println("测试userdao中的函数！")
	t.Run("验证用户名或密码", testLogin)
	t.Run("验证用户名", testRegist)
	t.Run("保存用户", testSave)
}

func testLogin(t *testing.T) {
	user, err := CheckUsernameAndPassword("admin", "123456")
	fmt.Printf("CheckUsernameAndPassword return: (%v, %v).\n", user, err)
}

func testRegist(t *testing.T) {
	user, err := CheckUsername("admin")
	fmt.Printf("CheckUsername return: (%v, %v).\n", user, err)
}

func testSave(t *testing.T) {
	err := SaveUser("admin_alpha", "123456", "admin@aliyun.com")
	fmt.Printf("testSave return: (%v).\n", err)
}
