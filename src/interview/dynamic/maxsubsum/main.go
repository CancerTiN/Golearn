package main

import "fmt"

func maxSubItemSumOf(slice []int) (start, end, sum int) {
	maxHere := slice[start]
	for i := 1; i < len(slice); i++ {
		if maxHere <= 0 {
			start = i
			maxHere = slice[start]
		} else {
			maxHere += slice[i]
		}
		if maxHere > sum {
			sum = maxHere
			end = i + 1
		}
	}
	return
}

func main() {
	test()
}

func test() {
	slice := []int{1, -5, 8, 3, -4, 15, -8}
	fmt.Println(maxSubItemSumOf(slice))
}
