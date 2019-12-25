package main

import "fmt"

func main() {
	r := addUpper(11)
	if r != 55 {
		fmt.Printf("addUpper error; return: %v; expect: %v.\n", r, 55)
	} else {
		fmt.Printf("addUpper right; return: %v; expect: %v.\n", r, 55)
	}
}

func addUpper(n int) int {
	var r int
	for i := 0; i <= n; i++ {
		r += i
	}
	return r
}
