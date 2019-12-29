package main

import (
	"fmt"
	"math/rand"
	"reflect"
)

func reflectTest01(i interface{}) {
	rTyp := reflect.TypeOf(i)
	fmt.Println("rTyp =", rTyp)
	fmt.Printf("type(rTyp) = %T\n", rTyp)

	rVal := reflect.ValueOf(i)
	fmt.Println("rVal =", rVal)
	fmt.Printf("type(rVal) = %T\n", rVal)

	fmt.Println("rVal.Int() =", rVal.Int())
	fmt.Println("rVal.Int()+rand.Int63() =", rVal.Int()+rand.Int63())
	fmt.Printf("type(rVal.Int()) = %T\n", rVal.Int())

	iV := rVal.Interface()
	fmt.Println("iV =", iV)
	fmt.Printf("type(iV) = %T\n", iV)

	num := iV.(int)
	fmt.Println("num =", num)
	fmt.Println("num+rand.Int() =", num+rand.Int())
	fmt.Printf("type(num) = %T\n", num)
}

func reflectTest02(i interface{}) {
	rTyp := reflect.TypeOf(i)
	fmt.Println("rTyp =", rTyp)
	fmt.Printf("type(rTyp) = %T\n", rTyp)

	rVal := reflect.ValueOf(i)
	fmt.Println("rVal =", rVal)
	fmt.Printf("type(rVal) = %T\n", rVal)

	fmt.Println("rTyp.Kind() =", rTyp.Kind())
	fmt.Println("rVal.Kind() =", rVal.Kind())
	fmt.Println("(rTyp.Kind() == rVal.Kind()) =", rTyp.Kind() == rVal.Kind())

	iV := rVal.Interface()
	fmt.Println("iV =", iV)
	fmt.Printf("type(iV) = %T\n", iV)

	switch iV.(type) {
	case Student:
		s, ok := iV.(Student)
		if ok {
			fmt.Println("s =", s)
			fmt.Printf("type(s) = %T\n", s)
			fmt.Println("s.Name =", s.Name)
			fmt.Println("s.Age =", s.Age)
		}
	case Monster:
		m, ok := iV.(Monster)
		if ok {
			fmt.Println("m.Name =", m.Name)
			fmt.Println("m.Age =", m.Age)
		}
	}
}

type Student struct {
	Name string
	Age  int
}

type Monster struct {
	Name string
	Age  int
}

func main() {
	if false {
		var num int = 100
		reflectTest01(num)
	}

	if true {
		s := Student{Name: "tom", Age: 20}
		reflectTest02(s)
	}
}
