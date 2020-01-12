package main

import (
	"errors"
	"fmt"
	"os"
)

type Queue struct {
	maxSize int
	array   []int
	front   int
	rear    int
}

func initQueue(maxSize int) *Queue {
	return &Queue{
		maxSize: maxSize,
		array:   make([]int, maxSize),
		front:   -1,
		rear:    -1,
	}
}

func (q *Queue) AddQueue(val int) (err error) {
	if q.rear == q.maxSize-1 {
		return errors.New("queue full")
	}
	q.rear++
	q.array[q.rear] = val
	return
}

func (q *Queue) ShowQueue() {
	if q.front == q.rear {
		fmt.Println("The queue is currently empty.")
		return
	}
	fmt.Println("The current situation of the queue is:")
	for i := q.front + 1; i <= q.rear; i++ {
		fmt.Printf("array[%d] = %d\n", i, q.array[i])
	}
}

func main() {
	q := initQueue(5)

	var key string
	var val int
	for {
		fmt.Println("Please enter an instruction to operate the queue.")
		fmt.Scanln(&key)

		switch key {
		case "add":
			fmt.Print("Enter the number to be queued:")
			fmt.Scanln(&val)
			err := q.AddQueue(val)
			if err != nil {
				fmt.Println("Successfully joined the queue.")
			} else {
				fmt.Println(err.Error())
			}
		case "get":
			fmt.Println("get")
		case "show":
			q.ShowQueue()
		case "exit":
			os.Exit(0)
		default:
			fmt.Println("The command entered does not exist.")
		}
	}
}
