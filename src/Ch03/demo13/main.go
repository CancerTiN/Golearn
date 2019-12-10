package main

import "fmt"

func main() {
	var i int = 100
	// var n1 float32 = i // error
	var n1 float32 = float32(i)
	var n2 int8 = int8(i)

	fmt.Printf("i = %v\n", i)
	fmt.Printf("n1 = %v\n", n1)
	fmt.Printf("n2 = %v\n", n2)
}
