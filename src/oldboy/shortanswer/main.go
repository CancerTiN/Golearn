package main

import "fmt"

func main() {
	question01()
	question02()
	question03()
	question04()
	question05()
}

func question01() {
	/*
		cap()和len()函数的区别是什么？
		cap()可以用来查看指定类型的容量
		len()可以用来查看指定类型的长度
	*/
}

func question02() {
	/*
		什么是byte？什么是rune？如何将[]byte和[]rune类型的值转换为字符串？
		byte是uint8的别名，在所有方面都等同于uint8。按照惯例，它用于区分字节值和8位无符号整数值。
		rune是int32的别名，在所有方面都等同于int32。按照惯例，它用于区分字符值和整数值。
	*/
	str1 := string([]byte{31, 127, 255})
	str2 := string([]rune{256, 512, 1024})
	fmt.Println(str1, str2)
}

func question03() {
	/*
		简述go语言中make和new的区别。
		1)内建函数new用来分配内存，它的第一个参数是一个类型，不是一个值，它的返回值是一个指向新分配类型零值的指针。
		2)内建函数make用来为slice，map或chan类型分配内存和初始化一个对象，跟new类似，第一个参数也是一个类型而不是一个值，
		  跟new不同的是，make返回类型的引用而不是指针，而返回值也依赖于具体传入的类型。
		3)new的作用是初始化一个指向类型的指针(*T)，make的作用是为slice，map或chan初始化并返回引用(T)。
	*/
}

func question04() {
	/*
		简述闭包的生命周期和作用范围。
		闭包是函数和它所引用的环境，也就是闭包=函数+引用环境。
	*/
}

func question05() {
	/*
		测试文件必须以什么结尾？功能测试函数必须以什么开头？压力测试函数必须以什么开头？
		测试文件必须以_test.go结尾。
		功能测试函数必须以Test开头。
		压力测试函数必须以Benchmark开头。
	*/
}
