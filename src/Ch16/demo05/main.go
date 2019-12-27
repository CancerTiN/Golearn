package main

import "fmt"

type Cat struct {
	Name string
	Age  int
}

func main() {
	allChan := make(chan interface{}, 3)

	allChan <- 10
	allChan <- "tom jack"
	cat := Cat{"Kitten", 4}
	allChan <- cat

	<-allChan
	<-allChan
	newCat := <-allChan
	fmt.Println("newCat =", newCat)
	fmt.Printf("type(newCat) = %T\n", newCat)
	// fmt.Println(newCat.Name) // newCat.Name undefined (type interface {} is interface with no methods)
	if c, ok := newCat.(Cat); ok {
		fmt.Println("c.Name =", c.Name)
	}
}
