package main

// func main() {
// 	result := kmp("ababcab", "ab")
// 	fmt.Println("Result: ", result)
// }

func main() {
	head := Node{-1, nil}
	last := &head

	for _, v := range []int{1, 2, 4, 3, 5, 6, 7} {
		node := Node{v, nil}
		last.Next = &node
		last = &node
	}

	result := segreGateEvenOdd(head.Next)
	printLinkList(result)
}

type Node struct {
	Val  int
	Next *Node
}
