package model

import (
	"fmt"
	"webapp/web04/utils"
)

type User struct {
	ID       int
	Username string
	Password string
	Email    string
}

func (u *User) AddUser() (err error) {
	sqlStr := "INSERT into users (username, password, email) VALUES (?, ?, ?)"
	inStmt, err := utils.Db.Prepare(sqlStr)
	if err != nil {
		fmt.Println("预编译出现异常：", err)
	}
	_, err = inStmt.Exec("admin", "123456", "admin@gmail.com")
	if err != nil {
		fmt.Println("执行出现异常：", err)
	}
	return
}

func (u *User) AddUserAlpha() (err error) {
	sqlStr := "INSERT into users (username, password, email) VALUES (?, ?, ?)"
	_, err = utils.Db.Exec(sqlStr, "admin_alpha", "666666", "admin_alpha@sina.com")
	if err != nil {
		fmt.Println("执行出现异常：", err)
	}
	return
}

func (u *User) GetUserByID() (user *User, err error) {
	var (
		id       int
		username string
		password string
		email    string
	)
	sqlStr := "SELECT id, username, password, email FROM users WHERE id = ?"
	row := utils.Db.QueryRow(sqlStr, u.ID)
	err = row.Scan(&id, &username, &password, &email)
	if err != nil {
		return
	}
	user = &User{ID: id, Username: username, Password: password, Email: email}
	return
}

func (u *User) GetUsers() (users []*User, err error) {
	var (
		id       int
		username string
		password string
		email    string
	)
	sqlStr := "SELECT id, username, password, email FROM users"
	rows, err := utils.Db.Query(sqlStr)
	if err != nil {
		return
	}
	for rows.Next() {
		err = rows.Scan(&id, &username, &password, &email)
		if err != nil {
			return
		} else {
			user := &User{ID: id, Username: username, Password: password, Email: email}
			users = append(users, user)
		}
	}
	return
}
