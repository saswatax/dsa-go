package stack

import "fmt"

func StockSpan(arr []int) []int {
	stack := ArrayStack{[]int{}, -1}

	for _, v := range arr {
		t, _ := stack.Peak()
		fmt.Println(t, stack.Size())
		temp, err := stack.Pop()
		fmt.Println(temp, err)
		stack.Push(v)
	}

	return stack.arr
}
