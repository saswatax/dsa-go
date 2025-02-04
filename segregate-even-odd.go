package main

func segreGateEvenOdd(head *Node) *Node {
	dummy := Node{-1, head}

	last := &dummy

	for last.Next != nil {
		last = last.Next
	}

	curLast := last
	cur := &dummy
	var fEven, fOdd *Node

	for cur.Next != fOdd {
		if cur.Next.Val%2 != 0 {
			curLast.Next = cur.Next
			curLast = cur.Next

			cur.Next.Next, cur.Next = nil, cur.Next.Next

			if fOdd == nil {
				fOdd = curLast
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

	i := &dummy
	j := &dummy
	var firstOdd *Node

	for j != nil && j.Next != nil {
		if j.Next.Val%2 == 0 {
			i.Next = j.Next
			i = j.Next
			j.Next = j.Next.Next
			j = j.Next.Next
		} else {
			j = j.Next
			if firstOdd == nil {
				firstOdd = j
			}
		}
	}

	i.Next = firstOdd

	return dummy.Next
}
