package main

import "fmt"

type Monkey struct {
	Name string
}

func (m *Monkey) climbing() {
	fmt.Println(m.Name, "is born to climb trees.")
}

type LittleMonkey struct {
	Monkey
}

func (l *LittleMonkey) Flying() {
	fmt.Println(l.Name, "will fly by learning.")
}

func (l *LittleMonkey) Swimming() {
	fmt.Println(l.Name, "can swim by learning.")
}

func main() {
	monkey := LittleMonkey{Monkey{"Goku"}}
	monkey.climbing()
	monkey.Flying()
	monkey.Swimming()
}
