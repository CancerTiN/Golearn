package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	intChan := make(chan int, 50)
	exitChan := make(chan bool, 1)

	go writeData(intChan)
	go readData(intChan, exitChan)

	//if b := <-exitChan; b {
	//	return
	//}

	for {
		v, ok := <-exitChan
		fmt.Println("ok =", ok)
		if !ok {
			fmt.Println("v =", v)
			break
		}
	}
}

func writeData(intChan chan int) {
	for i := 1; i <= 50; i++ {
		intChan <- i
		fmt.Printf("Write data %v.\n", i)
		rand.Seed(time.Now().UnixNano())
		time.Sleep(time.Second / time.Duration(rand.Intn(1000)+1))
	}
	close(intChan)
	fmt.Println("Close intChan.")
}

func readData(intChan chan int, exitChan chan bool) {
	for {
		v, ok := <-intChan
		if !ok {
			break
		}
		rand.Seed(time.Now().UnixNano())
		time.Sleep(time.Second / time.Duration(rand.Intn(1000)+1))
		fmt.Printf("Read data %v.\n", v)
	}
	exitChan <- true
	close(exitChan)
	fmt.Println("Close exitChan.")
}
