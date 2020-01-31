package main

import (
	"fmt"
	"time"
)

var ch1 chan int
var ch2 chan int
var chs = []chan int{ch1, ch2}
var numbers = []int{1, 2, 3, 4, 5}

func main() {
	switch 2 {
	case 1:
		demo01()
	case 2:
		demo02()
	}
}

func demo01() {
	select {
	case getChan(0) <- getNumber(2):
		fmt.Println("1th case is selected.")
	case getChan(1) <- getNumber(3):
		fmt.Println("2nd case is selected.")
	default:
		fmt.Println("default!.")
	}
	fmt.Println(ch1, ch2)
	ch1 = make(chan int)
	go func() {
		time.Sleep(time.Millisecond * 500)
		fmt.Println("Send 1 to ch1.")
		ch1 <- 1
	}()
	<-ch1
	fmt.Println("Received 1 from ch1.")
	close(ch1)
	//defer func() {
	//	recover()
	//}()
	close(ch2) // panic: close of nil channel
}

func getChan(i int) chan int {
	fmt.Printf("chs[%d]\n", i)
	return chs[i]
}

func getNumber(i int) int {
	fmt.Printf("numbers[%d]\n", i)
	return numbers[i]
}

func demo02() {
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)
	ch1 <- 3
	ch2 <- 5
	select {
	case <-ch1:
		fmt.Println("ch1 selected.")
		break
		fmt.Println("ch1 selected after break.")
	case <-ch2:
		fmt.Println("ch2 selected.")
		fmt.Println("ch2 selected without break.")
	}
}
