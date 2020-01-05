package model

import (
	"chatroom/common/message"
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"reflect"
)

var (
	MyUserDao *UserDao
)

type UserDao struct {
	pool *redis.Pool
}

func NewUserDao(pool *redis.Pool) (user *UserDao) {
	return &UserDao{pool: pool}
}

func (this *UserDao) getUserById(conn redis.Conn, id int) (user *User, err error) {
	res, err := redis.String(conn.Do("HGet", "users", id))
	fmt.Println("到了getUserById这里，error为：", err)
	if err != nil {
		fmt.Println("error不为nil，为：", err)
		if err == redis.ErrNil {
			fmt.Println("error具体为：", redis.ErrNil)
			err = ERROR_USER_NOTEXISTS
		} else {
			fmt.Println("未知error，实际为：", reflect.TypeOf(err))
		}
		return
	}
	user = &User{}
	err = json.Unmarshal([]byte(res), user)
	if err != nil {
		fmt.Println("反序列化错误：", err)
		return
	}
	return
}

func (this *UserDao) Login(userId int, userPwd string) (user *User, err error) {
	conn := this.pool.Get()
	defer conn.Close()
	fmt.Println("到了userDao这里，用户编号为：", userId)
	user, err = this.getUserById(conn, userId)
	fmt.Println("经过getUserById后，返回信息为：", user)
	fmt.Println("经过getUserById后，返回错误为：", err)
	if err != nil {
		fmt.Println("经过getUserId后，返回错误不为nil，马上返回")
		return
	}

	if user.UserPwd != userPwd {
		err = ERROR_USER_PWD
	}
	return
}

func (u *UserDao) Register(user *message.User) (err error) {
	fmt.Println("到了Register这里，user为：", user)
	conn := u.pool.Get()
	defer conn.Close()

	_, err = u.getUserById(conn, user.UserId)
	if err == nil {
		err = ERROR_USER_EXISTS
		return
	}

	data, err := json.Marshal(user)
	if err != nil {
		return
	}

	_, err = conn.Do("HSet", "users", user.UserId, string(data))
	if err != nil {
		fmt.Println("用户注册信息入库错误：", err)
		return
	}
	return
}
