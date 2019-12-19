package main

import (
	"Ch11/demo02/model"
	"fmt"
)

func main() {
	p := model.NewPerson("smith")
	p.SetAge(18)
	p.SetSalary(5000)
	fmt.Println("p ->", p)
	fmt.Printf("name: %v; age: %v; salary: %v\n", p.Name, p.GetAge(), p.GetSalary())
}
