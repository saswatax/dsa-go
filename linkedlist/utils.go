package linkedlist

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

func CreateLinkedList(arr []int) *Node {
	head := Node{-1, nil}
	last := &head

	for _, v := range arr {
		node := Node{v, nil}
		last.Next = &node
		last = &node
	}

	return head.Next
}

func PrintLinkedList(head *Node) {
	for cur := head; cur != nil; cur = cur.Next {
		fmt.Print(cur.Val, " ")
	}

	fmt.Println("")
}
