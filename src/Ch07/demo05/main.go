package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	var myChars [26]byte
	var myStrings [26]string
	var char byte

	for i := 0; i < len(myChars); i++ {
		char = 'A' + byte(i)
		myChars[i] = char
		myStrings[i] = fmt.Sprintf("%c", char)
		fmt.Printf("%c ", char)
	}
	fmt.Println()

	fmt.Println(myChars)
	for _, v := range myChars {
		fmt.Printf("%c ", v)
	}
	fmt.Println()
	fmt.Println(myStrings)

	var intArr = [...] int{1, -1, 9, 90, 12}
	fmt.Println(intArr)
	mIndex := 0
	mValue := intArr[mIndex]
	for i, v := range intArr {
		if mValue < intArr[i] {
			mIndex = i
			mValue = v
		}
	}
	fmt.Printf("index of max value is %d\n", mIndex)
	fmt.Printf("max value is %d\n", mValue)

	sum := 0
	for _, v := range intArr {
		sum += v
	}
	fmt.Println("sum =", sum)
	fmt.Println("mean =", float64(sum)/float64(len(intArr)))

	exec1()
}

func exec1() {
	fmt.Println("# exec1 begin #")
	rand.Seed(time.Now().UnixNano())
	var intArr [5]int
	for i := 0; i < len(intArr); i++ {
		intArr[i] = rand.Intn(100)
	}
	fmt.Println(intArr)
	var resArr [5]int
	for i, _ := range resArr {
		resArr[i] = intArr[len(intArr)-i-1]
	}
	fmt.Println(resArr)
	fmt.Println("# exec1 final #")
}
