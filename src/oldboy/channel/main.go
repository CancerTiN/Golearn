package main

import (
	"fmt"
	"time"
)

func useChan(ch chan int) {
	ch <- 1
}

func test01() {
	var ch1 chan int
	ch2 := make(chan int)
	//go useChan(ch1)
	//<-ch1
	go useChan(ch2)
	<-ch2
	fmt.Println(ch1, ch2)
}

func test02() {
	var nilCh chan int
	aCh := make(chan int, 1)
	//go useChan(nilCh)
	go useChan(aCh)
	time.Sleep(time.Second)
	<-aCh
	close(aCh)
	time.Sleep(time.Second)
	res := <-aCh
	//<-nilCh
	fmt.Println(nilCh, aCh, res)
}

func main() {
	switch 2 {
	case 1:
		test01()
	case 2:
		test02()
	}
	document01()
}

func document01() {
	// 无缓冲的channel是同步的，而有缓冲的channel是非同步的
}
