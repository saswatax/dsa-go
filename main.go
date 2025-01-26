package main

import (
	"fmt"
)

func main() {
	arr := []int{3, 2, 1, 5, 6, 4}
	heapSort(arr)
	fmt.Println("Result: ", arr)
}

// func main() {
// 	head := ListNode{-1, nil}
// 	last := &head

// 	for _, v := range []int{1, 2, 3, 4, 5, 6} {
// 		node := ListNode{v, nil}
// 		last.Next = &node
// 		last = &node
// 	}

// 	result := swapPairs(head.Next)
// 	fmt.Println("Result:", result)

// 	for cur := result; cur != nil; cur = cur.Next {
// 		fmt.Println(cur.Val)
// 	}
// }

// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }
