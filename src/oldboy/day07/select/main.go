package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	for i := 0; i < 10; i++ {
		select {
		case x := <-ch:
			fmt.Printf("Receive %d from channel.\n", x)
		case ch <- i:
			fmt.Printf("Send %d to channel.\n", i)
		}
	}
	close(ch)
}
