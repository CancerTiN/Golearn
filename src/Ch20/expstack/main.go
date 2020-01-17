package main

import (
	"errors"
	"fmt"
	"strconv"
)

const n = 20

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

func (s *Stack) IsOpe(val int) bool {
	// [42 *] [43 +] [45 -] [47 /]
	if val == 42 || val == 43 || val == 45 || val == 47 {
		return true
	} else {
		return false
	}
}

func (s *Stack) Cal(num1, num2, ope int) (res int, err error) {
	// [42 *] [43 +] [45 -] [47 /]
	//n1, err := strconv.ParseInt(string(byte(num1)), 0, 0)
	//n2, err := strconv.ParseInt(string(byte(num2)), 0, 0)
	//num1, num2 = int(n1), int(n2)
	fmt.Println(num1, num2)
	switch ope {
	case 42:
		res = num2 * num1
	case 43:
		res = num2 + num1
	case 45:
		res = num2 - num1
	case 47:
		res = num2 / num1
	default:
		err = errors.New("unsupported operator")
	}
	return
}

func (s *Stack) PriorityOf(ope int) (pri int) {
	if ope == 42 || ope == 47 {
		pri = 1
	} else if ope == 43 || ope == 45 {
		pri = 0
	}
	return
}

func main() {
	fmt.Println('+', '-', '*', '/')

	var (
		err     error
		num1    int
		num2    int
		ope     int
		result  int
		val64   int64
		exp     string
		keepStr string
	)

	fmt.Println(keepStr)

	numStack := &Stack{Max: n, Top: -1}
	opeStack := &Stack{Max: n, Top: -1}

	exp = "3+2*6-2"
	exp = "03+0300*0006-0400"

	fmt.Println("exp =", exp)

	for index := 0; index < len(exp); index++ {
		fmt.Println("index =", index)

		charUnit8 := exp[index] // uint8
		charInt := int(charUnit8)

		fmt.Println("string(charInt) =", string(charInt))

		if opeStack.IsOpe(charInt) {
			if opeStack.Top == -1 {
				err = opeStack.Push(charInt)
				if err != nil {
					fmt.Printf("opeStack.Push(%d) error: %v.", charInt, err)
				}
			} else {
				if opeStack.PriorityOf(opeStack.Arr[opeStack.Top]) >= opeStack.PriorityOf(charInt) {
					num1, err = numStack.Pop()
					num2, err = numStack.Pop()
					ope, err = opeStack.Pop()
					result, err = opeStack.Cal(num1, num2, ope)
					err = numStack.Push(result)
					err = opeStack.Push(charInt)
				} else {
					err = opeStack.Push(charInt)
				}
			}
		} else {
			if index == len(exp)-1 {
				fmt.Println("lastStr =", string(charUnit8))
				val64, err = strconv.ParseInt(string(charUnit8), 10, 0)
				err = numStack.Push(int(val64))
				continue
			}
			keepStr += string(charUnit8)
			fmt.Println("keepStr =", keepStr)
			for nextIndex := index + 1; nextIndex <= len(exp); nextIndex++ {
				if nextIndex == len(exp) {
					fmt.Println("nextIsEnd: keepStr =", keepStr)
					val64, err = strconv.ParseInt(keepStr, 10, 0)
					err = numStack.Push(int(val64))
					keepStr = ""
					index = nextIndex - 1
					break
				}
				nextCharUnit8 := exp[nextIndex]
				if opeStack.IsOpe(int(nextCharUnit8)) {
					fmt.Println("nextIsOpe: keepStr =", keepStr)
					val64, err = strconv.ParseInt(keepStr, 10, 0)
					err = numStack.Push(int(val64))
					keepStr = ""
					index = nextIndex - 1
					break
				} else {
					keepStr += string(nextCharUnit8)
				}
			}
			//val64, err = strconv.ParseInt(string(charUnit8), 0, 0)
			//err = numStack.Push(int(val64))
			//if err != nil {
			//	fmt.Printf("numStack.Push(%d) error: %v.", charInt, err)
			//}
		}
	}

	for opeStack.Top > -1 {
		num1, err = numStack.Pop()
		num2, err = numStack.Pop()
		ope, err = opeStack.Pop()
		result, err = opeStack.Cal(num1, num2, ope)
		err = numStack.Push(result)
	}

	val, err := numStack.Pop()
	if err != nil {
		fmt.Printf("numStack.Pop() error: %v.", err)
	} else {
		fmt.Println(exp, "=", val)
	}
}
