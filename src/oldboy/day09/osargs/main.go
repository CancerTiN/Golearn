package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("%#v\n", os.Args)
	fmt.Printf("Type of os.Args is %T.\n", os.Args)
}
