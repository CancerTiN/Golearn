package main

import "fmt"

func main() {
	// type 1
	var i int
	fmt.Println("i =", i)

	// type 2
	// Type inference
	var num = 10.11
	fmt.Println("num =", num)

	// type 3
	// var name string; name = "tom"
	name := "tom"
	fmt.Println("name =", name)
}
