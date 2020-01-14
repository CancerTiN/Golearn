package main

import (
	"fmt"
)

func main() {
	slice := []int{-9, 78, 0, 23, -567, 70}
	//slice = []int{-1, -1, -1, -1}
	//slice = []int{3, 3, 3}
	//slice = []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	fmt.Println(slice)
	myQuickSort(slice, 0, len(slice)-1)
	fmt.Println(slice)

	slice = []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	fmt.Println(slice)
	quickSort(slice, 0, len(slice)-1)
	fmt.Println(slice)
}

func quickSort(slice []int, left int, right int) {
	var l = left
	var r = right
	var pivot = slice[(left+right)/2]
	var n int
	for n = 0; l < r; n++ {
		fmt.Printf("A {n: %d, l: %d, r: %d}\n", n, l, r)
		for slice[l] < pivot {
			l++
		}
		for pivot < slice[r] {
			r--
		}
		fmt.Printf("B {n: %d, l: %d, r: %d}\n", n, l, r)
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
		fmt.Printf("C {n: %d, l: %d, r: %d}\n", n, l, r)
		if n > 10 {
			break
		}
	}
	if l == r {
		l++
		r--
	}
	fmt.Printf("D {n: %d, l: %d, r: %d}\n", n, l, r)
	if left < r {
		quickSort(slice, left, r)
	}
	if l < right {
		quickSort(slice, l, right)
	}
}

func myQuickSort(slice []int, left int, right int) {
	pivot := slice[(left+right)/2]
	l, r := left, right
	for l < r {
		// You cannot use <if> here, because the data transformation of <l> and <r>
		// should be isolated from the outer <for> loop.
		for slice[l] < pivot {
			l++
		}
		for pivot < slice[r] {
			r--
		}
		if slice[l] >= pivot && pivot >= slice[r] {
			slice[l], slice[r] = slice[r], slice[l]
		}
		if l >= r {
			break
		}
		if slice[l] == pivot {
			l++
		}
		if pivot == slice[r] {
			r--
		}
		fmt.Printf("slice[%d]=%d; pivot=%d; slice[%d]=%d\n", l, slice[l], pivot, r, slice[r])
	}
	if l == r {
		l--
		r++
	}
	if left < l {
		myQuickSort(slice, left, l)
	}
	if r < right {
		myQuickSort(slice, r, right)
	}
}
