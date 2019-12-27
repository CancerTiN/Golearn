package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	go test()
	for i := 0; i < 10; i++ {
		fmt.Printf("main(%v) Hello, golang!\n", strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}

func test() {
	for i := 0; i < 20; i++ {
		fmt.Printf("test(%v) Hello, golang!\n", strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}
