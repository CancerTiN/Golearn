package main

import (
	"fmt"
	"math/rand"
	"sort"
)

type Hero struct {
	Name string
	Age  int
}

type HeroSlice []Hero

func (hs HeroSlice) Len() int {
	return len(hs)
}

func (hs HeroSlice) Less(i, j int) bool {
	// return hs[i].Age < hs[j].Age // Sort Ascending
	return hs[i].Name < hs[j].Name // Sort Ascending
}

func (hs HeroSlice) Swap(i, j int) {
	hs[i], hs[j] = hs[j], hs[i]
}

type Student struct {
	Name  string
	Age   int
	Score float64
}

func main() {

	var intSlice = []int{0, -1, 10, 7, 90}
	sort.Ints(intSlice)

	fmt.Println(intSlice)

	var heroes HeroSlice
	for i := 0; i < 10; i++ {
		hero := Hero{
			Name: fmt.Sprintf("hero~%d", rand.Intn(100)),
			Age:  rand.Intn(100),
		}
		heroes = append(heroes, hero)
	}
	for _, v := range heroes {
		fmt.Println(v)
	}

	fmt.Println("--------  Sort  --------")

	sort.Sort(heroes)
	for _, v := range heroes {
		fmt.Println(v)
	}
}
