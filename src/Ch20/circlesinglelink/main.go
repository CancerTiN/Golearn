package main

import "fmt"

type CatNode struct {
	no   int
	name string
	next *CatNode
}

func InsertCatNode(head, newCatNode *CatNode) {
	if head.next == nil {
		head.no = newCatNode.no
		head.name = newCatNode.name
		head.next = head
		fmt.Println(newCatNode, "has been added to CircleSingleLink.")
		return
	}
	temp := head
	for {
		if temp.next == head {
			break
		}
		temp = temp.next
	}
	temp.next = newCatNode
	newCatNode.next = head
}

func ListCircleLink(head *CatNode) {
	temp := head
	if temp.next == nil {
		fmt.Println("The circle link is currently empty.")
		return
	}
	fmt.Println("The situation of the circular link is as follows:")
	for {
		fmt.Println(temp)
		if temp.next == head {
			break
		}
		temp = temp.next
	}
}

func DeleteCatNode(head *CatNode, id int) {

}

func main() {
	head := &CatNode{}

	cat1 := &CatNode{no: 1, name: "tom1"}
	cat2 := &CatNode{no: 2, name: "tom2"}
	cat3 := &CatNode{no: 3, name: "tom3"}

	InsertCatNode(head, cat1)
	InsertCatNode(head, cat2)
	InsertCatNode(head, cat3)

	fmt.Printf("%v, %p\n", head, head)
	fmt.Printf("%v, %p\n", cat1, cat1)
	fmt.Printf("%v, %p\n", cat2, cat2)
	fmt.Printf("%v, %p\n", cat3, cat3)
	ListCircleLink(head)
}
