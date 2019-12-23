package main

import "fmt"

type Student struct{}

func TypeJudge(items ...interface{}) {
	for i, x := range items {
		switch x.(type) {
		case bool:
			fmt.Printf("The type of argument %d is boolean, its value is %v.\n", i, x)
		case float32:
			fmt.Printf("The type of argument %d is float32, its value is %v.\n", i, x)
		case float64:
			fmt.Printf("The type of argument %d is float64, its value is %v.\n", i, x)
		case int, int32, int64:
			fmt.Printf("The type of argument %d is integer, its value is %v.\n", i, x)
		case string:
			fmt.Printf("The type of argument %d is string, its value is %v.\n", i, x)
		case Student:
			fmt.Printf("The type of argument %d is Student, its value is %v.\n", i, x)
		case *Student:
			fmt.Printf("The type of argument %d is *Student, its value is %v.\n", i, x)
		default:
			fmt.Printf("The type of argument %d cannot be determined, its value is %v.\n", i, x)
		}
	}
}

func main() {
	var n1 float32 = 1.1
	var n2 float64 = 2.3
	var n3 int32 = 30
	var name string = "tom"
	address := "beijing"
	n4 := 300

	s1 := Student{}
	s2 := &Student{}

	TypeJudge(n1, n2, n3, name, address, n4, s1, s2)
}
