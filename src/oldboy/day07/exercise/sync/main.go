package main

import (
	"fmt"
	"math/rand"
	"time"
)

type result struct {
	job int64
	sum int64
}

func (r *result) Solve() {
	for n := r.job; n > 0; n /= 10 {
		r.sum += n % 10
	}
}

var (
	nSize      = 100
	nWorkers   = 24
	jobChan    = make(chan int64, nSize)
	resultChan = make(chan result, nSize)
)

func producer(ch chan<- int64) {
	for {
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)))
		ch <- rand.Int63()
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)))
	}
}

func worker(jch <-chan int64, rch chan<- result) {
	for {
		var r result
		r.job = <-jch
		r.Solve()
		rch <- r
	}
}

func main() {
	go producer(jobChan)
	for i := 0; i < nWorkers; i++ {
		go worker(jobChan, resultChan)
	}
	var count int
	for r := range resultChan {
		count++
		fmt.Println(count, r)
	}
}
