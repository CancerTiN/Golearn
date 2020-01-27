package main

import "fmt"

func main() {
	switch 4 {
	case 1:
		test01()
	case 2:
		test02()
	case 3:
		test03()
	case 4:
		test04()
	}
}

func test01() {
	if true {
		defer fmt.Print("1")
	} else {
		defer fmt.Print("2")
	}
	fmt.Print("3")
}

func test02() {
	defer fmt.Print(1)
	defer func() { recover() }()
	defer fmt.Print(2)
	panic("abc")
	// 21
}

func test03() {
	defer fmt.Print(1)
	defer recover() // recover() is invalid
	defer fmt.Print(2)
	panic("abc")
	// 21panic: abc
}

func test04() {
	defer fmt.Print(1)
	defer fmt.Print(2)
	panic("abc")
	// 21panic: abc
}
