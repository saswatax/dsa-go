package linkedlist

func Reverse(head *Node) *Node {
	var pre *Node = nil
	cur := head

	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}

	return pre
}
