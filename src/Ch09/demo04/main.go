package main

import "fmt"

func main() {
	cities := make(map[string]string)
	cities["no1"] = "beijing"
	cities["no2"] = "tianjin"
	cities["no3"] = "shanghai"
	fmt.Println("cities =", cities)
	for k, v := range cities {
		fmt.Printf("k = %v; v = %v.\n", k, v)
	}

	fmt.Printf("There are %v key-value pairs in cities.\n", len(cities))

	students := make(map[string]map[string]string)
	students["stu01"] = make(map[string]string)
	students["stu01"]["name"] = "tom"
	students["stu01"]["sex"] = "male"
	students["stu01"]["address"] = "beijing"
	students["stu02"] = make(map[string]string)
	students["stu02"]["name"] = "mary"
	students["stu02"]["sex"] = "female"
	students["stu02"]["address"] = "shanghai"
	fmt.Println("students =", students)
	for k1, v1 := range students {
		fmt.Println("k1 =", k1)
		for k2, v2 := range v1 {
			fmt.Printf("\t k2 = %v; v2 = %v.\n", k2, v2)
		}
	}
}
