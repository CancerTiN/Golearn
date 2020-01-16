package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 45533
// 30969
const n = 30969

func try() {
	slice := []int{10, 34, 19, 100, 80}
	fmt.Println(slice)
	Insert(slice)
	fmt.Println(slice)
}

func main() {
	var flag bool = false
	if flag {
		try()
	} else {
		rand.Seed(time.Now().UnixNano())
		var array [n]int
		for i := 0; i < n; i++ {
			array[i] = rand.Intn(n)
		}
		testBubble(array)
		testSelect(array)
		testInsert(array)
	}
}

func testBubble(array [n]int) {
	slice := array[:]
	fmt.Println(slice)
	startTime := float64(time.Now().UnixNano()) / 1e9
	Bubble(slice)
	endTime := float64(time.Now().UnixNano()) / 1e9
	fmt.Println(slice)
	fmt.Printf("BubbleSort takes %f seconds.\n", endTime-startTime)
}

func testSelect(array [n]int) {
	slice := array[:]
	fmt.Println(slice)
	startTime := float64(time.Now().UnixNano()) / 1e9
	Select(slice)
	endTime := float64(time.Now().UnixNano()) / 1e9
	fmt.Println(slice)
	fmt.Printf("SelectSort takes %f seconds.\n", endTime-startTime)
}

func testInsert(array [n]int) {
	slice := array[:]
	fmt.Println(slice)
	startTime := float64(time.Now().UnixNano()) / 1e9
	Insert(slice)
	endTime := float64(time.Now().UnixNano()) / 1e9
	fmt.Println(slice)
	fmt.Printf("InsertSort takes %f seconds.\n", endTime-startTime)
}
