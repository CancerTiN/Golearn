package main

import "fmt"

func main() {
	// type 1
	var men map[string]string
	men = make(map[string]string, 2)
	men["no1"] = "songjiang"
	men["no2"] = "wuyong"
	men["no1"] = "wusong"
	men["no3"] = "wuyong"
	fmt.Println("men =", men)

	// type 2
	cities := make(map[string]string)
	cities["no1"] = "beijing"
	cities["no2"] = "tianjin"
	cities["no3"] = "shanghai"
	fmt.Println("cities =", cities)

	// type 3
	heroes := map[string]string{
		"hero1": "songjiang",
		"hero2": "lujunyi",
		"hero3": "wuyong",
	}
	heroes["hero4"] = "linchong"
	fmt.Println("heroes =", heroes)

	// example
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
	fmt.Println("student2 =", students["stu02"])
}
