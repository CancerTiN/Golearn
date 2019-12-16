package main

import "fmt"

func main() {
	var arr [4][6]int
	arr[1][2] = 1
	arr[2][1] = 2
	arr[2][3] = 3
	fmt.Println(arr)
	for _, row := range arr {
		for _, v := range row {
			fmt.Print(v, " ")
		}
		fmt.Println()
	}

	test()
}

func test()  {
	
}
