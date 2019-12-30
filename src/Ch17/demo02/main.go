package main

import (
	"fmt"
	"reflect"
)

func reflect01(intPtr interface{}) {
	rVal := reflect.ValueOf(intPtr)
	fmt.Println("rVal =", rVal)
	fmt.Println("rVal.Kind() =", rVal.Kind())
	fmt.Printf("type(rVal) = %T\n", rVal)

	rValElm := rVal.Elem()
	fmt.Println("rValElm =", rValElm)
	fmt.Println("rValElm.Kind() =", rValElm.Kind())
	fmt.Printf("type(rValElm) = %T\n", rValElm)

	rValElm.SetInt(20)
}

func main() {
	var num int = 10
	fmt.Println("num =", num)
	reflect01(&num)
	fmt.Println("num =", num)
}
