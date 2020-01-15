package main

import (
	"errors"
	"fmt"
)

var nMax = 1000

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
	for n := 1; n < nMax; n++ {
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
	for n := 0; n < nMax; n++ {
		if n != 0 && temp.Next == head.Next {
			break
		}
		temp = temp.Next
		fmt.Printf("%d [%p] => {Id: %d, Next: ([%p] => %v)}\n", n, temp, temp.Id, temp.Next, temp.Next)
	}
}

func DeleteCircleLink(head *Node, id int) (err error) {
	if head.Next == nil {
		err = errors.New("head node points to <nil>")
		return
	}
	temp := head
	for n := 0; n < nMax; n++ {
		if n != 0 && temp.Next == head.Next {
			err = fmt.Errorf("node of which id is %d was not found", id)
			break
		}
		if n == 0 && temp.Next.Id == id {
			if temp.Next.Next == temp.Next {
				temp.Next = nil
				break
			} else {
				temp.Next = temp.Next.Next
			}
		}
		if temp.Next.Id == id {
			temp.Next = temp.Next.Next
			break
		}
		temp = temp.Next
	}
	return
}

func main() {
	fmt.Println("Joseph problem will be solved here!")
	Play(5, 2, 3)
	//Play(4, 3, 2)
}

func Play(n, k, m int) {
	var slice []int

	head := &Node{}

	for id := 1; id <= n; id++ {
		InsertCircleLink(head, &Node{Id: id})
	}

	ListCircleLink(head)

	temp := head.Next
	for {
		if temp.Id == k {
			break
		}
		temp = temp.Next
	}

	for head.Next != nil {
		for step := 0; step < m-1; step++ {
			temp = temp.Next
		}
		outId := temp.Id
		slice = append(slice, outId)
		temp = temp.Next
		err := DeleteCircleLink(head, outId)
		if err != nil {
			fmt.Println("DeleteCircleLink error:", err)
		} else {
			fmt.Printf("Succeed in deleting node of which id is %d.\n", outId)
		}
		if temp == nil {
			temp = head.Next
		}
		ListCircleLink(head)
	}

	fmt.Println("slice:", slice)
}
