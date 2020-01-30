package main

import (
	"fmt"
	"strconv"
)

func main() {
	intStr := "10000"
	rInt64, err := strconv.ParseInt(intStr, 0, 0)
	fmt.Printf("strconv.ParseInt return: (%v, %v).\n", rInt64, err)
	fmt.Printf("%#v, %T\n", rInt64, rInt64)

	rInt, err := strconv.Atoi(intStr)
	fmt.Printf("strconv.Atoi return: (%v, %v).\n", rInt, err)
	fmt.Printf("%#v, %T\n", rInt, rInt)

	rInt32 := int32(97)
	rStr := fmt.Sprintf("%d", rInt32)
	fmt.Printf("%#v, %T\n", rStr, rStr)

	rStr = strconv.Itoa(int(rInt32))
	fmt.Printf("%#v, %T\n", rStr, rStr)

	boolStr := "true"
	rBool, err := strconv.ParseBool(boolStr)
	fmt.Printf("strconv.ParseBool return: (%v, %v).\n", rBool, err)
	fmt.Printf("%#v, %T\n", rBool, rBool)

	floatStr := "1.234"
	rFloat64, err := strconv.ParseFloat(floatStr, 0)
	fmt.Printf("strconv.ParseFloat return: (%v, %v).\n", rFloat64, err)
	fmt.Printf("%#v, %T\n", rFloat64, rFloat64)
}
