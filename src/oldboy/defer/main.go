package main

import (
	"fmt"
	"strings"
)

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

func test05() {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d", i)
	}
}

type Slice []int

func NewSlice() Slice {
	return make(Slice, 0)
}

func (s Slice) Add(elem int) *Slice {
	s = append(s, elem)
	fmt.Print(elem)
	return &s
}

func test06() {
	s := NewSlice()
	defer s.Add(1).Add(2)
	s.Add(3)
}

func hundredDividedBy(i int) {
	fmt.Println(100 / i)
}

func test07() {
	defer fmt.Println(strings.Repeat("a", 6)) // 3
	defer fmt.Println(strings.Repeat("b", 6)) // 2
	defer hundredDividedBy(0)                 // 4
	defer fmt.Println(strings.Repeat("c", 6)) // 1
}

func test08() {
	var i = 10
	defer fmt.Printf("defer i = %d\n", i)
	i = 1000
	fmt.Printf("i = %d\n", i)
}

func main() {
	switch 8 {
	case 1:
		test01()
	case 2:
		test02()
	case 3:
		test03()
	case 4:
		test04()
	case 5:
		test05()
	case 6:
		test06()
	case 7:
		test07()
	case 8:
		test08()
	default:
		document01()
	}
}

func document01() {
	// panic比defer执行的优先级高
}
