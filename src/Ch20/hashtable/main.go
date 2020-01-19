package main

import (
	"fmt"
)

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

func (e *EmpLink) Get(id int) (name string, err error) {
	if e.Head == nil {
		err = fmt.Errorf("id error <%d>", id)
		return
	}
	var ok bool
	for temp := e.Head.Next; temp != nil; temp = temp.Next {
		if id == temp.Id {
			name = temp.Name
			ok = true
			break
		}
	}
	if !ok {
		err = fmt.Errorf("id error <%d>", id)
	}
	return
}

func (e *EmpLink) Delete(id int) (err error) {
	if e.Head == nil {
		err = fmt.Errorf("id error <%d>", id)
		return
	}
	var ok bool
	for temp := e.Head.Next; temp != nil; temp = temp.Next {
		if id == temp.Id {
			temp.Last.Next = temp.Next
			if temp.Next != nil {
				temp.Next.Last = temp.Last
			}
			ok = true
			break
		}
	}
	if e.Head.Next == nil {
		e.Head = nil
	}
	if !ok {
		err = fmt.Errorf("id error <%d>", id)
	}
	return
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

func (h *HashTable) Get(id int) (name string, err error) {
	linkNo := h.HashFunc(id)
	if linkNo < 0 || linkNo >= nLink {
		err = fmt.Errorf("hash result <%d> is out of bound [0, %d)", linkNo, nLink)
	} else {
		name, err = h.LinkArr[linkNo].Get(id)
	}
	return
}

func (h *HashTable) Delete(id int) (err error) {
	linkNo := h.HashFunc(id)
	if linkNo < 0 || linkNo >= nLink {
		err = fmt.Errorf("hash result <%d> is out of bound [0, %d)", linkNo, nLink)
	} else {
		err = h.LinkArr[linkNo].Delete(id)
	}
	return
}

func (h *HashTable) HashFunc(id int) (res int) {
	res = id % 7
	return
}

func main() {
	emp0 := &Emp{Id: 0, Name: "Kareem Abdul-Jabbar"}
	emp1 := &Emp{Id: 1, Name: "Nate Archibald"}
	emp2 := &Emp{Id: 2, Name: "Paul Arizin"}
	emp3 := &Emp{Id: 3, Name: "Charles Barkley"}
	emp4 := &Emp{Id: 4, Name: "Rick Barry"}
	emp5 := &Emp{Id: 5, Name: "Elgin Baylor"}
	emp6 := &Emp{Id: 6, Name: "Dave Bing"}
	emp7 := &Emp{Id: 7, Name: "LARRY BIRD"}
	hashTable := &HashTable{}
	hashTable.Insert(emp0)
	hashTable.Insert(emp1)
	hashTable.Insert(emp2)
	hashTable.Insert(emp3)
	hashTable.Insert(emp4)
	hashTable.Insert(emp5)
	hashTable.Insert(emp6)
	hashTable.Insert(emp7)
	hashTable.Show()
	for _, id := range []int{-1, 3, 8} {
		name, err := hashTable.Get(id)
		if err != nil {
			fmt.Printf("hashTable.Get(%d) error: %v.\n", id, err)
		} else {
			fmt.Printf("hashTable.Get(%d) return: %s.\n", id, name)
		}
	}
	id := 4
	err := hashTable.Delete(4)
	if err != nil {
		fmt.Printf("hashTable.Delete(%d) error: %v.\n", id, err)
	} else {
		fmt.Printf("hashTable.Delete(%d) successfully.\n", id)
	}
	hashTable.Show()
}
