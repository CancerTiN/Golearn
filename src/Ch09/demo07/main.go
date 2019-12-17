package main

import "fmt"

type Stu struct {
	Name    string
	Age     int
	Address string
}

func modify(map1 map[int]int) {
	map1[10] = 900
}

func main() {
	map1 := make(map[int]int)
	map1[1] = 90
	map1[2] = 88
	map1[10] = 1
	map1[20] = 2
	fmt.Println(map1) // map[1:90 2:88 10:1 20:2]
	modify(map1)
	fmt.Println(map1) // map[1:90 2:88 10:900 20:2]

	students := make(map[string]Stu)
	stu1 := Stu{"tom", 18, "beijing"}
	stu2 := Stu{"mary", 18, "shanghai"}
	students["no1"] = stu1
	students["no2"] = stu2
	fmt.Println(students)

	for k, v := range students {
		fmt.Printf("student id is %v\n", k)
		fmt.Printf("\tstudent name is %v\n", v.Name)
		fmt.Printf("\tstudent age is %v\n", v.Age)
		fmt.Printf("\tstudent address is %v\n", v.Address)
		fmt.Println()
	}
}
