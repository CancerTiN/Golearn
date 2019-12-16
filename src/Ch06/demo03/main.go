package main

import "fmt"

func main() {
	n := 2
	fmt.Printf("f(%d) return %d\n", n, f(n))
}

func f(n int) int {
	if n == 1 {
		return 3
	} else {
		return 2*f(n-1) + 1
	}
}
