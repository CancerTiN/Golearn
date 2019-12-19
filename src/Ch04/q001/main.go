package main

import "fmt"

func main() {
	deferCall()
}

func deferCall() {
	defer func() { fmt.Println("Before printing") }()
	defer func() { fmt.Println("Printing") }()
	defer func() { fmt.Println("After printing") }()

	panic("Trigger an exception")
}
