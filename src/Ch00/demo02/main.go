package main

import "fmt"

func main() {
	var v1 int     // 声明
	v1 = 10        // 赋值
	var v2 float64 // 声明不赋值
	var v3 = 10.11 // 类型推导
	v4 := "major"  // 短变量声明
	// 多变量声明方式1
	var (
		v5 = 100
		v6 = "200"
	)
	// 多变量声明方式2
	v7, v8 := "300", 400
	fmt.Println(v1, v2, v3, v4, v5, v6, v7, v8)
}
