package main

const n = 7

type Emp struct {
	Id   int
	Name string
	Next *Emp
}

type EmpLink struct {
	Head *Emp
}

func (e *EmpLink) Insert(emp *Emp) {

}

type HashTable struct {
	LinkArr [n]EmpLink
}

func (h *HashTable) Insert(emp *Emp) {
	linkNo := h.HashFunc(emp.Id)
	h.LinkArr[linkNo].Insert(emp)
}

func (h *HashTable) HashFunc(id int) (res int) {
	res = id % 7
	return
}

func main() {

}
