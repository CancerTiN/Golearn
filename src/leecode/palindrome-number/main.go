package main

import "fmt"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	slice := make([]int, 0)
	for x > 0 {
		slice = append(slice, x%10)
		x /= 10
	}
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		if slice[i] == slice[j] {
			continue
		} else {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isPalindrome(12133121))
}
