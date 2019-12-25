package main

import (
	"encoding/json"
	"fmt"
)

type Monster struct {
	Name     string
	Age      int
	Birthday string
	Salary   float64
	Skill    string
}

func marshalStruct() string {
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
	return string(data)
}

func unmarshalStruct(s string) {
	var monster Monster
	err := json.Unmarshal([]byte(s), &monster)
	if err != nil {
		fmt.Println("Deserialization failed error:")
		fmt.Println(err)
	}
	fmt.Println("monster =", monster)
}

func unmarshalMap(s string) {
	var m map[string]interface{}
	err := json.Unmarshal([]byte(s), &m)
	if err != nil {
		fmt.Println("Deserialization failed error:")
		fmt.Println(err)
	}
	fmt.Println("m =", m)
}

func marshalSlice() string {
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
	return string(data)
}

func unmarshalSlice(s string) {
	var slice []interface{}
	err := json.Unmarshal([]byte(s), &slice)
	if err != nil {
		fmt.Println("Deserialization failed error:")
		fmt.Println(err)
	}
	fmt.Println("slice =", slice)
}

func main() {
	s := marshalStruct()
	unmarshalStruct(s)
	unmarshalMap(s)
	s = marshalSlice()
	unmarshalSlice(s)
}
