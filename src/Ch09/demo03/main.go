package main

import "fmt"

func main() {
	cities := make(map[string]string)
	cities["no1"] = "beijing"
	cities["no2"] = "tianjin"
	cities["no3"] = "shanghai"
	fmt.Println("cities =", cities)

	// modify
	cities["no3"] = "shanghai~"
	fmt.Println("cities =", cities)

	// delete
	delete(cities, "no1")
	delete(cities, "no4")
	fmt.Println("cities =", cities)

	// find
	for _, key := range [2]string{"no1", "no2"} {
		val, ok := cities[key]
		if ok {
			fmt.Printf("Find the key-value pair with the key %v and the value %v\n", key, val)
		} else {
			fmt.Printf("No value found for key %v\n", key)
		}
	}

	// delete all keys
	cities = make(map[string]string)
	fmt.Println("cities =", cities)
}
