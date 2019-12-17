package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	// type 1
	var p1 Person
	fmt.Println(p1)

	// type 2
	p2 := Person{}
	// p2 := Person{"tom", 18}
	p2.Name = "tom"
	p2.Age = 18
	fmt.Println(p2)

	// type 3
	// var p3 *Person = new(Person)
	p3 := new(Person)
	// p3.Name = "smith" // right
	// `p3.Name = "smith"` is equivalent to `(*p3).Name = "smith"`, which is actually handled by Golang.
	// The way the program is called at the bottom is `(*p3).Name = "smith"`.
	(*p3).Name = "smith"
	(*p3).Age = 30
	fmt.Println(*p3)
	p3.Name = "john"
	p3.Age = 100
	fmt.Println(*p3)
	fmt.Println(p3) // pointer

	// type 4
	var person *Person = &Person{"mary", 60}
	// var person *Person = &Person{}
	fmt.Println(person) // pointer
	(*person).Name = "scott"
	(*person).Age = 88
	fmt.Println(*person)
}
