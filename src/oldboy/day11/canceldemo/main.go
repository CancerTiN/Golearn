package main

import (
	"context"
	"fmt"
	"time"
)

func gen(ctx context.Context) (dst chan int) {
	dst = make(chan int)
	n := 1
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("goroutine exit")
				return
			case dst <- n:
				n++
			}
		}
	}()
	fmt.Println("gen exit")
	return
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer func() {
		time.Sleep(time.Nanosecond)
		fmt.Println("1 nanosecond after main exit")
	}()
	defer cancel()
	dst := gen(ctx)
	for n := range dst {
		fmt.Println("main n:", n)
		if n == 5 {
			break
		}
	}
	fmt.Println("main exit")
}
