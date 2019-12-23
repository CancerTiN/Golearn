package main

import "fmt"

type Usb interface {
	Say()
}

type Student struct{}

func (this *Student) Say() {
	fmt.Println("Student is specking.")
}

func main() {
	var s Student = Student{}
	// Student does not implement Usb (Say method has pointer receiver)
	// var u Usb = s // error
	var u Usb = &s
	u.Say()
	fmt.Println("here", u)
}
