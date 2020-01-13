package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CircleQueue struct {
	maxSize int
	array   []int
	head    int
	tail    int
}

func initCircleQueue(maxSize int) *CircleQueue {
	return &CircleQueue{
		maxSize: maxSize,
		array:   make([]int, maxSize),
		head:    0, // open
		tail:    0, // close
	}
}

func (c *CircleQueue) IsFull() bool {
	return (c.tail+1)%c.maxSize == c.head
}

func (c *CircleQueue) IsEmpty() bool {
	return c.tail == c.head
}

func (c *CircleQueue) Size() int {
	return (c.tail + c.maxSize - c.head) % c.maxSize
}

func (c *CircleQueue) Push(val int) (err error) {
	if c.IsFull() {
		return errors.New("the tail of the queue has reached the border")
	}
	c.array[c.tail] = val
	c.tail = (c.tail + 1) % c.maxSize
	return
}

func (c *CircleQueue) Pop() (val int, err error) {
	if c.IsEmpty() {
		err = errors.New("the queue is currently empty")
		return
	}
	val = c.array[c.head]
	c.head = (c.head + 1) % c.maxSize
	return
}

func (c *CircleQueue) List() {
	size := c.Size()
	if size == 0 {
		fmt.Println("The queue is currently empty.")
		return
	}
	fmt.Println("The current situation of the queue is:")
	tempHead := c.head
	for i := 0; i < size; i++ {
		fmt.Printf("queue[%d] = %d\n", tempHead, c.array[tempHead])
		tempHead = (tempHead + 1) % c.maxSize
	}
}

func main() {
	if false {
		moduloDemo()
	}

	cq := initCircleQueue(5)

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
				err = cq.Push(val)
				if err != nil {
					fmt.Printf("Fail to add data to the queue: %v.\n", err.Error())
				} else {
					fmt.Println("Succeed in adding data to the queue.")
				}
			}
		case "get":
			val, err := cq.Pop()
			if err != nil {
				fmt.Printf("Fail to get data from the queue: %v.\n", err.Error())
			} else {
				fmt.Printf("The value taken from the queue is %d.\n", val)
			}
		case "show":
			cq.List()
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

func moduloDemo() {
	m := 10
	for i := 1; i < m; i++ {
		fmt.Println(i % m)
	}
}
