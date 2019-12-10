package main

import (
	"fmt"
	"unsafe"
)

// Demonstrate the use of bool types
func main() {
	var b = false
	fmt.Println("b =", b)
	fmt.Println("footprint of b is", unsafe.Sizeof(b))
}
