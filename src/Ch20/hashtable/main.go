package main

import "fmt"

const nLink = 7

type Emp struct {
	Id   int
	Name string
	Last *Emp
	Next *Emp
}

func (e *Emp) Show() {
	fmt.Printf("{Id: %d, Name: %s}", e.Id, e.Name)
}

type EmpLink struct {
	Head *Emp
}

func (e *EmpLink) Insert(emp *Emp) {
	if e.Head == nil {
		e.Head = &Emp{}
	}
	temp := e.Head
	for temp.Next != nil {
		temp = temp.Next
		if (e.Head == temp.Last || temp.Last.Id < emp.Id) && emp.Id < temp.Id {
			emp.Last = temp.Last
			emp.Last.Next = emp
			temp.Last = emp
			emp.Next = temp
			return
		}
	}
	temp.Next = emp
	emp.Last = temp
}

func (e *EmpLink) Show() {
	if e.Head == nil {
		fmt.Println("The link is empty.")
		return
	}
	for temp := e.Head.Next; temp != nil; temp = temp.Next {
		temp.Show()
		if temp.Next != nil {
			fmt.Print(" <-> ")
		}
	}
	fmt.Println()
}

type HashTable struct {
	LinkArr [nLink]EmpLink
}

func (h *HashTable) Insert(emp *Emp) {
	linkNo := h.HashFunc(emp.Id)
	h.LinkArr[linkNo].Insert(emp)
}

func (h *HashTable) Show() {
	for i, empLink := range h.LinkArr {
		fmt.Printf("The situation of LinkArr[%d] is:\n", i)
		empLink.Show()
	}
}

func (h *HashTable) HashFunc(id int) (res int) {
	res = id % 7
	return
}

func main() {
	emp1 := &Emp{Id: 0, Name: "Kareem Abdul-Jabbar"}
	emp7 := &Emp{Id: 7, Name: "LARRY BIRD"}
	hashTable := &HashTable{}
	hashTable.Insert(emp1)
	hashTable.Insert(emp7)
	hashTable.Show()
}
