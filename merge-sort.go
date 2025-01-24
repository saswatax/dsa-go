package main

func mergeSort(arr []int, left, right int) {
	if left >= right {
		return
	}

	mid := left + (right-left)/2
	mergeSort(arr, left, mid)
	mergeSort(arr, mid+1, right)
	merge(arr, left, mid, right)
}

func merge(arr []int, left, mid, right int) {
	len1, len2 := mid-left+1, right-mid
	arr1 := make([]int, len1)
	arr2 := make([]int, len2)

	for i := 0; i < len1; i++ {
		arr1[i] = arr[left+i]
	}

	for j := 0; j < len2; j++ {
		arr2[j] = arr[mid+1+j]
	}

	m, n, l := 0, 0, left
	for m < len1 && n < len2 {
		if arr1[m] > arr2[n] {
			arr[l] = arr2[n]
			n++
		} else {
			arr[l] = arr1[m]
			m++
		}
		l++
	}

	for m < len1 {
		arr[l] = arr1[m]
		m++
		l++
	}

	for n < len2 {
		arr[l] = arr2[n]
		n++
		l++
	}
}
