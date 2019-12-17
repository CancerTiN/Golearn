package main

import (
	"fmt"
	"runtime"
	"unsafe"
)

type Cat struct {
	Name  string
	Age   int
	Color string
	Hobby string
}

func main() {
	test01()
	test02()
	test03()
}

func test01() {
	framePC, _, _, _ := runtime.Caller(0)
	fmt.Printf("Enter the function %v.\n", runtime.FuncForPC(framePC).Name())

	cat11Name := "xiaobai"
	cat1Age := 3
	cat1Color := "white"
	cat21Name := "xiaohua"
	cat2Age := 100
	cat2Color := "color"
	fmt.Println(cat11Name, cat1Age, cat1Color)
	fmt.Println(cat21Name, cat2Age, cat2Color)
}

func test02() {
	framePC, _, _, _ := runtime.Caller(0)
	fmt.Printf("Enter the function %v.\n", runtime.FuncForPC(framePC).Name())

	catNames := [...]string{"xiaobai", "xiaohua"}
	catAges := [...]int{3, 100}
	catColors := [...]string{"white", "color"}
	fmt.Println(catNames, catAges, catColors)
}

func test03() {
	framePC, _, _, _ := runtime.Caller(0)
	fmt.Printf("Enter the function %v.\n", runtime.FuncForPC(framePC).Name())

	var cat1 Cat
	fmt.Printf("The address of cat1 is %p.\n", &cat1)
	cat1.Name = "xiaobai"
	cat1.Age = 3
	cat1.Color = "white"
	cat1.Hobby = "Eat fish"
	fmt.Println("cat1 =", cat1)
	fmt.Println("The cat1's information is as follows:")
	fmt.Println("\tname =", cat1.Name)
	fmt.Println("\tage =", cat1.Age)
	fmt.Println("\tcolor =", cat1.Color)
	fmt.Println("\thobby =", cat1.Hobby)
	fmt.Printf("The address of cat1 name is %p.\n", &cat1.Name)   // 16 bytes
	fmt.Printf("The address of cat1 age is %p.\n", &cat1.Age)     // 8 bytes
	fmt.Printf("The address of cat1 color is %p.\n", &cat1.Color) // 16 bytes
	fmt.Printf("The address of cat1 hobby is %p.\n", &cat1.Hobby) // 16 bytes
	fmt.Printf("The size of cat1 name is %v.\n", unsafe.Sizeof(cat1.Name))
	fmt.Printf("The size of cat1 age is %v.\n", unsafe.Sizeof(cat1.Age))
	fmt.Printf("The size of cat1 color is %v.\n", unsafe.Sizeof(cat1.Color))
	fmt.Printf("The size of cat1 hobby is %v.\n", unsafe.Sizeof(cat1.Hobby))
}
