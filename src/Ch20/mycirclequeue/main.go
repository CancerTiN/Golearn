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
	size    int
}

func initCircleQueue(maxSize int) *CircleQueue {
	return &CircleQueue{
		maxSize: maxSize,
		array:   make([]int, maxSize),
	}
}

func (c *CircleQueue) IsFull() bool {
	return c.size == c.maxSize
}

func (c *CircleQueue) IsEmpty() bool {
	return c.size == 0
}

func (c *CircleQueue) Push(val int) (err error) {
	if c.IsFull() {
		return errors.New("the tail of the queue has reached the border")
	}
	c.array[c.tail] = val
	c.tail = (c.tail + 1) % c.maxSize
	c.size++
	return
}

func (c *CircleQueue) Pop() (val int, err error) {
	if c.IsEmpty() {
		err = errors.New("the queue is currently empty")
		return
	}
	val = c.array[c.head]
	c.head = (c.head + 1) % c.maxSize
	c.size--
	return
}

func (c *CircleQueue) List() {
	if c.size == 0 {
		fmt.Println("The queue is currently empty.")
		return
	}
	fmt.Println("The current situation of the queue is:")
	i := c.head
	for j := 0; j < c.size; j++ {
		fmt.Printf("queue[%d] = %d\n", i, c.array[i])
		i = (i + 1) % c.maxSize
	}
}

func main() {
	if false {
		moduloDemo()
	}

	cq := initCircleQueue(3)

	var key string
	var val int

	for {
		fmt.Println("Please enter an instruction to operate the queue")
		fmt.Println("1. Enter <add> to add data to the queue")
		fmt.Println("2. Enter <get> to get data from the queue")
		fmt.Println("3. Enter <show> to display the queue")
		fmt.Println("4. Enter <exit> to exit the operation")
		fmt.Print("Please select {add, get, show, exit}: ")
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
