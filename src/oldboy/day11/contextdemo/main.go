package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func ff(ctx context.Context) {
	defer wg.Done()
	for loop := true; loop; {
		fmt.Println("ff time:", time.Now().UnixNano())
		time.Sleep(time.Millisecond * 500)
		select {
		case <-ctx.Done():
			loop = false
		default:
		}
	}
	time.Sleep(time.Second * 3)
	fmt.Println("Three seconds after the loop ends.")
}

func f(ctx context.Context) {
	defer wg.Done()
	wg.Add(1)
	go ff(ctx)
	for loop := true; loop; {
		fmt.Println(time.Now())
		time.Sleep(time.Second)
		select {
		case <-ctx.Done():
			loop = false
		default:
		}
	}
	time.Sleep(time.Second)
	fmt.Println("One second after the loop ends.")
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go f(ctx)
	time.Sleep(time.Second * 5)
	cancel()
	wg.Wait()
}
