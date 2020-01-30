package main

import "fmt"

const n = 100

func f1(ch chan int) {
	for i := 0; i < n; i++ {
		ch <- i
	}
	close(ch)
}

func f2(ch1, ch2 chan int) {
	for i := 0; i < n; i++ {
		x := <-ch1
		ch2 <- x * x
	}
	close(ch2)
}

func f3(ch1, ch2 chan int) {
	for {
		x, ok := <-ch1
		if !ok {
			break
		}
		ch2 <- x * x
	}
	close(ch2)
}

func main() {
	var ch1 = make(chan int)
	var ch2 = make(chan int, n)
	go f1(ch1)
	switch 3 {
	case 2:
		go f2(ch1, ch2)
	case 3:
		go f3(ch1, ch2)
	}
	switch "range" {
	case "range":
		for r := range ch2 {
			fmt.Println(r)
		}
	default:
		for i := 0; i < n; i++ {
			fmt.Println(<-ch2)
		}
	}
}
