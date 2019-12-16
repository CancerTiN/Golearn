package main

import "fmt"

func main() {
	n := 8
	fmt.Printf("No.%d Fibonacci number is %d\n", n, fbn(n))
}

func fbn(n int) int {
	if n <= 2 {
		return 1
	} else {
		return fbn(n-1) + fbn(n-2)
	}
}
