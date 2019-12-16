package main

import "fmt"

func main() {
	slice := []int{1, 8, 10, 89, 1000, 1234}
	i := binarySearch(89, slice, 0, len(slice)-1)
	fmt.Println(i)

	fmt.Println()

	array := [6]int{1, 8, 10, 89, 1000, 1234}
	index := binaryFind(&array, 0, len(array)-1, 1000)
	fmt.Println(index)
}

func binaryFind(array *[6]int, leftIndex int, rightIndex int, findValue int) int {
	fmt.Println(leftIndex, rightIndex)
	resultIndex := -1
	middleIndex := (leftIndex + rightIndex) / 2
	if leftIndex >= rightIndex {
		return resultIndex
	}
	if (*array)[middleIndex] == findValue {
		resultIndex = middleIndex
	} else if findValue < (*array)[middleIndex] {
		resultIndex = binaryFind(array, leftIndex, middleIndex, findValue)
	} else if (*array)[middleIndex] < findValue {
		resultIndex = binaryFind(array, middleIndex, rightIndex, findValue)
	}
	return resultIndex
}

func binarySearch(v int, s []int, l int, r int) int {
	fmt.Println(l, r)
	m := (l + r) / 2
	if l == r {
		return -1
	}
	if s[m] == v {
		return m
	} else if v < s[m] {
		return binarySearch(v, s, l, m)
	} else if s[m] < v {
		return binarySearch(v, s, m, r)
	}
	return -1
}
