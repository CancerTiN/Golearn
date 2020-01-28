package main

import "fmt"

func test01() {
	中文变量 := "变量内容"
	fmt.Println("中文变量：", 中文变量)
	/*
		对于变量声明下面的说法都是正确的
		1)变量名的首字母不可以是数字
		2)变量名可以为中文
		3)变量名明中不能包含$
		4)变量名不能为关键字
	*/
}

func test02() {
	// var x1 = nil
	var x2 interface{} = nil
	// var x3 string = nil
	var x4 error = nil
	fmt.Println(x2, x4)
}

func test03() {
	var m map[string]int
	var f func(int) int
	var ch <-chan int
	fmt.Println(m, f, ch)
}

func main() {
	switch 3 {
	case 1:
		test01()
	case 2:
		test02()
	case 3:
		test03()
	}
}
