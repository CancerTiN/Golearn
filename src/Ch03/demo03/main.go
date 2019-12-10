package main

import "fmt"

var globalN1 = 100
var globalN2 = 200
var globalName1 = "jack"

var (
	globalN3    = 300
	globalN4    = 400
	globalName2 = "peter"
)

func main() {
	// Declare multiple variables at once
	// type 1
	var n1, n2, n3 int
	fmt.Println("n1 =", n1, "n2 =", n2, "n3 =", n3)

	// type 2
	var n4, name1, n5 = 100, "tom!", 888
	fmt.Println("n4 =", n4, "name1 =", name1, "n5 =", n5)

	// type 3 Type inference
	n6, name2, n7 := 100, "tom~", 888
	fmt.Println("n6 =", n6, "name2 =", name2, "n7 =", n7)

	// type 4 Output global variables type
	fmt.Println("globalN1 =", globalN1, "globalN2 =", globalN2, "globalName1 =", globalName1)

	// type 5 Output global variables
	fmt.Println("globalN3 =", globalN3, "globalN4 =", globalN4, "globalName2 =", globalName2)
}
