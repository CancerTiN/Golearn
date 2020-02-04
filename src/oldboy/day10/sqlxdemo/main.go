package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

type user struct {
	Id   int
	Name string
	Age  int
}

func init() {
	var (
		diverName      = "mysql"
		dataSourceName = "root:root@tcp(127.0.0.1:3306)/sql_test"
		err            error
	)
	DB, err = sqlx.Connect(diverName, dataSourceName)
	if err != nil {
		panic(err)
	}
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
}

func getDemo() {
	var u user
	query := "SELECT `id`, `name`, `age` FROM `users` WHERE `id`=1"
	err := DB.Get(&u, query)
	fmt.Printf("DB.Get(%#v, %#v) return: (%#v).\n", &u, query, err)
}

func selectDemo() {
	users := make([]user, 0)
	query := "SELECT `id`, `name`, `age` FROM `users`"
	err := DB.Select(&users, query)
	fmt.Printf("DB.Select(%#v, %#v) return: (%#v).\n", users, query, err)
}

func main() {
	switch "select" {
	case "get":
		getDemo()
	case "select":
		selectDemo()
	}
}
