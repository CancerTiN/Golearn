package main

import "fmt"

func main() {
	var num int = 9
	fmt.Println("num =", num)

	const tax int = 0
	fmt.Println("tax =", tax)

	const z = 9 / 3
	// const y = num / 9 // Const initializer 'num' is not a constant
	fmt.Println("z =", z)

	const (
		a = iota
		b
		c
		d
		e // Declared constants may not be used.
	)
	fmt.Println(a, b, c, d)

	const (
		f = iota // 0
		g        // 1
		_
		i = iota // 3
		j        // 4
	)
	fmt.Println(f, g, i, j)
}
