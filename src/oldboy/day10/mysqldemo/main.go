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

func queryRow(n int) (u user, err error) {
	query := "SELECT `id`, `name`, `age` FROM `users` WHERE `id`=?"
	row := DB.QueryRow(query, n)
	err = row.Scan(&u.id, &u.name, &u.age)
	return
}

func queryRows(closedMin int) (users []user, err error) {
	query := "SELECT `id`, `name`, `age` FROM `users` WHERE `id`>?"
	rows, err := DB.Query(query, closedMin)
	if err != nil {
		return
	}
	defer func() { _ = rows.Close() }()
	for rows.Next() {
		var u user
		err = rows.Scan(&u.id, &u.name, &u.age)
		if err != nil {
			return
		}
		users = append(users, u)
	}
	return
}

func insertRow(name string, age int) (res sql.Result, err error) {
	query := "INSERT INTO `users` (`name`, `age`) VALUES (?, ?)"
	res, err = DB.Exec(query, name, age)
	if err != nil {
		return
	}
	return
}

func updateRow(id int, age int) (res sql.Result, err error) {
	query := "UPDATE `users` SET `age`=? WHERE `id`=?"
	res, err = DB.Exec(query, age, id)
	return
}

func deleteRow(id int) (res sql.Result, err error) {
	query := "DELETE FROM `users` WHERE `id`=?"
	res, err = DB.Exec(query, id)
	return
}

func prepareInsertRows(m map[string]int) (resMap map[string]sql.Result, err error) {
	query := "INSERT INTO `users` (`name`, `age`) VALUES (?, ?)"
	stmt, err := DB.Prepare(query)
	if err != nil {
		return
	}
	defer func() { _ = stmt.Close() }()
	resMap = make(map[string]sql.Result, len(m))
	for name, age := range m {
		res, err := stmt.Exec(name, age)
		if err != nil {
			break
		}
		resMap[name] = res
	}
	return
}

func main() {
	defer func() {
		_ = DB.Close()
	}()

	for i := 1; i < 5; i++ {
		u, err := queryRow(i)
		fmt.Printf("queryRow(%#v) return: (%#v, %#v).\n", i, u, err)
	}

	closedMin := 0
	users, err := queryRows(closedMin)
	fmt.Printf("queryRows(%#v) return: (%#v, %#v).\n", closedMin, users, err)

	var (
		found bool
		id    int
		name  = "图朝阳"
		age   int
	)
	for _, u := range users {
		if u.name == name {
			found = true
		}
	}
	if !found {
		res, err := insertRow(name, age)
		fmt.Printf("insertRow(%#v, %#v) return: (%#v, %#v).\n", name, age, res, err)
	}

	id = 5
	age = 9000
	res, err := updateRow(id, age)
	fmt.Printf("updateRow(%#v, %#v) return: (%#v, %#v).\n", id, age, res, err)
	fmt.Println(res.LastInsertId())
	fmt.Println(res.RowsAffected())

	id = 2
	res, err = deleteRow(id)
	fmt.Printf("deleteRow(%#v) return: (%#v, %#v).\n", id, res, err)
	fmt.Println(res.LastInsertId())
	fmt.Println(res.RowsAffected())

	m := map[string]int{
		"六七强": 30,
		"王相机": 32,
		"天说":  72,
		"白慧姐": 40,
	}
	resMap, err := prepareInsertRows(m)
	fmt.Printf("prepareInsertRows(%#v) return: (%#v, %#v).\n", m, resMap, err)
}
