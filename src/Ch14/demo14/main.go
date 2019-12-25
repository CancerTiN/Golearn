package main

import (
	"encoding/json"
	"fmt"
)

type Monster struct {
	Name     string `json:"monster_name"`
	Age      int    `json:"monster_age"`
	Birthday string
	Salary   float64
	Skill    string
}

func testStruct() {
	monster := Monster{
		Name:     "OxDevil",
		Age:      500,
		Birthday: "2011-11-11",
		Salary:   8000.0,
		Skill:    "BullMagicFist",
	}
	data, err := json.Marshal(&monster)
	if err != nil {
		fmt.Println("Serialization failed error:")
		fmt.Println(err)
	}
	fmt.Printf("The type of data is %T.\n", data)
	fmt.Println("data =", data)
	fmt.Println("data string =", string(data))
}

func testMap() {
	// var m map[string]interface{}
	m := make(map[string]interface{})
	m["name"] = "RedBabe"
	m["age"] = 30
	m["address"] = "HongYaDong"

	data, err := json.Marshal(m)
	if err != nil {
		fmt.Println("Serialization failed error:")
		fmt.Println(err)
	}
	fmt.Printf("The type of data is %T.\n", data)
	fmt.Println("data =", data)
	fmt.Println("data string =", string(data))
}

func testSlice() {
	var slice []map[string]interface{}

	m1 := make(map[string]interface{})
	m1["name"] = "jack"
	m1["age"] = "7"
	m1["address"] = "Beijing"
	slice = append(slice, m1)

	m2 := make(map[string]interface{})
	m2["name"] = "tom"
	m2["age"] = "20"
	m2["address"] = [2]string{"Mexico", "Hawaii"}
	slice = append(slice, m2)

	data, err := json.Marshal(slice)
	if err != nil {
		fmt.Println("Serialization failed error:")
		fmt.Println(err)
	}
	fmt.Printf("The type of data is %T.\n", data)
	fmt.Println("data =", data)
	fmt.Println("data string =", string(data))
}

func testFloat64() {
	var num float64 = 2345.67

	data, err := json.Marshal(num)
	if err != nil {
		fmt.Println("Serialization failed error:")
		fmt.Println(err)
	}
	fmt.Printf("The type of data is %T.\n", data)
	fmt.Println("data =", data)
	fmt.Println("data string =", string(data))
}

func main() {
	testStruct()
	testMap()
	testSlice()
	testFloat64()
}
