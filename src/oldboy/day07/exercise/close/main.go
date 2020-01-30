package main

import (
	"fmt"
	"math/rand"
)

var (
	nJobs      = 100
	nWorkers   = 24
	jobChan    = make(chan int64, 1)
	resultChan = make(chan int64, nJobs)
	doneChan   = make(chan bool, nWorkers)
)

func producer(jobChan chan<- int64) {
	for i := 0; i < nJobs; i++ {
		jobChan <- rand.Int63()
	}
	close(jobChan)
	fmt.Println("---- close jobChan ----")
}

func worker(id int, jobChan <-chan int64, resultChan chan<- int64) {
	for {
		j, ok := <-jobChan
		if !ok {
			break
		}
		r := digitSumOf(j)
		fmt.Printf("worker<%d> convert %d to %d\n", id, j, r)
		resultChan <- r
	}
	doneChan <- true
}

func digitSumOf(i int64) (sum int64) {
	//var (
	//	residue   = i
	//	remainder int64
	//	flag      bool
	//)
	//for !flag {
	//	remainder = residue % 10
	//	if residue/10 != 0 {
	//		residue /= 10
	//	} else {
	//		flag = true
	//	}
	//	sum += remainder
	//}
	for r := i; r > 0; r /= 10 {
		sum += r % 10
	}
	return
}

func main() {
	fmt.Println("main start")
	go producer(jobChan)
	for i := 0; i < nWorkers; i++ {
		go worker(i, jobChan, resultChan)
	}
	go func() {
		for {
			if cap(doneChan) == len(doneChan) {
				close(resultChan)
				fmt.Println("---- close resultChan ----")
				break
			}
		}
	}()
	for {
		r, ok := <-resultChan
		if !ok {
			break
		}
		fmt.Println("main export:", r)
	}
	fmt.Println("main end")
}
