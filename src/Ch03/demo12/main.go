package main

import "fmt"

func main() {
	var a int
	var b float32
	var c float64
	var d bool
	var e string
	fmt.Printf("a=%d,b=%f,c=%f,d=%v,e=%v\n", a, b, c, d, e)
	fmt.Printf("a=%d,b=%v,c=%f,d=%v,e=%v\n", a, b, c, d, e)
}
