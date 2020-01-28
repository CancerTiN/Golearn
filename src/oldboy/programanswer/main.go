package main

import (
	"fmt"
	"sync"
	"time"
)

func isPalindrome(s string) bool {
	runeSlice := []rune(s)
	if len(runeSlice) < 1 {
		return false
	}
	for i, j := 0, len(runeSlice)-1; i <= j; i, j = i+1, j-1 {
		if runeSlice[i] == runeSlice[j] {
			continue
		} else {
			return false
		}
	}
	return true
}

func question01() {
	// 写一个函数判断一个字符串是否是回文。
	s := "上海自来水来自海上"
	s = "abba"
	fmt.Println(isPalindrome(s))
}

func sayOdd(group *sync.WaitGroup, chOdd, chEven chan int) {
	defer group.Done()
	for i := 1; i <= 10; i++ {
		fmt.Println("sayOdd:", <-chOdd)
		chEven <- 2 * i
	}
}

func sayEven(group *sync.WaitGroup, chOdd, chEven chan int) {
	defer group.Done()
	for i := 1; i <= 10; i++ {
		chOdd <- 2*i - 1
		fmt.Println("sayEven", <-chEven)
	}
}

func question02() {
	// 使用2个协程交替执行，使其能顺序输出1-20的自然数，一个子协程输出奇数，另一个输出偶数。
	group := &sync.WaitGroup{}
	chOdd := make(chan int)
	chEven := make(chan int)
	group.Add(2)
	go sayOdd(group, chOdd, chEven)
	go sayEven(group, chOdd, chEven)
	group.Wait()
}

func solve(total int) (result [][3]int) {
	a, b := 0, total/2
	c := total - a - b
	for a := 0; a <= total/2; a++ {
		for b := total / 2; b >= 0; b-- {
			c = total - a - b
			if a*a+b*b == c*c {
				result = append(result, [3]int{a, b, c})
			}
		}
		c = total - a - b
	}
	return
}

func question03() {
	// 如果a+b+c==1000，且a*a+b*b=c*c，要求a、b和c为自然数，如果求出所有a、b和c可能的组合？
	startTime := time.Now().UnixNano()
	r := solve(1000)
	endTime := time.Now().UnixNano()
	fmt.Printf("Solver takes %v nanoseconds.\n", endTime-startTime)
	for _, s := range r {
		fmt.Println(s)
	}
}

func main() {
	switch 3 {
	case 1:
		question01()
	case 2:
		question02()
	case 3:
		question03()
	}
}
