package main

import "fmt"

func main() {
	s := make([]string, 10)
	fmt.Printf("type of s is %T\n", s)

	//m := make(map[string]string, 10)
	m := make(map[string]string)
	fmt.Printf("type of s is %T\n", m)
	m["no1"] = "songjiang"
	m["no2"] = "wuyong"
	m["no1"] = "wusong"
	m["no3"] = "wuyong"
	fmt.Println(m)
}
