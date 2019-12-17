package main

import (
	"fmt"
	"sort"
)

func main() {
	map1 := make(map[int]int)
	map1[10] = 100
	map1[1] = 13
	map1[4] = 56
	map1[8] = 90
	fmt.Println(map1)

	// sort key
	var keys []int
	for k, _ := range map1 {
		keys = append(keys, k)
	}
	fmt.Println("Keys before sorting is", keys)
	sort.Ints(keys)
	fmt.Println("Keys after sorting is", keys)

	for _, k := range keys {
		fmt.Printf("map[%v] = %v\n", k, map1[k])
	}
}
