package main

import "fmt"

type myFunType func(int, int) int

func main() {
	f := getSum
	fmt.Printf("type of f is %T\n", f)
	fmt.Printf("type of getSum is %T\n", getSum)

	res1 := f(10, 40)
	fmt.Println("res1 =", res1)

	res2 := myFun(getSum, 50, 60)
	fmt.Println("res2 =", res2)

	type myInt int
	var num1 myInt
	var num2 int
	num1 = 40
	num2 = 50
	fmt.Println("num1 =", num1)
	fmt.Println("num2 =", num2)

	res3 := myFunTwo(getSum, 500, 600)
	fmt.Println("res3 =", res3)

	a, b := getSumAndSub(1, 2)
	fmt.Printf("a = %v; b = %v\n", a, b)
}

func getSum(n1 int, n2 int) int {
	return n1 + n2
}

func myFun(funvar func(int, int) int, num1 int, num2 int) int {
	return funvar(num1, num2)
}

func myFunTwo(funvar myFunType, num1 int, num2 int) int {
	return funvar(num1, num2)
}

func getSumAndSub(n1 int, n2 int) (sum int, sub int) {
	sum = n1 + n2
	sub = n1 - n2
	return
}
