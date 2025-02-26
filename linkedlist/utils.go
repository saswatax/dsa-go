package linkedlist

import "fmt"

type Node[V any] struct {
	Val  V
	Next *Node[V]
}

func CreateLinkedList[V any](arr []V) *Node[V] {
	var zero V
	dummy := Node[V]{zero, nil}
	last := &dummy

	for _, v := range arr {
		node := Node[V]{v, nil}
		last.Next = &node
		last = &node
	}

	return dummy.Next
}

func PrintLinkedList[V any](head *Node[V]) {
	for cur := head; cur != nil; cur = cur.Next {
		fmt.Print(cur.Val, " ")
	}

	fmt.Println("")
}
