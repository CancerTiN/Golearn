package main

import "fmt"

type AInterface interface {
	Test01()
	Test02()
}

type BInterface interface {
	Test01()
	Test03()
}

type Student struct{}

func (s Student) Test01() {}
func (s Student) Test02() {}
func (s Student) Test03() {}

func main() {
	// The interface is based on the implementation of the method.
	s := Student{}
	var a AInterface = s
	var b BInterface = s
	fmt.Println("a =", a, "b =", b, "ok~")

}
