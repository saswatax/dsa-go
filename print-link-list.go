package main

import "fmt"

func printLinkList(head *Node) {
	for cur := head; cur != nil; cur = cur.Next {
		fmt.Print(cur.Val, " ")
	}
	fmt.Println("")
}
