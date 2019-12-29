package main

import (
	"fmt"
	"reflect"
)

func test01(f interface{}) {
	rVal := reflect.ValueOf(f)
	fmt.Println("rVal.Type() =", rVal.Type())
	fmt.Println("rVal.Kind() =", rVal.Kind())

	iV := rVal.Interface()
	fmt.Println("iV =", iV)

	switch iV.(type) {
	case float64:
		v, ok := iV.(float64)
		if ok {
			fmt.Println("v =", v)
		}
	}
}

func main() {
	var v float64 = 1.2
	test01(v)
}
