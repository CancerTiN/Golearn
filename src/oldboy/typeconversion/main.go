package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func test01() {
	type MyInt int
	var i int = 1
	var j MyInt = MyInt(i)
	fmt.Println(reflect.TypeOf(j))
}

func test02() {
	// 指针可以转化为有效的JSON文本
	i := 1
	var ii *int = &i
	di, err := json.Marshal(ii)
	fmt.Println(err)
	var io *int
	err = json.Unmarshal(di, &io)
	fmt.Println(err)
	fmt.Println(ii, di, *io)
}

func test03() {
	// channel不可以转化为有效的JSON文本
	var chi chan int
	dch, err := json.Marshal(chi)
	fmt.Println(err)
	var cho chan int
	err = json.Unmarshal(dch, &cho)
	fmt.Println(err)
	fmt.Println(chi, dch, cho)
}

func test04() {
	// func不可以转化为有效的JSON文本
	fi := func() {}
	fmt.Println(fi)
	df, err := json.Marshal(fi)
	fmt.Println(err)
	var fo func()
	err = json.Unmarshal(df, &fo)
	fmt.Println(err)
	fmt.Println(fi, df, fo)
}

func main() {
	switch 4 {
	case 1:
		test01()
	case 2:
		test02()
	case 3:
		test03()
	case 4:
		test04()
	}
}
