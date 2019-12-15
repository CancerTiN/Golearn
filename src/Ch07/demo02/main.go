package main

import "fmt"

func main() {
	var intArr [3]int // int64 occupies 8 bytes.
	// After we define the array, each element of the array has a default value of 0.
	fmt.Println(intArr)
	fmt.Printf("The address of inArr is %p\n", &intArr)
	fmt.Printf("The address of inArr[0] is %p\n", &intArr[0])
	fmt.Printf("The address of inArr[1] is %p\n", &intArr[1])
	fmt.Printf("The address of inArr[2] is %p\n", &intArr[2])
	intArr[0] = 10
	intArr[1] = 20
	intArr[2] = 30
	fmt.Println(intArr)
	fmt.Printf("The address of inArr is %p\n", &intArr)
	fmt.Printf("The address of inArr[0] is %p\n", &intArr[0])
	fmt.Printf("The address of inArr[1] is %p\n", &intArr[1])
	fmt.Printf("The address of inArr[2] is %p\n", &intArr[2])

	if false {
		testScore()
	}

	var numArr01 [3]int = [3]int{1, 2, 3}
	fmt.Println("numArr01 =", numArr01)

	var numArr02 = [3]int{4, 5, 6}
	fmt.Println("numArr02 =", numArr02)

	var numArr03 = [...]int{7, 8, 9}
	fmt.Println("numArr03 =", numArr03)

	var numArr04 = [...]int{1: 800, 0: 900, 2: 999}
	fmt.Println("numArr04 =", numArr04)

	strArr05 := [...]string{1: "tom", 0: "jack", 2: "marry"}
	fmt.Println("strArr05 =", strArr05)

}

func testScore() {
	var score [5]float64

	for i := 0; i < len(score); i++ {
		fmt.Printf("Please enter the value of element [%d]: ", i)
		fmt.Scanln(&score[i])
	}
	for i := 0; i < len(score); i++ {
		fmt.Printf("score[%d] = %v\n", i, score[i])
	}
}
