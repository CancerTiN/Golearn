package main

import (
	"fmt"
	"math/rand"
	"time"
)

const n = 100

var ch = make(chan int, n)

func hello(i int) {
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(n)*rand.Intn(n)))
	ch <- i
	fmt.Printf("Hello, goroutine%v!\n", i)
}

func main() {
	for i := 0; i < n; i++ {
		go hello(i)
	}
	fmt.Printf("Main run with %#v channels!\n", len(ch))
	time.Sleep(time.Second)
	fmt.Printf("Main end with %#v channels!\n", len(ch))
}
