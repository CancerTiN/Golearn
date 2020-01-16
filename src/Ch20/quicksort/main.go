package main

import "fmt"

func main() {
	slice := []int{-9, 78, 0, 23, -567, 70, 123, 90, -23}
	fmt.Println(slice)
	QuickSort(slice, 0, len(slice)-1)
	fmt.Println(slice)
}

func QuickSort(slice []int, left, right int) {
	if right < 1 {
		return
	}
	l := left
	r := right
	pivot := slice[(left+right)/2]
	for l < r {
		for slice[l] < pivot {
			l++
		}
		for pivot < slice[r] {
			r--
		}
		if l >= r {
			break
		}
		slice[l], slice[r] = slice[r], slice[l]
		if slice[l] == pivot {
			r--
		}
		if pivot == slice[r] {
			l++
		}
	}
	if l == r {
		l++
		r--
	}
	if left < r {
		QuickSort(slice, left, r)
	}
	if l < right {
		QuickSort(slice, l, right)
	}
}
