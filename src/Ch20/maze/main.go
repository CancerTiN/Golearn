package main

import (
	"fmt"
	"time"
)

const (
	nRow = 8
	nCol = 7
)

func SetMap(mapMat *[nRow][nCol]int, i, j int) bool {
	// display(mapMat)
	time.Sleep(time.Millisecond * 100)
	if mapMat[6][5] == 2 {
		return true
	} else {
		if mapMat[i][j] == 0 {
			mapMat[i][j] = 2
			if SetMap(mapMat, i+1, j) {
				return true
			} else if SetMap(mapMat, i, j+1) {
				return true
			} else if SetMap(mapMat, i-1, j) {
				return true
			} else if SetMap(mapMat, i, j-1) {
				return true
			} else {
				mapMat[i][j] = 3
				return false
			}
		} else {
			return false
		}
	}
}

func main() {
	var mapMat [nRow][nCol]int

	for i := 0; i < nRow; i++ {
		if i == 0 || i == nRow-1 {
			for j := 0; j < nCol; j++ {
				mapMat[i][j] = 1
			}
		} else {
			mapMat[i][0] = 1
			mapMat[i][nCol-1] = 1
		}
	}
	mapMat[3][1] = 1
	mapMat[3][2] = 1
	mapMat[5][2] = 1
	mapMat[5][3] = 1
	mapMat[5][4] = 1
	mapMat[5][5] = 1
	mapMat[2][3] = 1

	display(&mapMat)
	SetMap(&mapMat, 1, 1)
	fmt.Println("After SetMap:")
	display(&mapMat)
}

func display(mapMat *[nRow][nCol]int) {
	for i := 0; i < nRow; i++ {
		for j := 0; j < nCol; j++ {
			fmt.Print(mapMat[i][j], " ")
		}
		fmt.Println()
	}
}
