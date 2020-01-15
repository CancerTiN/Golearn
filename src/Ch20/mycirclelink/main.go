package main

import (
	"errors"
	"fmt"
)

type Node struct {
	Id   int
	Next *Node
}

func InsertCircleLink(head, newNode *Node) {
	if head.Next == nil {
		head.Next = newNode
		newNode.Next = newNode
		fmt.Println("The first node was inserted.")
		return
	}
	temp := head.Next
	for n := 1; n < 10; n++ {
		if temp.Next == head.Next {
			fmt.Printf("Insert the node at position <%d>.\n", n)
			newNode.Next = temp.Next
			temp.Next = newNode
			break
		}
		temp = temp.Next
	}
}

func ListCircleLink(head *Node) {
	if head.Next == nil {
		fmt.Println("CircleLink is empty.")
		return
	}
	temp := head
	for n := 0; n < 10; n++ {
		if n != 0 && temp.Next == head.Next {
			break
		}
		temp = temp.Next
		fmt.Printf("%d [%p] => {Id: %d, Next: ([%p] => %v)}\n", n, temp, temp.Id, temp.Next, temp.Next)
	}
}

func GetHeadNextId(head *Node) (id int, err error) {
	if head.Next == nil {
		err = errors.New("head node points to <nil>")
		return
	}
	id = head.Next.Id
	return
}

func main() {
	head := &Node{}

	nod1 := &Node{Id: 25}
	nod2 := &Node{Id: 37}
	nod3 := &Node{Id: 49}
	nod4 := &Node{Id: 61}

	InsertCircleLink(head, nod1)

	ListCircleLink(head)

	InsertCircleLink(head, nod2)
	InsertCircleLink(head, nod3)
	InsertCircleLink(head, nod4)

	ListCircleLink(head)

	id, err := GetHeadNextId(head)
	if err != nil {
		fmt.Println("GetHeadNextId error:", err)
	} else {
		fmt.Printf("Head node points to node of which id is %d.\n", id)
	}
}
