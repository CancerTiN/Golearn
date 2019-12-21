package main

type AInterface interface {
	Test01()
	Test02()
}

type BInterface interface {
	Test01()
	Test03()
}

// duplicate method Test01 // error
// type CInterface interface {
// 	AInterface
// 	BInterface
// }

func main() {
	// var c CInterface
	// fmt.Println("c =", c)
}
