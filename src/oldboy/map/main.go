package main

import "fmt"

func test01() {
	var m1 map[string]int
	// m1["one"] = 1
	var m2 map[string]int
	m2 = make(map[string]int)
	m2["one"] = 2
	fmt.Println(m1, m2)
}

func main() {
	test01()
}
