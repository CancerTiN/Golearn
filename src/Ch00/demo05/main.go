package main

import (
	"fmt"
	"strconv"
)

func main() {
	str := "ABCDEFG"
	//      0123456
	fmt.Println(str[4 : len(str)-1])
	fmt.Println(str[4:len(str)])

	var keepStr string

	index := 3
	keepStr = string(str[index])
	fmt.Println(keepStr)

	for i := index + 1; i < len(str); i++ {
		keepStr += string(str[i])
		fmt.Println(keepStr)
	}

	s := "0400"
	base := 10
	bitSize := 0
	val64, err := strconv.ParseInt(s, base, bitSize)
	if err != nil {
		fmt.Printf("strconv.ParseInt(%s, %d, %d) error: %v", s, base, bitSize, err)
	} else {
		fmt.Println("val64 =", val64)
	}
}
