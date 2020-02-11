package main

import (
	"fmt"
)

func isValid(s string) bool {
	var (
		slice []int32
		m     = map[int32]int32{')': '(', '}': '{', ']': '['}
	)
	for _, b := range s {
		if b == '(' || b == '{' || b == '[' {
			slice = append(slice, b)
		} else {
			if len(slice) == 0 {
				return false
			}
			if m[b] == slice[len(slice)-1] {
				slice = slice[:len(slice)-1]
			} else {
				return false
			}
		}
	}
	if len(slice) > 0 {
		return false
	}
	return true
}

func main() {
	for _, s := range []string{"()", "()[]{}", "(]", "([)]", "{[]}"} {
		fmt.Println(s, isValid(s))
	}
}
