package main

import (
	"fmt"

	"github.com/saswatax/dsa-go/linkedlist"
)

func main() {
	head := linkedlist.CreateLinkedList([]int{1, 2})

	fmt.Print("")

	result := linkedlist.SegreGateEvenOddSingleTraversal(head)
	linkedlist.PrintLinkedList(result)
}
