package main

func segreGateEvenOdd(head *Node) *Node {
	dummy := Node{-1, head}

	last := &dummy

	for last.Next != nil {
		last = last.Next
	}

	cur := &dummy
	var fEven, fOdd *Node

	for cur.Next != fOdd {
		if cur.Next.Val%2 != 0 {
			last.Next = cur.Next
			last = cur.Next

			cur.Next.Next, cur.Next = nil, cur.Next.Next

			if fOdd == nil {
				fOdd = last
			}
		} else {
			cur = cur.Next

			if fEven == nil {
				fEven = cur
			}
		}
	}

	if fEven == nil {
		return cur.Next
	}

	return fEven
}

func segreGateEvenOddSingleTraverSal(head *Node) *Node {
	dummy := Node{-1, head}

	var firstEven, firstOdd *Node
	lastEven, lastOdd := &dummy, &dummy

	cur := &dummy

	for cur.Next != nil {
		if cur.Next.Val%2 == 0 {
			lastEven.Next = cur.Next
			lastEven = cur.Next

			if firstEven == nil {
				firstEven = cur.Next
			}
		} else {
			lastOdd.Next = cur.Next
			lastOdd = cur.Next

			if firstOdd == nil {
				firstOdd = cur.Next
			}
		}

		cur.Next, cur.Next.Next = cur.Next.Next, nil
	}

	lastEven.Next = firstOdd

	return firstEven
}
