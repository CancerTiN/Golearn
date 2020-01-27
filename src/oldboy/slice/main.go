package main

import "fmt"

func test01() {
	// s1 := make([]int)
	s2 := make([]int, 0)
	s3 := make([]int, 5, 10)
	s4 := []int{1, 2, 3, 4, 5}
	fmt.Println(s2, s3, s4)
}

func test02() {
	s1 := []int{
		1, 2, 3,
		4, 5, 6,
	}
	//s2 := []int{
	//	1, 2, 3,
	//	4, 5, 6
	//}
	s3 := []int{
		1, 2, 3,
		4, 5, 6}
	s4 := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(s1, s3, s4)
}

func test03() {
	var s1 []int
	// s1[0] = 6
	var s2 []int
	s2 = make([]int, 0)
	s2 = append(s2, 1)
	fmt.Println(s1, s2)
}

func main() {
	switch 3 {
	case 1:
		test01()
	case 2:
		test02()
	case 3:
		test03()
	}
}
