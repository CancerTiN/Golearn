package main

import "fmt"

func main() {
	n := 1
	fmt.Printf("%d peaches on day %d", day(n), n)
}

func day(n int) int {
	if n == 10 {
		return 1
	} else {
		// day(9) / 2 - 1 == day(10)
		return 2 * (day(n+1) + 1)
	}
}
