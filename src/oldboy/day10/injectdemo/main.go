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

func sqlInjectDemo() {
	var name string
	name = "xxx' or 1=1 #"
	name = "xxx' union select * from users #"
	users := make([]user, 0)
	query := fmt.Sprintf("SELECT `id`, `name`, `age` FROM `users` WHERE `name`='%s'", name)
	fmt.Printf("query = %#v.\n", query)
	err := DB.Select(&users, query)
	fmt.Printf("DB.Select(%#v, %#v) return: (%#v).\n", users, query, err)
}

func main() {
	sqlInjectDemo()
}
