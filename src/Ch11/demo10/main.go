package main

import "fmt"

type Point struct {
	x int
	y int
}

func main() {
	var a interface{}
	var point = Point{1, 2}
	a = point
	fmt.Println(a)
	var b Point
	// cannot use a (type interface {}) as type Point in assignment: need type assertion
	// b = a // error
	b = a.(Point) // ok
	fmt.Println(b)

	var x interface{}
	var c float32 = 2.1
	x = c
	// interface conversion: interface {} is float32, not float64
	// y := x.(float64) // error
	// y := x.(float32)
	// y, ok := x.(float64)
	if y, ok := x.(float64); ok {
		fmt.Printf("The type of y is %T\n", y)
		fmt.Printf("The value of y is %v\n", y)
	} else {
		fmt.Println("Type conversion failed.")
	}
	fmt.Println("Continue ...")
}
