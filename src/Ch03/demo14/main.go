package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "hello, world"
	fmt.Println(len(s))
	fmt.Println(s[0], s[7])

	c := s[len(s)-1]
	fmt.Println(c)

	fmt.Println(s[0:5])
	fmt.Println(s[:5])
	fmt.Println(s[7:])
	fmt.Println(s[:])
	fmt.Printf("pointer of s is %p\n", &s)

	t := s[:]
	fmt.Printf("pointer of s is %p\n", &t)

	fmt.Println("goodbye" + s[5:])

	a := "left foot"
	b := a
	a += ", right foot"
	fmt.Println(a)
	fmt.Println(b)

	fmt.Println("\xe4\xb8\x96")
	fmt.Println("\u4e16")
	fmt.Println("\U00004e16")

	d := "Hello, \xe4\xb8\x96\xe7\x95\x8c"
	fmt.Println(len(d))                    // 13
	fmt.Println(utf8.RuneCountInString(d)) // 9
	for i := 0; i < len(d); {
		r, size := utf8.DecodeRuneInString(d[i:])
		fmt.Printf("%d\t%c\n", i, r)
		i += size
	}

	for i, r := range "Hello, \xe4\xb8\x96\xe7\x95\x8c" {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}

	e := "\xe3\x83\x97\xe3\x83\xad\xe3\x82\xb0\xe3\x83\xa9\xe3\x83\xa0"
	fmt.Println(e)
	f := []rune(e)
	fmt.Println(f)
	fmt.Printf("%x\n", f)
	fmt.Println(string(f))

	fmt.Println(string(65))            // "A"
	fmt.Println(fmt.Sprintf("%v", 65)) // "65"
}
