package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("There are %v parameters on the command line.\n", len(os.Args))
	for i, v := range os.Args {
		fmt.Printf("os.Args[%v] = %v\n", i, v)
	}
}
