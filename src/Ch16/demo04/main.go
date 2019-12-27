package main

import "fmt"

func main() {
	var intChan chan int
	intChan = make(chan int, 3)

	fmt.Println("intChan =", intChan)
	fmt.Println("len(intChan) =", len(intChan))
	fmt.Println("cap(intChan) =", cap(intChan))

	intChan <- 10
	num := 211
	intChan <- num
	intChan <- 50

	fmt.Println("intChan =", intChan)
	fmt.Println("len(intChan) =", len(intChan))
	fmt.Println("cap(intChan) =", cap(intChan))

	// intChan <- 98 // fatal error: all goroutines are asleep - deadlock!
	<-intChan
	intChan <- 98

	var n1 int
	n1 = <-intChan
	fmt.Println("n1 =", n1)
	fmt.Println("len(intChan) =", len(intChan))
	fmt.Println("cap(intChan) =", cap(intChan))

	n2 := <-intChan
	fmt.Println("n2 =", n2)
	n3 := <-intChan
	fmt.Println("n3 =", n3)
	n4 := <-intChan // fatal error: all goroutines are asleep - deadlock!
	fmt.Println("n4 =", n4)
}
