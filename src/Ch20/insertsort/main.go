package main

import "fmt"

func main() {
	slice := []int{23, 0, 12, 56, 34, -1, 55}
	fmt.Println(slice)
	InsertSort(slice)
	fmt.Println(slice)
}

func InsertSort(slice []int) {
	for i := 1; i < len(slice); i++ {
		insertIndex := i - 1
		insertValue := slice[i]
		for insertIndex >= 0 && slice[insertIndex] < insertValue {
			slice[insertIndex+1] = slice[insertIndex]
			insertIndex--
		}
		if insertIndex+1 != i {
			slice[insertIndex+1] = insertValue
		}
	}
}
