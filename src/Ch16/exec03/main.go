package main

import (
	"fmt"
	"time"
)

func main() {
	var flag bool
	var primeSlice []int

	startTime := float64(time.Now().UnixNano()) / 1e9
	for num := 1; num <= 80000; num++ {
		flag = true
		for i := 2; i < num; i++ {
			if num%i == 0 {
				flag = false
				break
			}
		}
		if flag {
			primeSlice = append(primeSlice, num)
		}
	}
	endTime := float64(time.Now().UnixNano()) / 1e9
	fmt.Printf("It takes %v seconds to use the normal method.\n", endTime-startTime)
	fmt.Printf("A total of %v prime numbers were found.\n", len(primeSlice))
	fmt.Println("The main thread exits.")
}
