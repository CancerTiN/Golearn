package main

import "fmt"

func main() {
	var positiveCount int
	var negativeCount int
	var num int
	for {
		fmt.Print("Please enter an integer: ")
		fmt.Scanln(&num)
		if num == 0 {
			break
		}
		if num > 0 {
			positiveCount++
			continue
		}
		negativeCount++
	}
	fmt.Printf("Positive count: %d\n", positiveCount)
	fmt.Printf("Negative count: %d\n", negativeCount)
}
