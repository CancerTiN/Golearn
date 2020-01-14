package main

import (
	"fmt"
)

type HeroNode struct {
	no       int
	name     string
	nickname string
	pre      *HeroNode
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
	newHeroNode.pre = temp
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
			flag = true
			break
		}
		temp = temp.next
	}
	if !flag {
		fmt.Printf("Fail to insert because {no: %d} already exists.\n", newHeroNode.no)
		return
	} else {
		newHeroNode.next = temp.next
		if newHeroNode.next != nil {
			newHeroNode.next.pre = newHeroNode
		}
		temp.next = newHeroNode
		newHeroNode.pre = temp
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
		if temp.next != nil {
			temp.next.pre = temp
		}
	} else {
		fmt.Printf("Fail to delete because {no: %d} does not exists.\n", id)
	}
}

func ListHeroNode(head *HeroNode) {
	temp := head
	if temp.next == nil {
		fmt.Println("The double link is currently empty.")
		return
	}
	for {
		if temp.next == nil {
			break
		}
		fmt.Printf("{%d, %s, %s}", temp.next.no, temp.next.name, temp.next.nickname)
		temp = temp.next
		if temp.next != nil {
			fmt.Print(" <=> ")
		} else {
			fmt.Println()
		}
	}
}

func ReverseListHeroNode(head *HeroNode) {
	temp := head
	if temp.next == nil {
		fmt.Println("The linked list is currently empty.")
		return
	}
	for {
		if temp.next == nil {
			break
		}
		temp = temp.next
	}
	for {
		fmt.Printf("{%d, %s, %s}", temp.no, temp.name, temp.nickname)
		temp = temp.pre
		if temp.pre == nil {
			fmt.Println()
			break
		} else {
			fmt.Print(" <=> ")
		}
	}
}

func main() {
	head := &HeroNode{}

	hero1 := &HeroNode{
		no:       1,
		name:     "Song Jiang",
		nickname: "Rain in time",
	}

	hero2 := &HeroNode{
		no:       2,
		name:     "Lu Junyi",
		nickname: "Jade Unicorn",
	}

	hero3 := &HeroNode{
		no:       3,
		name:     "Lin Chong",
		nickname: "Leopard head",
	}

	hero4 := &HeroNode{
		no:       4,
		name:     "Wu Yong",
		nickname: "Chitastar",
	}

	key := 1
	switch key {
	case 0:
		InsertHeroNode(head, hero1)
		InsertHeroNode(head, hero2)
		InsertHeroNode(head, hero3)
		InsertHeroNode(head, hero4)
	case 1:
		InsertSortedHeroNode(head, hero3)
		InsertSortedHeroNode(head, hero1)
		InsertSortedHeroNode(head, hero2)
		InsertSortedHeroNode(head, hero4)
	}

	ListHeroNode(head)
	DeleteHeroNode(head, 4)
	ReverseListHeroNode(head)
}
