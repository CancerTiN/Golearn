package main

import "runtime"

func test1() {
	runtime.Gosched()
}

func main() {

}
