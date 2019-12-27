package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	myMap = make(map[int]int)
	lock  sync.Mutex
)

func main() {
	for i := 1; i <= 20; i++ {
		// go test(i) // fatal error: concurrent map writes
		go test(i)
	}
	fmt.Println("After goroutine.")
	lock.Lock()
	fmt.Println("Before Print.")
	for i, v := range myMap {
		fmt.Printf("myMap[%v] = %v\n", i, v)
	}
	fmt.Println("After Print.")
	lock.Unlock()
}

func test(n int) {
	r := 1
	for i := 1; i <= n; i++ {
		r *= i
	}
	lock.Lock()
	time.Sleep(time.Second)
	myMap[n] = r
	lock.Unlock()
}
