package main

import (
	"fmt"
	"time"
)

func main() {
	slice := []int{10, 34, 19, 100, 80}
	fmt.Println(slice)
	startTime := float64(time.Now().UnixNano()) / 1e9
	SelectSort(slice)
	endTime := float64(time.Now().UnixNano()) / 1e9
	fmt.Println(slice)
	fmt.Printf("SelectSort takes %f seconds.\n", endTime-startTime)
}

func SelectSort(slice []int) {
	for i := 0; i < len(slice); i++ {
		maxIndex := i
		maxValue := slice[maxIndex]

		for j := maxIndex + 1; j < len(slice); j++ {
			if slice[j] > maxValue {
				maxIndex = j
				maxValue = slice[maxIndex]
			}
		}

		if maxIndex != i {
			slice[i], slice[maxIndex] = slice[maxIndex], slice[i]
		}
	}
}
