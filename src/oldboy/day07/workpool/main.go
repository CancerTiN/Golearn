package main

import (
	"fmt"
	"time"
)

const n = 5

var count int

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("worker<%d> start job<%d>\n", id, j)
		time.Sleep(time.Second)
		results <- j * 2
		count++
		fmt.Printf("worker<%d> end job<%d>\n", id, j)
		if count == n {
			fmt.Println("reach border, close channel")
			close(results)
			break
		}
	}
}

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}
	for j := 1; j <= n; j++ {
		jobs <- j
	}
	close(jobs)
	for {
		_, ok := <-results
		if !ok {
			break
		}
	}
}
