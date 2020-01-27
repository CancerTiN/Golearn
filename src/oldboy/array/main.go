package main

import "fmt"

func main() {
	test01()
}

func test01() {
	var a1 = [5]int{1, 2, 3, 4, 5}
	var a2 = [...]int{1, 2, 3, 4, 5}
	// var a3 [5]int = {1, 2, 3, 4, 5}
	var a4 = [5]int{2: 4, 4: 61}
	fmt.Println(a1, a2, a4)
}
