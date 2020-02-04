package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

type user struct {
	id   int
	name string
	age  int
}

func init() {
	var (
		diverName      = "mysql"
		dataSourceName = "root:root@tcp(127.0.0.1:3306)/sql_test"
		err            error
	)
	// Does not verify that the username and password are correct.
	DB, err = sql.Open(diverName, dataSourceName)
	if err != nil {
		panic(err)
	}
	// Try to connect to the database.
	fmt.Printf("sql.Open return: (%v, %v).\n", DB, err)
	err = DB.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Printf("DB.Ping return: (%v).\n", err)
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
}

func transactionDoneDemo() {
	var (
		tx    *sql.Tx
		err   error
		query string
		res   sql.Result
	)
	tx, err = DB.Begin()
	if err != nil {
		fmt.Printf("DB.Begin() return: (%#v, %#v).\n", tx, err)
		return
	}
	query = "UPDATE `users` SET `age`=`age`-2 WHERE `id`=1"
	res, err = tx.Exec(query)
	if err != nil {
		fmt.Printf("tx.Exec(%#v) return: (%#v, %#v).\n", query, res, err)
		err = tx.Rollback()
		return
	}
	fmt.Println()
	query = "UPDATE `users` SET `age`=`age`+2 WHERE `id`=3"
	res, err = tx.Exec(query)
	if err != nil {
		fmt.Printf("tx.Exec(%#v) return: (%#v, %#v).\n", query, res, err)
		err = tx.Rollback()
		return
	}
	err = tx.Commit()
	if err != nil {
		fmt.Printf("tx.Commit() return: (%#v).\n", err)
		err = tx.Rollback()
		return
	}
	fmt.Println("The transaction executed successfully.")
}

func transactionFailDemo() {
	var (
		tx    *sql.Tx
		err   error
		query string
		res   sql.Result
	)
	tx, err = DB.Begin()
	if err != nil {
		fmt.Printf("DB.Begin() return: (%#v, %#v).\n", tx, err)
		return
	}
	query = "UPDATE `users` SET `age`=`age`-2 WHERE `id`=1"
	res, err = tx.Exec(query)
	if err != nil {
		fmt.Printf("tx.Exec(%#v) return: (%#v, %#v).\n", query, res, err)
		err = tx.Rollback()
		return
	}
	fmt.Println()
	// Transaction executed successfully with 0 row affected!
	query = "UPDATE `users` SET `age`=`age`+2 WHERE `id`=2"
	res, err = tx.Exec(query)
	if err != nil {
		fmt.Printf("tx.Exec(%#v) return: (%#v, %#v).\n", query, res, err)
		err = tx.Rollback()
		return
	}
	query = "UPDATE `xxxxx` SET `age`=`age`+2 WHERE `id`=2"
	res, err = tx.Exec(query)
	if err != nil {
		fmt.Printf("tx.Exec(%#v) return: (%#v, %#v).\n", query, res, err)
		err = tx.Rollback()
		return
	}
	err = tx.Commit()
	if err != nil {
		fmt.Printf("tx.Commit() return: (%#v).\n", err)
		err = tx.Rollback()
		return
	}
	fmt.Println("The transaction executed successfully.")
}

func main() {
	switch "fail" {
	case "done":
		transactionDoneDemo()
	case "fail":
		transactionFailDemo()
	}
}
