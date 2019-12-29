package main

import (
	"fmt"
	"reflect"
)

func main() {
	var str string = "tom"
	fmt.Println("str =", str)
	if false {
		fs := reflect.ValueOf(str)
		fs.SetString("jack") // panic: reflect: reflect.flag.mustBeAssignable using unaddressable value
	}
	rValPtr := reflect.ValueOf(&str)
	rValElm := rValPtr.Elem()
	rValElm.SetString("jack")
	fmt.Println("str =", str)
}
