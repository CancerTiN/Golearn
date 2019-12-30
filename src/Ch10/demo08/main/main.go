package main

import (
	"Ch10/demo08/model"
	"fmt"
)

func main() {
	p := model.NewPerson()
	fmt.Println(p)
	p.SetSalary(10000)
	fmt.Println(p)
}
