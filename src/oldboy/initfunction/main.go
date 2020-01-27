package main

import "fmt"

func init() {
	fmt.Println(1)
}

func init() {
	fmt.Println(2)
}

func main() {
	document01()
}

func document01() {
	/*
		关于init函数，下边说法均正确：
		1)一个包中，可以包含多个init函数
		2)程序编译时，先执行导入包的init函数，再执行本包内的init函数
		3)main包中，能有init函数
		4)init函数不可以被其它函数调用
	*/
}
