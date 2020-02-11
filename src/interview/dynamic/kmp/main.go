package main

import "fmt"

func violentMatch(str1, str2 string) (m int) {
	i, j := 0, 0
	runeSlice1 := []rune(str1)
	runeSlice2 := []rune(str2)
	for i < len(runeSlice1) && j < len(runeSlice2) {
		if runeSlice1[i] == runeSlice2[j] {
			i++
			j++
		} else {
			i -= j - 1
			j = 0
		}
	}
	if j == len(runeSlice2) {
		m = i - j
	} else {
		m = -1
	}
	return
}

func main() {
	testViolentMatch()
}

func testViolentMatch() {
	str1 := "中中人 国中人你国中 国中人你国中人你国中你好"
	str2 := "国中人你国中你"
	index := violentMatch(str1, str2)
	fmt.Println(index)
}
