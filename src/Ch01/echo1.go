package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		ns := sep + os.Args[i]
		s += ns
		sep = " "
	}
	fmt.Println(s)
}
