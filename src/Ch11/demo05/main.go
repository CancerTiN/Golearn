package main

import "fmt"

type AInterface interface {
	Say()
}

type Student struct {
	Name string
}

func (s Student) Say() {
	fmt.Println("Students said ...")
}

type integer int

func (i integer) Say() {
	fmt.Println("integer say i is", i)
}

type Binterface interface {
	Hello()
}

type Monster struct {
}

func (m Monster) Hello() {
	fmt.Println("Hello, monster!")
}

func (m Monster) Say() {
	fmt.Println("Speaking, monster!")
}

func main() {
	var a AInterface
	fmt.Println("a =", a)
	// a.Say() // panic: runtime error: invalid memory address or nil pointer dereference
	var student Student
	a = student
	fmt.Println("a =", a)
	a.Say()

	var i integer = 10
	var b AInterface = i
	b.Say()

	var monster Monster
	var am AInterface = monster
	var bm Binterface = monster
	am.Say()
	bm.Hello()
}
