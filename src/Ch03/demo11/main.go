package main

import "fmt"

func main() {
	var address string = "北京长城 110 hello world"
	fmt.Println("address =", address)

	var str1 = "hello"
	// str[0] = 'a' // Can not assign to str[0]
	fmt.Println("str1 =", str1)

	str3 := `反引号`
	fmt.Println("str3 =", str3)

	var str4 = "hello " + "world"
	str4 += " haha!"
	fmt.Println("str4 =", str4)

	var str5 = "z" + "y" + "x" +
		"w" + "v" + "w" +
		"t" + "s" + "r" +
		"q" + "p" + "o"
	fmt.Println("str5 =", str5)
}
