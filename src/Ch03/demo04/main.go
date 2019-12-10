package main

import "fmt"

// Notes on the use of variables
func main() {
	// The data value of this area can change continuously within the same type range
	var i int
	i = 30
	fmt.Println("i =", i)
	i = 50
	fmt.Println("i =", i)

	// error 1 Cannot change data type
	// i = 10.2

	// error 2 Variables cannot have the same name in the same scope
	// var i int = 59
	// i := 99

	// If the variable is not assigned an initial value, the compiler will use the default value
}
