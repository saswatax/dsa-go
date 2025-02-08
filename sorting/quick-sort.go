package sorting

func QuickSort(arr []int) {
	quickSort(arr, 0, len(arr)-1)
}

func quickSort(arr []int, left, right int) {
	if left >= right {
		return
	}

	p := lomutoPartition(arr, left, right)
	quickSort(arr, left, p-1)
	quickSort(arr, p+1, right)
}

func hoarePartition(arr []int, left, right int) int {
	pivot := arr[left]
	i, j := left, right

	for {
		for arr[i] < pivot {
			i++
		}
		for arr[j] > pivot {
			j--
		}

		if i >= j {
			break
		}

		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}

	return j
}

func lomutoPartition(arr []int, left, right int) int {
	pivot := arr[right]
	i := left

	for j := left; j < right; j++ {
		if arr[j] <= pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}

	arr[i], arr[right] = arr[right], arr[i]

	return i
}
