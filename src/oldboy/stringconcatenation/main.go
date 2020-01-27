package main

import "fmt"

func main() {
	test01()
}

func test01() {
	str1 := "abc" + "123"
	fmt.Println("str1:", str1)
	str2 := fmt.Sprintf("abc%d", 123)
	fmt.Println("str2:", str2)
}
