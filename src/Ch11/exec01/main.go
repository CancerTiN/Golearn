package main

import "fmt"

type MethodUtils struct{}

func (mu MethodUtils) Print1() {
	for i := 1; i <= 10; i++ {
		for j := 1; j <= 8; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func (mu MethodUtils) Print2(m int, n int) {
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func (mu MethodUtils) area(length float64, width float64) float64 {
	return length * width
}

func (mu *MethodUtils) JudgeNum(num int) {
	if num%2 == 0 {
		fmt.Println("even")
	} else {
		fmt.Println("odd")
	}
}

func (mu MethodUtils) Print3(m int, n int, key string) {
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			fmt.Print(key)
		}
		fmt.Println()
	}
}

func main() {
	var mu MethodUtils
	mu.Print1()
	mu.Print2(5, 20)
	area := mu.area(2.5, 8.7)
	fmt.Println("area =", area)
	mu.JudgeNum(11)
	mu.Print3(7, 4, "@")
}
