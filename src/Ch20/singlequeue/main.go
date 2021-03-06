package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
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

func (q *Queue) Add(val int) (err error) {
	if q.rear == q.maxSize-1 {
		return errors.New("the rear of the queue has reached the border")
	}
	q.rear++
	q.array[q.rear] = val
	return
}

func (q *Queue) Get() (val int, err error) {
	if q.front == q.rear {
		err = errors.New("the queue is currently empty")
		return
	}
	q.front++
	val = q.array[q.front]
	return
}

func (q *Queue) Show() {
	if q.front == q.rear {
		fmt.Println("The queue is currently empty.")
		return
	}
	fmt.Println("The current situation of the queue is:")
	for i := q.front + 1; i <= q.rear; i++ {
		fmt.Printf("queue[%d] = %d\n", i, q.array[i])
	}
}

func main() {
	q := initQueue(2)

	var key string
	var val int

	for {
		fmt.Println("Please enter an instruction to operate the queue")
		fmt.Println("1. Enter <add> to add data to the queue")
		fmt.Println("2. Enter <get> to get data from the queue")
		fmt.Println("3. Enter <show> to display the queue")
		fmt.Println("4. Enter <exit> to exit the operation")
		fmt.Print("Please select (1-4): ")
		key = strings.TrimSpace(input())
		switch key {
		case "add":
			fmt.Print("Enter the number to be queued: ")
			i, err := strconv.ParseInt(strings.TrimSpace(input()), 0, 0)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				val = int(i)
				err = q.Add(val)
				if err != nil {
					fmt.Printf("Fail to add data to the queue: %v.\n", err.Error())
				} else {
					fmt.Println("Succeed in adding data to the queue.")
				}
			}
		case "get":
			val, err := q.Get()
			if err != nil {
				fmt.Printf("Fail to get data from the queue: %v.\n", err.Error())
			} else {
				fmt.Printf("The value taken from the queue is %d.\n", val)
			}
		case "show":
			q.Show()
		case "exit":
			os.Exit(0)
		default:
			fmt.Println("The command entered does not exist.")
		}
	}
}

func input() (lineStr string) {
	var showReadLineReturn bool
	reader := bufio.NewReader(os.Stdin)
	line, isPrefix, err := reader.ReadLine()
	if showReadLineReturn {
		fmt.Println("line:", line)
		fmt.Println("isPrefix:", isPrefix)
		fmt.Println("err:", err)
	}
	lineStr = string(line)
	return
}
