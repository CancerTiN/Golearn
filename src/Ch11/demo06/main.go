package main

import "fmt"

type BInterface interface {
	test01()
}

type CInterface interface {
	test02()
}

type AInterface interface {
	BInterface
	CInterface
	test03()
	// cannot use student (type Student) as type AInterface in assignment:
	// Student does not implement AInterface (missing test04 method)
	// test04()
}

type Student struct{}

func (s Student) test01() {}

func (s Student) test02() {}

func (s Student) test03() {}

type T interface{}

func main() {
	var student Student
	var a AInterface = student
	a.test01()

	var t T = student
	fmt.Println("t =", t)

	// There is no method for the empty interface.
	// All types implement the empty interface.
	// We can assign any variable to the empty interface!
	var ti interface{} = student
	fmt.Println("ti =", ti)
	var num float64 = 8.8
	ti = num
	fmt.Println("ti =", ti)
	ti = "anything"
	fmt.Println("ti =", ti)
}
