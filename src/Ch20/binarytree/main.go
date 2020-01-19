package main

import "fmt"

type Node struct {
	Key   int
	Value string
	Left  *Node
	Right *Node
}

func (n *Node) Show() {
	fmt.Printf("{Key: %d, Value: %s}\n", n.Key, n.Value)
}

func PreOrder(node *Node) {
	if node != nil {
		node.Show()
		PreOrder(node.Left)
		PreOrder(node.Right)
	}
}

func InfixOrder(node *Node) {
	if node != nil {
		InfixOrder(node.Left)
		node.Show()
		InfixOrder(node.Right)
	}
}

func PostOrder(node *Node) {
	if node != nil {
		PostOrder(node.Left)
		PostOrder(node.Right)
		node.Show()
	}
}

func main() {
	root := &Node{Key: 0, Value: "A"}

	left1 := &Node{Key: 1, Value: "B"}
	right1 := &Node{Key: 2, Value: "C"}

	root.Left = left1
	root.Right = right1

	node10 := &Node{Key: 10, Value: "z"}
	node11 := &Node{Key: 11, Value: "y"}

	left1.Left = node10
	left1.Right = node11

	right2 := &Node{Key: 3, Value: "D"}

	right1.Right = right2

	fmt.Println("Preorder traversal:")
	PreOrder(root)
	fmt.Println("In-order traversal:")
	InfixOrder(root)
	fmt.Println("Postorder traversal:")
	PostOrder(root)
}
