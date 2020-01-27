package main

import "fmt"

type student struct {
	Name string
	Age  int
}

func test01() {
	m := make(map[string]*student)
	students := []student{
		{Name: "zebu", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	for _, s := range students {
		m[s.Name] = &s
	}
	for k, v := range m {
		fmt.Println(k, v.Name, v.Age)
	}
}

func swap(a, b *int) (*int, *int) {
	fmt.Printf("swap{a: %v; b: %v.}\n", a, b)
	a, b = b, a
	fmt.Printf("swap{a: %v; b: %v.}\n", a, b)
	return a, b
}

func test02() {
	a, b := 3, 4
	c, d := swap(&a, &b) // 3 4 4 3
	fmt.Printf("c: %v; d: %v.\n", c, d)
	fmt.Printf("*c: %v; *d: %v.\n", *c, *d)
	a = *c // 4 4 4 4
	b = *d // 4 4 4 4
	fmt.Printf("a: %v; b: %v.\n", a, b)
	fmt.Printf("&a: %v; &b: %v.\n", &a, &b)
}

func main() {
	switch 2 {
	case 1:
		test01()
	case 2:
		test02()
	}
}
