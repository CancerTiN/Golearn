package main

import "fmt"

func main() {
	test01()
}

func test01() {
	res := (1+6)/2*4 ^ 2 + 10%3<<3
	fmt.Println(res)
}
