package main

import "fmt"

func main() {
	document01()
}

func document01() {
	switch true {
	case true:
		fmt.Println("case1")
	case true:
		fmt.Println("case2")
	}
	/*
		关于switch语句，下边说法都正确：
		1)条件表达式并非必须为常量或整数
		2)switch中，可以出现多个条件相同的case
		3)不需要用break来明确退出一个case
		4)只有在case中明确添加fallthrough关键字，才会继续执行紧跟的下一个case
	*/
}
