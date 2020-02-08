package main

import (
	"fmt"
)

func Bubble(slice []int) {
	var exchange bool
	for i := len(slice) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if slice[j] > slice[j+1] {
				slice[j], slice[j+1] = slice[j+1], slice[j]
				exchange = true
			}
		}
		if !exchange {
			break
		}
	}
}

func Select(slice []int) {
	for i := 0; i < len(slice)-1; i++ {
		for j := i + 1; j < len(slice); j++ {
			if slice[i] > slice[j] {
				slice[i], slice[j] = slice[j], slice[i]
			}
		}
	}
}

func Insert(slice []int) {
	for i := 1; i < len(slice); i++ {
		insertIndex := i
		insertValue := slice[insertIndex]
		for j := i - 1; j >= 0; j-- {
			if slice[j] >= insertValue {
				slice[j+1] = slice[j]
				insertIndex = j
			}
		}
		if insertIndex != i {
			slice[insertIndex] = insertValue
		}
	}
}

func Quick(slice []int, left, right int) {
	l, r := left, right
	pivot := slice[(l+r)/2]
	for l < r {
		for slice[l] < pivot {
			l++
		}
		for pivot < slice[r] {
			r--
		}
		slice[l], slice[r] = slice[r], slice[l]
		if slice[l] == pivot {
			l++
		}
		if pivot == slice[r] {
			r--
		}
	}
	if l == r {
		l++
		r--
	}
	if left < r {
		Quick(slice, left, r)
	}
	if l < right {
		Quick(slice, l, right)
	}
}

func main() {
	slice := []int{3, 44, 38, 5, 47, 15, 36, 26, 27, 2, 46, 4, 19, 50, 48}
	fmt.Println(slice)
	switch "quick" {
	case "bubble":
		Bubble(slice)
	case "select":
		Select(slice)
	case "insert":
		Insert(slice)
	case "quick":
		Quick(slice, 0, len(slice)-1)
	}
	fmt.Println(slice)
}
