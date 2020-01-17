package main

import (
	"errors"
	"fmt"
)

const n = 5

type Stack struct {
	Max int
	Top int
	Arr [n]int
}

func (s *Stack) Push(val int) (err error) {
	if s.Top == s.Max-1 {
		err = errors.New("stack full")
		return
	}
	s.Top++
	s.Arr[s.Top] = val
	return
}

func (s *Stack) Pop() (val int, err error) {
	if s.Top == -1 {
		err = errors.New("stack empty")
		return
	}
	val = s.Arr[s.Top]
	s.Top--
	return
}

func (s *Stack) List() {
	if s.Top == -1 {
		fmt.Println("Stack empty.")
		return
	}
	fmt.Println("Stack situation:")
	for i := s.Top; i > -1; i-- {
		fmt.Printf("Arr[%d] = %v\n", i, s.Arr[i])
	}
}

func main() {
	stack := &Stack{Max: n, Top: -1}
	for i := 1; i < 7; i++ {
		err := stack.Push(i)
		if err != nil {
			fmt.Printf("stack.Push(%d) error: %v.\n", i, err)
		}
	}

	stack.List()

	for i := 1; i < 7; i++ {
		val, err := stack.Pop()
		if err != nil {
			fmt.Printf("stack.Pop() error: %v.\n", err)
		} else {
			fmt.Println("stack.Pop() return:", val)
		}
		stack.List()
	}
}
