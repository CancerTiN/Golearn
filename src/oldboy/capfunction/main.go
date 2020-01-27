package main

import "fmt"

func main() {
	switch 0 {
	default:
		document01()
	}
}

func document01() {
	fmt.Println(cap([2]int{}))
	fmt.Println(cap([]int{0, 1, 2}))
	fmt.Println(cap(make(chan int, 4)))
	//关于cap函数的适用类型包括：数组、切片和管道
}
