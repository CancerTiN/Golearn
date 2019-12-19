package main

import "fmt"

func main() {
	slice := []int{0, 1, 2, 3}
	m := make(map[int]*int)

	for key, val := range slice {
		m[key] = &val
		fmt.Printf("key: %v; address: %p; content: %v; value: %v\n", key, &val, val, *(&val))
	}

	for k, v := range m {
		fmt.Println(k, "->", *v)
	}
}
