package main

type Node struct {
	Id   int
	Next *Node
}

func InsertCircleLink(head, newNode *Node) {
	if head.Next == nil {
		head.Next = newNode
		newNode.Next = newNode
		return
	}
	temp := head.Next
	for {
		if temp.Next == head.Next {
			newNode.Next = temp.Next
			temp.Next = newNode
		}
	}
}

func main() {

}
