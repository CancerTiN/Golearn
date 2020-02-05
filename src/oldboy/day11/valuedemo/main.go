package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

type TraceCode string

func worker(ctx context.Context) {
	defer wg.Done()
	key := TraceCode("TRACE_CODE")
	value, ok := ctx.Value(key).(string)
	if !ok {
		fmt.Println("invalid trace code")
	}
	for loop := true; loop; {
		fmt.Println("worker working, trace code:", value)
		time.Sleep(time.Millisecond * 10)
		select {
		case <-ctx.Done():
			loop = false
		default:
		}
	}
	fmt.Println("worker done")
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*50)
	key := TraceCode("TRACE_CODE")
	value := "12512312234"
	ctx = context.WithValue(ctx, key, value)
	wg.Add(1)
	go worker(ctx)
	time.Sleep(time.Second * 5)
	fmt.Println("5 seconds after worker start")
	fmt.Println("still cancel")
	cancel()
	wg.Wait()
	fmt.Println("main exit")
}
