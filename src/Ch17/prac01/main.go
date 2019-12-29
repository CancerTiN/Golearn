package main

import (
	"fmt"
	"reflect"
)

type Monster struct {
	Name  string  `json:"name"`
	Age   int     `json:"age"`
	Score float32 `json:"score"`
	Sex   string
}

func (m Monster) Print() {
	fmt.Println("---start---")
	fmt.Println(m)
	fmt.Println("----end----")
}

func (m Monster) GetSum(n1, n2 int) int {
	return n1 + n2
}

func (m Monster) Set(name string, age int, score float32, sex string) {
	m.Name = name
	m.Age = age
	m.Score = score
	m.Sex = sex
}

func testStruct(i interface{}) {
	rValue := reflect.ValueOf(i)
	rValueKind := rValue.Kind()
	fmt.Printf("The kind of the incoming interface is %v.\n", rValueKind)
	rType := reflect.TypeOf(i)
	fmt.Printf("The type of the incoming interface is %v.\n", rType)
	if rValueKind != reflect.Struct {
		fmt.Println("The kind of the incoming interface is not a struct.")
		return
	}

	rValueNumField := rValue.NumField()
	fmt.Printf("This struct has %v fields.\n", rValueNumField)
	for fieldIndex := 0; fieldIndex < rValueNumField; fieldIndex++ {
		rTypeField := rType.Field(fieldIndex)
		fmt.Printf("The type of field %d is %v.\n", fieldIndex, rTypeField)
		fmt.Printf("The value of field %d is %v.\n", fieldIndex, rValue.Field(fieldIndex))
		jsonTag := rTypeField.Tag.Get("json")
		if jsonTag != "" {
			fmt.Printf("The json tag of field %d is %v.\n", fieldIndex, jsonTag)
		}
	}

	rValueNumMethod := rValue.NumMethod()
	fmt.Printf("This struct has %v methods.\n", rValueNumMethod)
	for methodIndex := 0; methodIndex < rValueNumMethod; methodIndex++ {
		fmt.Printf("The value of method %d is %v.\n", methodIndex, rValue.Method(methodIndex))
	}
	// Print
	rValue.Method(1).Call(nil)
	// GetSum
	var paramSlice []reflect.Value
	paramSlice = append(paramSlice, reflect.ValueOf(10))
	paramSlice = append(paramSlice, reflect.ValueOf(40))
	result := rValue.Method(0).Call(paramSlice)
	fmt.Printf("The result is %v.\n", result)
	fmt.Printf("The result[0] is %v.\n", result[0])
	fmt.Printf("The result[0].Int() is %v.\n", result[0].Int())

}

func main() {
	var m Monster = Monster{
		Name:  "Weasel",
		Age:   400,
		Score: 30.8,
	}
	testStruct(m)
}
