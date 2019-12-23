package main

import "fmt"

type Usb interface {
	Start()
	Stop()
}

type Phone struct {
}

func (p Phone) Start() {
	fmt.Println("Phone started working ...")
}

func (p Phone) Stop() {
	fmt.Println("Phone turned off ...")
}

type Camera struct {
}

func (c Camera) Start() {
	fmt.Println("Camera started working ...")
}

func (c Camera) Stop() {
	fmt.Println("Camera turned off ...")
}

type Computer struct {
}

func (c Computer) Working(usb Usb) {
	usb.Start()
	usb.Stop()
}

func main() {
	computer := Computer{}
	phone := Phone{}
	camera := Camera{}

	computer.Working(phone)
	computer.Working(camera)
}
