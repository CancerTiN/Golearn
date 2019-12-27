package main

import "fmt"

func main() {
	intChan := make(chan int, 3)
	intChan <- 100
	intChan <- 200
	close(intChan)
	// intChan <- 300 // panic: send on closed channel
	fmt.Println("ok!")
	n1 := <-intChan
	fmt.Println("n1 =", n1)

	hundredIntChan := make(chan int, 100)
	for i := 0; i < cap(hundredIntChan); i++ {
		hundredIntChan <- i * 2
	}
	close(hundredIntChan)
	for v := range hundredIntChan {
		fmt.Println("v =", v)
	}
}
