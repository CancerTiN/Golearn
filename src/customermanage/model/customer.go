package model

import "fmt"

type Customer struct {
	Id     int
	Name   string
	Gender string
	Age    int
	Phone  string
	Email  string
}

func InitCustomer(id int, name string, gender string, age int, phone string, email string) Customer {
	return Customer{id, name, gender, age, phone, email}
}

func NewCustomer(name string, gender string, age int, phone string, email string) Customer {
	return Customer{Name: name, Gender: gender, Age: age, Phone: phone, Email: email}
}

func (c Customer) GetInfo() string {
	return fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v", c.Id, c.Name, c.Gender, c.Age, c.Phone, c.Email)
}
