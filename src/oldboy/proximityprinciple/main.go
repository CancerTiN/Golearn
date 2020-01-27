package main

import "fmt"

func main() {
	test1098()
	fmt.Println()
	test10910()
}

func test1098() {
	fmt.Println("test1098")
	var a int = 10
	fmt.Print(a)
	{
		a = 9
		fmt.Print(a)
		a = 8
	}
	fmt.Print(a)
}

func test10910() {
	fmt.Println("test10910")
	var a int = 10
	fmt.Print(a)
	{
		a := 9
		fmt.Print(a)
		a = 8
	}
	fmt.Print(a)
}
