package dao

import (
	"bookstore/model"
	"bookstore/utils"
	"fmt"
)

func CheckUsernameAndPassword(username, password string) (user *model.User, err error) {
	user = &model.User{}
	sqlStr := "SELECT id, username, password, email FROM users WHERE username = ? AND password = ?"
	row := utils.Db.QueryRow(sqlStr, username, password)
	err = row.Scan(&user.ID, &user.Username, &user.Password, &user.Email)
	return
}

func CheckUsername(username string) (user *model.User, err error) {
	user = &model.User{}
	sqlStr := "SELECT id, username, password, email FROM users WHERE username = ?"
	row := utils.Db.QueryRow(sqlStr, username)
	err = row.Scan(&user.ID, &user.Username, &user.Password, &user.Email)
	return
}

func SaveUser(username, password, email string) (err error) {
	sqlStr := "INSERT INTO users (username, password, email) VALUES (?, ?, ?)"
	sqlRes, err := utils.Db.Exec(sqlStr, username, password, email)
	fmt.Printf("utils.Db.Exec return: (%v, %v).\n", sqlRes, err)
	return
}
