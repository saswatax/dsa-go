package linkedlist

func Reverse[V any](head *Node[V]) *Node[V] {
	var pre *Node[V]
	cur := head

	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}

	return pre
}
