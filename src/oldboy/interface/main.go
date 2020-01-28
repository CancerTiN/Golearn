package main

import "fmt"

type TransInfo struct{}

type Fragment interface {
	Exec(transInfo *TransInfo) error
}

type GetPodAction struct{}

func (g GetPodAction) Exec(transInfo *TransInfo) error {
	return nil
}

func test01() {
	var f1 Fragment = new(GetPodAction)
	//var f2 Fragment = GetPodAction
	var f3 Fragment = &GetPodAction{}
	var f4 Fragment = GetPodAction{}
	fmt.Println(f1, f3, f4)
}

type A interface {
	method1()
}

type B interface {
	method1()
	method2()
}

type implementB struct{}

func (i implementB) method1() {}

func (i implementB) method2() {}

func test02() {
	var a A = implementB{}
	a.method1()
}

func main() {
	switch 2 {
	case 1:
		test01()
	case 2:
		test02()
	default:
		document01()
		document02()
		document03()
	}
}

func document01() {
	// 关于接口和结构体的正确说法是：
	// 1)一个结构体只需要实现接口要求的所有方法，就可以认为这个结构体实现了该接口
	// 2)实现“类”时只需要关心应该实现哪些方法即可
	// 3)结构体实现接口时，不需要导入接口所在的包
	// 4)使用方按自身需求来定义接口，且无需关心接口的具体实现细节
}

func document02() {
	/*
		关于接口，下面说法正确：
		1)只要两个接口拥有相同的方法列表，那么它们就是等价的，可以相互赋值
		2)如果接口A的方法列表是接口B的方法列表的子集，那么接口B的具体实现可以赋值给A
		3)接口可以继承
	*/
}

func document03() {
	// interface不一定是一个值类型
	// interface{}可以指向任意对象的类型
}
