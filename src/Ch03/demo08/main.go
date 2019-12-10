package main

import "fmt"

func main() {
	var price float32 = 89.12
	fmt.Println("price =", price)

	var num1 float32 = -0.00089
	var num2 float64 = -7809656.09
	fmt.Println("num1 =", num1, "num2 =", num2)

	// Mantissa parts may be lost, resulting in loss of accuracy
	var num3 float32 = -123.0000901
	var num4 float64 = -123.0000901
	fmt.Println("num3 =", num3, "num4 =", num4)

	// Golang's floats are declared as float64 by default
	var num5 = 1.1
	fmt.Printf("type of num5 is %T\n", num5)

	num6 := 5.12
	num7 := .123
	fmt.Println("num6 =", num6, "num7 =", num7)

	num8 := 5.1234e2
	num9 := 5.1234e2
	fmt.Println("num8 =", num8, "num9 =", num9)
}
