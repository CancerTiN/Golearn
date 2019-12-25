package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type CharCount struct {
	EnglishCount int
	NumberCount  int
	SpaceCount   int
	OtherCount   int
}

func main() {
	name := `D:\Workspace\Golearn\src\Ch14\demo00\python.txt`
	file, err := os.Open(name)
	if err != nil {
		fmt.Println("Open error:")
		fmt.Println(err)
		return
	}
	defer file.Close()

	//var count CharCount

	reader := bufio.NewReader(file)
	var i int
	var charCount CharCount
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Read string error:")
			fmt.Println(err)
		}
		for j, v := range str {
			fmt.Printf("line: %v; index: %v; byte: %v; char: '%c'.\n", i, j, v, v)
			switch {
			case v >= 'a' && v <= 'z':
				fallthrough
			case v >= 'A' && v <= 'z':
				charCount.EnglishCount++
			case v == ' ' || v == '\t':
				charCount.SpaceCount++
			case v >= '0' && v <= '9':
				charCount.NumberCount++
			default:
				charCount.OtherCount++
			}
		}
	}
	fmt.Println("The statistical results are as follows:")
	fmt.Printf("English count: %v.\n", charCount.EnglishCount)
	fmt.Printf("Number count: %v.\n", charCount.NumberCount)
	fmt.Printf("Space count: %v.\n", charCount.SpaceCount)
	fmt.Printf("Other count: %v.\n", charCount.OtherCount)
}
