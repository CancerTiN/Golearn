package main

import (
	"fmt"
)

const n = 3
const m = 4

func display(matrix [n + 1][m + 1]int) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Print(matrix[i][j], " ")
		}
		fmt.Println()
	}
}

func main() {
	var weights = []int{1, 4, 3}
	var values = []int{1500, 3000, 2000}
	var vMat [n + 1][m + 1]int
	var path [n + 1][m + 1]int

	for i := 0; i < n+1; i++ {
		vMat[i][0] = 0
	}
	for j := 0; j < m+1; j++ {
		vMat[0][j] = 0
	}

	display(vMat)

	for i := 1; i < len(vMat); i++ {
		for j := 1; j < len(vMat[i]); j++ {
			if weights[i-1] > j {
				vMat[i][j] = vMat[i-1][j]
			} else {
				oldValue := vMat[i-1][j]
				newValue := values[i-1] + vMat[i-1][j-weights[i-1]]
				if oldValue < newValue {
					vMat[i][j] = newValue
					path[i][j] = 1
				} else {
					vMat[i][j] = oldValue
				}
			}
		}
	}

	fmt.Println()
	display(vMat)
	fmt.Println()
	display(path)

	r := len(path) - 1
	c := len(path[r]) - 1
	for r > 0 && c > 0 {
		if path[r][c] == 1 {
			fmt.Printf("Put product %#v in the backpack at path[%#v][%#v].\n", r, r, c)
			c -= weights[r-1]
		}
		r--
	}
}
