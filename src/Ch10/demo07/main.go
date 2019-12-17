package main

import (
	"encoding/json"
	"fmt"
)

type Monster struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Skill string `json:"skill"`
}

func main() {
	monster := Monster{"OxDevil", 500, "BananaFan"}

	jsonStr, err := json.Marshal(monster)
	if err != nil {
		fmt.Println("json processing error:", err)
	}
	fmt.Println("json bytes:", jsonStr)
	fmt.Println("json strings:", string(jsonStr))
}
