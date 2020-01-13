package main

import "fmt"

type HeroNode struct {
	no       int
	name     string
	nickname string
	next     *HeroNode
}

func InsertHeroNode(head, newHeroNode *HeroNode) {
	temp := head
	for {
		if temp.next == nil {
			break
		}
		temp = temp.next
	}
	temp.next = newHeroNode
}

func ListHeroNode(head *HeroNode) {
	temp := head
	if temp.next == nil {
		fmt.Println("The linked list is currently empty.")
	}
	for {
		if temp.next == nil {
			break
		}
		fmt.Printf("{%d, %s, %s}", temp.next.no, temp.next.name, temp.next.nickname)
		temp = temp.next
		if temp.next != nil {
			fmt.Print(" ==> ")
		}
	}
	fmt.Println()
}

func InsertSortedHeroNode(head, newHeroNode *HeroNode) {
	temp := head
	flag := true
	for {
		if temp.next == nil {
			break
		} else if temp.next.no > newHeroNode.no {
			break
		} else if temp.next.no == newHeroNode.no {
			flag = false
			break
		}
		temp = temp.next
	}
	if !flag {
		fmt.Printf("Fail to insert because {no: %d} already exists.\n", newHeroNode.no)
		return
	} else {
		newHeroNode.next = temp.next
		temp.next = newHeroNode
	}
}

func DeleteHeroNode(head *HeroNode, id int) {
	temp := head
	flag := false
	for {
		if temp.next == nil {
			break
		} else if temp.next.no == id {
			flag = true
			break
		}
		temp = temp.next
	}
	if flag {
		temp.next = temp.next.next
	} else {
		fmt.Printf("Fail to delete because {no: %d} does not exists.\n", id)
	}
}

func main() {
	head := &HeroNode{}

	hero1 := &HeroNode{
		no:       1,
		name:     "Song Jiang",
		nickname: "Rain in time",
		next:     nil,
	}

	hero2 := &HeroNode{
		no:       2,
		name:     "Lu Junyi",
		nickname: "Jade Unicorn",
		next:     nil,
	}

	hero3 := &HeroNode{
		no:       3,
		name:     "Lin Chong",
		nickname: "Leopard head",
		next:     nil,
	}

	hero4 := &HeroNode{
		no:       4,
		name:     "Wu Yong",
		nickname: "Chitastar",
		next:     nil,
	}

	key := 1
	switch key {
	case 0:
		InsertHeroNode(head, hero1)
		InsertHeroNode(head, hero2)
		InsertHeroNode(head, hero3)
	case 1:
		InsertSortedHeroNode(head, hero3)
		InsertSortedHeroNode(head, hero1)
		InsertSortedHeroNode(head, hero2)
		InsertSortedHeroNode(head, hero4)
	}

	ListHeroNode(head)
	DeleteHeroNode(head, 2)
	ListHeroNode(head)
}
