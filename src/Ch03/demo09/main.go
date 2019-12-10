package main

import "fmt"

func main() {
	var c1 byte = 'a'
	var c2 byte = '0'

	// When we directly output the byte value, we output the ASCII code corresponding to the variable
	fmt.Println("c1 =", c1)
	fmt.Println("c2 =", c2)

	// If we want to output the corresponding characters, we need to use formatted output
	fmt.Printf("c1 = %c\n", c1)
	fmt.Printf("c2 = %c\n", c2)

	// var c3 byte = '北' // overflow
	var c3 int = '北'
	fmt.Printf("c3 = %c\n", c3)
	fmt.Printf("ASCII code of c3 = %d\n", c3)

	// var c4 int = 22269 // '国'
	var c4 int = 120
	fmt.Printf("c4 = %c\n", c4)

	// The character type can be operated, which is equivalent to an integer,
	// the operation is performed according to the code value
	var n1 = 100 + 'a'
	fmt.Println("n1 =", n1)
}
