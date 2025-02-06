package stack

import "fmt"

func StockSpan(arr []int) []int {
	var stack ArrayStack[int] = []int{1, 2, 3}

	for _, v := range arr {
		t, _ := stack.Peak()
		fmt.Println(t, stack.Size())
		temp, err := stack.Pop()
		fmt.Println(temp, err)
		stack.Push(v)
	}

	return stack
}
