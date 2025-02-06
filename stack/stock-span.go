package stack

func StockSpan(arr []int) []int {
	if len(arr) == 0 {
		return nil
	}

	stack := ArrayStack[int]{[]int{0}}
	result := make([]int, 0, len(arr))
	result = append(result, 1)

	for i := 1; i < len(arr); i++ {
		for {
			val, err := stack.Peak()
			if err == nil && arr[val] <= arr[i] {
				stack.Pop()
			} else {
				break
			}
		}

		val, err := stack.Peak()
		if err == nil {
			result = append(result, i-val)
		} else {
			result = append(result, i+1)
		}

		stack.Push(i)
	}

	return result
}
