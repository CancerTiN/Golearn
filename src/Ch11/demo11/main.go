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
	fmt.Println("Phone started working ...")
}

func (p Phone) Stop() {
	fmt.Println("Phone stopped working ...")
}

func (p Phone) Call() {
	fmt.Println("Phone call ...")
}

type Camera struct {
	name string
}

func (c Camera) Start() {
	fmt.Println("Camera started working ...")
}

func (c Camera) Stop() {
	fmt.Println("Camera stopped working ...")
}

type Computer struct {
}

func (c Computer) Working(usb Usb) {
	usb.Start()
	if phone, ok := usb.(Phone); ok {
		phone.Call()
	}
	usb.Stop()
}

func main() {
	var usbArr [3]Usb

	usbArr[0] = Phone{"vivo"}
	usbArr[1] = Phone{"mi"}
	usbArr[2] = Camera{"nikon"}

	var computer Computer
	for _, v := range usbArr {
		computer.Working(v)
		fmt.Println()
	}
}
