package main

import "fmt"

func main() {
	var arr = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(arr)

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			fmt.Printf("%v\t", arr[i][j])
		}
		fmt.Println()
	}

	for i, v := range arr {
		for j, x := range v {
			fmt.Printf("arr[%v][%v] = %v\t", i, j, x)
		}
		fmt.Println()
	}
}
