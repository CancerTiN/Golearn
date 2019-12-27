package main

import (
	"fmt"
	"time"
)

func putNum(intChan chan int) {
	for i := 1; i <= 80000; i++ {
		intChan <- i
	}
	close(intChan)
}

func primeNum(intChan chan int, primeChan chan int, exitChan chan bool) {
	var flag bool
	for {
		num, ok := <-intChan
		if !ok {
			break
		}
		flag = true
		for i := 2; i < num; i++ {
			if num%i == 0 {
				flag = false
				break
			}
		}
		if flag {
			primeChan <- num
		}
	}
	fmt.Println("The coroutine could not get the data and exited.")
	exitChan <- true
}

func main() {
	inChan := make(chan int, 1000)
	primeChan := make(chan int, 2000)
	exitChan := make(chan bool, 4)

	startTime := float64(time.Now().UnixNano()) / 1e9
	go putNum(inChan)

	for i := 0; i < cap(exitChan); i++ {
		go primeNum(inChan, primeChan, exitChan)
	}

	go func() {
		for i := 0; i < cap(exitChan); i++ {
			<-exitChan
		}
		endTime := float64(time.Now().UnixNano()) / 1e9
		fmt.Printf("Using coroutines takes %f seconds.\n", endTime-startTime)
		close(primeChan)
	}()

	var pn int
	for {
		_, ok := <-primeChan
		if !ok {
			break
		}
		pn++
	}
	fmt.Printf("A total of %v prime numbers were found.\n", pn)
	close(exitChan)

	fmt.Println("The main thread exits.")
}
