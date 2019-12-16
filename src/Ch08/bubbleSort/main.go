package main

import (
	"fmt"
)

func main() {
	slice := []int{24, 69, 80, 57, 13}
	fmt.Println("origin slice:", slice)
	fmt.Println("sorted slice:", bubbleSortSlice(slice))

	array := [...]int{24, 69, 80, 57, 13}
	fmt.Println("origin array:", array)
	fmt.Println("sorted array:", bubbleSortArray(&array))
}

func bubbleSortArray(array *[5]int) [5]int {
	for n := len(array) - 1; n > 0; n-- {
		for j := 0; j < n; j++ {
			if (*array)[j] < (*array)[j+1] {
				(*array)[j], (*array)[j+1] = (*array)[j+1], (*array)[j]
			}
		}
		fmt.Println(*array)
	}
	return *array
}

func bubbleSortSlice(slice []int) []int {
	n := len(slice) - 1
	var eFlag bool
	for {
		eFlag = false
		for i := 0; i < n; i++ {
			if slice[i] > slice[i+1] {
				slice[i], slice[i+1] = slice[i+1], slice[i]
				eFlag = true
			}
		}
		if !eFlag {
			return slice
		}
		n--
		if n == 0 {
			break
		}
		fmt.Println(slice)
	}
	return slice
}
