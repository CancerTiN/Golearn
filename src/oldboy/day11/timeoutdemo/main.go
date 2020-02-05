package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func worker(ctx context.Context) {
	defer wg.Done()
	for loop := true; loop; {
		fmt.Println("db connecting ...")
		time.Sleep(time.Millisecond * 10)
		select {
		case <-ctx.Done():
			loop = false
		default:
		}
	}
	fmt.Println("worker done!")
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*50)
	wg.Add(1)
	go worker(ctx)
	time.Sleep(time.Second * 5)
	fmt.Println("5 seconds after call worker.")
	fmt.Println("still cancel.")
	cancel()
	wg.Wait()
	fmt.Println("main exit.")
}
