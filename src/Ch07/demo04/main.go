package main

import (
	"fmt"
	"math"
)

func main() {
	vf := 32.1
	vi := int(vf)
	fmt.Printf("type of vf is %T\n", vf)
	fmt.Printf("type of vi is %T\n", vi)
	fmt.Printf("type of vi is %T\n", math.Round(1.1))
	fmt.Printf("type of vi is %v\n", math.Round(1.1))

	var arr01 [3]int
	arr01[0] = 1
	arr01[1] = 30
	arr01[2] = 11
	// arr01[2] = int(float64(1.1)) // error
	arr01[2] = int(math.Round(1.1))
	fmt.Println(arr01)

	var arr02 [3]float32
	var arr03 [3]string
	var arr04 [3]bool
	fmt.Println(arr02)
	fmt.Println(arr03)
	fmt.Println(arr04)

	var arr05 [3]string
	var index int = 2
	arr05[index] = "tom"

	arr06 := [3]int{11, 22, 33}
	fmt.Printf("address of main array06 is %p\n", &arr06)
	test01(arr06)
	fmt.Printf("main array06 is %v\n", arr06)
	test02(&arr06)
	fmt.Printf("main array06 is %v\n", arr06)
}

// The length is part of the array type,
// and you need to consider the length of the array when passing function parameters.
func test01(arr [3]int) {
	arr[0] = 88
	fmt.Printf("test01 array is %v\n", arr)
}

func test02(arr *[3]int) {
	fmt.Printf("address of test02 array is %p\n", arr)
	fmt.Printf("address of pointer of test02 array is %p\n", &arr)
	(*arr)[0] = 88
	fmt.Printf("test02 array is %v\n", *arr)
}
