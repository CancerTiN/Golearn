package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	d := time.Now().Add(time.Millisecond * 50)
	ctx, cancel := context.WithDeadline(context.Background(), d)

	defer func() {
		fmt.Println("still cancel")
		cancel()
	}()

	select {
	case <-time.After(time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}
