package main

import (
	"fmt"
	"math/rand"
	"time"
)

const n = 45533

func main() {
	rand.Seed(time.Now().UnixNano())
	var array [n]int
	for i := 0; i < n; i++ {
		array[i] = rand.Intn(n)
	}

	testSelect(array)
	testBubble(array)
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
