package main

import "fmt"

func main() {
	fmt.Println("Calling function (test1)")
	test1(4)
	fmt.Println("Calling function (test2)")
	test2(4)
}

func test1(n int) {
	if n > 2 {
		n--
		test1(n)
	}
	fmt.Println("n =", n)
}

func test2(n int) {
	if n > 2 {
		n--
		test2(n)
	} else {
		fmt.Println("n =", n)
	}
}
