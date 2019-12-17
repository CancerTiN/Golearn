package main

import "fmt"

type Person struct {
	Name   string
	Age    int
	Scores [5]float64
	ptr    *int
	slice  []int
	dict   map[string]string
}

type Monster struct {
	Name string
	Age  int
}

func main() {
	var p1 Person
	fmt.Println(p1)

	if p1.ptr == nil {
		fmt.Println("ok1")
	}
	if p1.slice == nil {
		fmt.Println("ok2")
	}
	if p1.dict == nil {
		fmt.Println("ok3")
	}

	// p1.slice[0] = 100 // error
	// Be sure to use the make function to allocate memory before using slicing.
	p1.slice = make([]int, 10)
	p1.slice[0] = 100

	// p1.dict["key1"] = "tom~" // error
	// Be sure to use the make function to allocate memory before using map.
	p1.dict = make(map[string]string)
	p1.dict["key1"] = "tom~"

	fmt.Println(p1)

	var m1 Monster
	m1.Name = "OxDevil"
	m1.Age = 500
	m2 := m1 // Structure is value type, and this operation is a value copy.
	m2.Name = "QingNiuJing"

	fmt.Println("monster1 =", m1) // {OxDevil 500}
	fmt.Println("monster2 =", m2) // {QingNiuJing 500}
}
