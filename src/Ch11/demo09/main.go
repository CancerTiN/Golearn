package main

import "fmt"

type Usb interface {
	Start()
	Stop()
}

type Phone struct {
	name string
}

func (p Phone) Start() {
	fmt.Println("The phone started to work.")
}

func (p Phone) Stop() {
	fmt.Println("The phone stopped working.")
}

type Camera struct {
	name string
}

func (c Camera) Start() {
	fmt.Println("The camera started to work.")
}

func (c Camera) Stop() {
	fmt.Println("The camera stopped working.")
}

func main() {
	var usbArr [3]Usb
	usbArr[0] = Phone{"vivo"}
	usbArr[1] = Phone{"xiaomi"}
	usbArr[2] = Camera{"nikon"}
	fmt.Println(usbArr)
}
