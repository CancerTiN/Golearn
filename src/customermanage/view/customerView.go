package main

import (
	"fmt"
	"strings"
)

type customerView struct {

}

func (cv *customerView) mainMenu() {
	for {
		fmt.Println(strings.Repeat("-", 64))

		fmt.Println("Customer Information Management Software")

	}
}

func main() {
	fmt.Println("ok!")
}
