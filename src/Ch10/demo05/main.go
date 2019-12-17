package main

import "fmt"

type Point struct {
	x int
	y int
}

//type Rect struct {
//	leftUp Point
//	rightDown Point
//}

type Rect struct {
	leftUp, rightDown Point
}

type Rect2 struct {
	leftUp, rightDown *Point
}

func main() {
	r1 := Rect{Point{1, 2}, Point{3, 4}}
	fmt.Printf("The address of r1.leftUp.x is %p\n", &r1.leftUp.x)
	fmt.Printf("The address of r1.leftUp.y is %p\n", &r1.leftUp.y)
	fmt.Printf("The address of r1.rightDown.x is %p\n", &r1.rightDown.x)
	fmt.Printf("The address of r1.rightDown.y is %p\n", &r1.rightDown.y)

	r2 := Rect2{&Point{10, 20}, &Point{30, 40}}
	fmt.Printf("The address of r2.leftUp is %p\n", &r2.leftUp)
	fmt.Printf("The address of r2.rightDown is %p\n", &r2.rightDown)
	fmt.Printf("The content of r2.leftUp is %p\n", r2.leftUp)
	fmt.Printf("The content of r2.rightDown is %p\n", r2.rightDown)
}
