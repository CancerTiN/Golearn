package main

import "fmt"

func main() {

	var heroes [3]string = [3]string{"songjiang", "wuyong", "lujunyi"}
	fmt.Println(heroes)

	for index, value := range heroes {
		fmt.Printf("heroes[%v] = %v\n", index, value)
		fmt.Printf("heroes[%v] = %v\n", index, heroes[index])
	}

	for _, v := range heroes {
		fmt.Printf("The value of the element is %v\n", v)
	}
}
