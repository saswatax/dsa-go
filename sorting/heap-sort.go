package sorting

func HeapSort(arr []int) {
	n := len(arr)

	for i := n/2 - 1; i >= 0; i-- {
		maxHeapify(arr, n, i)
	}

	for i := n - 1; i > 0; i-- {
		arr[i], arr[0] = arr[0], arr[i]
		maxHeapify(arr, i, 0)
	}
}

func maxHeapify(arr []int, n, i int) {
	max := i
	left := i*2 + 1
	right := i*2 + 2

	if left < n && arr[left] > arr[max] {
		max = left
	}

	if right < n && arr[right] > arr[max] {
		max = right
	}

	if max != i {
		arr[max], arr[i] = arr[i], arr[max]
		maxHeapify(arr, n, max)
	}
}
