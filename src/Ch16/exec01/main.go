package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	testNum := 200000
	chanNum := 4
	intChan := make(chan int, 1000)
	primeChan := make(chan int, testNum)
	exitChan := make(chan int, chanNum)

	go putNumber(testNum, intChan)
	for i := 0; i < chanNum; i++ {
		go pullAndPushPrime(intChan, primeChan, exitChan, i)
	}

	epoch := 0
	for {
		time.Sleep(time.Second)
		epoch++
		fmt.Println("epoch =", epoch)
		if len(exitChan) == 4 {
			close(exitChan)
			break
		}
	}

	for {
		v, ok := <-exitChan
		if !ok {
			close(primeChan)
			break
		}
		fmt.Println("v =", v)
	}

	for p := range primeChan {
		fmt.Printf("%v is a prime number.\n", p)
	}

	rand.Seed(time.Now().UnixNano())
	time.Sleep(time.Second / time.Duration(rand.Intn(100)+1))
}

func putNumber(testNum int, intChan chan int) {
	for n := 1; n <= testNum; n++ {
		//rand.Seed(time.Now().UnixNano())
		//time.Sleep(time.Second / time.Duration(rand.Intn(n*100)+1))
		intChan <- n
	}
	close(intChan)
	fmt.Println("Complete function putNumber.")
}

func pullAndPushPrime(intChan, primeChan, exitChan chan int, mark int) {
	for {
		n, ok := <-intChan
		if !ok {
			exitChan <- mark
			break
		} else if isPrime(n) {
			primeChan <- n
		}
	}
	fmt.Printf("Complete function pullAndPushPrime %v.\n", mark)
}

func isPrime(n int) bool {
	if n < 4 {
		return true
	}
	if n%2 == 0 || n%3 == 0 || n%5 == 0 {
		return false
	}
	var flag bool
	for i := 2; i < n/2+1; i++ {
		if n%i == 0 {
			flag = true
			break
		}
	}
	if flag {
		return false
	} else {
		return true
	}
}
