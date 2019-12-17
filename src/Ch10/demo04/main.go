package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	var p1 Person
	p1.Name = "ming"
	p1.Age = 10

	var p2 *Person = &p1

	fmt.Println((*p2).Age)
	fmt.Println(p2.Age)

	fmt.Printf("p2.Name=%v\tp1.Name=%v\n", (*p2).Name, p1.Name)

	p2.Name = "tom~" // value reference
	// *p2.Name = "tom~" // error
	// (*p2).Name = "tom~" // right

	fmt.Printf("p2.Name=%v\tp1.Name=%v\n", p2.Name, p1.Name)
	fmt.Printf("p2.Name=%v\tp1.Name=%v\n", (*p2).Name, p1.Name)

	fmt.Printf("The address of p1 is %p.\n", &p1)
	fmt.Printf("The content of p2 is %p.\n", p2)
	fmt.Printf("The address of p2 is %p.\n", &p2)

	p3 := p1 // value copy

	fmt.Printf("p3.Name=%v\tp1.Name=%v\n", p3.Name, p1.Name)

	p3.Name = "mary"

	fmt.Printf("p3.Name=%v\tp1.Name=%v\n", p3.Name, p1.Name)

	fmt.Printf("The address of p1 is %p.\n", &p1)
	fmt.Printf("The address of p3 is %p.\n", &p3)
}
