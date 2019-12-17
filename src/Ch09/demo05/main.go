package main

import "fmt"

func main() {
	var monsters []map[string]string
	monsters = make([]map[string]string, 2)
	fmt.Printf("The address of monsters is %p.\n", &monsters)
	test(monsters)
	exec(&monsters)
	// The following writing can dynamically add content to the map.
	monster := map[string]string{
		"name": "FireCloudEvilGod",
		"age":  "200",
	}
	monsters = append(monsters, monster)
	fmt.Printf("The address of monsters is %p.\n", &monsters)
	fmt.Println("monsters =", monsters)
}

func test(monsters []map[string]string) {
	fmt.Printf("The address of monsters is %p.\n", &monsters)
	fmt.Println("monsters =", monsters)
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("Found the following error in the test function:")
			fmt.Println(err)
		}
	}()
	if monsters[0] == nil {
		monsters[0] = make(map[string]string, 2)
		monsters[0]["name"] = "OxDevil"
		monsters[0]["age"] = "500"
	}
	if monsters[1] == nil {
		monsters[1] = make(map[string]string, 2)
		monsters[1]["name"] = "YuTuJing"
		monsters[1]["age"] = "400"
	}

	// The wording below will lead to cross-border.
	if monsters[2] == nil {
		monsters[2] = make(map[string]string, 2)
		monsters[2]["name"] = "Vixen"
		monsters[2]["age"] = "300"
	}
	fmt.Println("It will not print here.")
	fmt.Println("monsters =", monsters)
}

func exec(monsters *[]map[string]string) {
	fmt.Printf("The address of monsters is %p.\n", monsters)
	fmt.Println("monsters =", *monsters)
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("Found the following error in the exec function:")
			fmt.Println(err)
		}
	}()
	if (*monsters)[0] == nil {
		(*monsters)[0] = make(map[string]string, 2)
		(*monsters)[0]["name"] = "OxDevil"
		(*monsters)[0]["age"] = "500"
	}
	if (*monsters)[1] == nil {
		(*monsters)[1] = make(map[string]string, 2)
		(*monsters)[1]["name"] = "YuTuJing"
		(*monsters)[1]["age"] = "400"
	}

	// The wording below will lead to cross-border.
	if (*monsters)[2] == nil {
		(*monsters)[2] = make(map[string]string, 2)
		(*monsters)[2]["name"] = "Vixen"
		(*monsters)[2]["age"] = "300"
	}
	fmt.Println("It will not print here.")
	fmt.Println("monsters =", *monsters)
}
