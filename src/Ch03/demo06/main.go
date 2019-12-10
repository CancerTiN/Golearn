package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var i int = 1
	fmt.Println("i =", i)

	// -128 <= j <= 127
	var j int8 = 127
	fmt.Println("j =", j)

	// 0 <= k <= 255
	var k uint8 = 255
	fmt.Println("k =", k)

	var a int = 8900
	fmt.Println("a =", a)

	// b >= 0
	var b uint = 1
	fmt.Println("b =", b)

	// 0 <= c <= 255
	var c byte = 255
	fmt.Println("c =", c)

	// Integer usage details
	var n1 = 100
	// How to view the data type of a variable
	// fmt.Printf() can be used to do formatted output
	fmt.Printf("type of n1 is %T\n", n1)

	// How to view the occupied byte size and data type of a variable in a program
	var n2 int64 = 10
	fmt.Printf("type of n2 is %T, number of bytes occupied by n2 is %d\n", n2, unsafe.Sizeof(n2))

	// When using integer variables in Golang,
	// the principle of keeping small and not big must be observed,
	// under the premise of ensuring the correct operation of the program,
	// try to use data types that take up less space.
	var age byte = 150
	fmt.Println("age =", age)
}
