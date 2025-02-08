package main

func findKthLargest(nums []int, k int) int {
	l := len(nums) - k
	left, right := 0, len(nums)-1

	for left <= right {
		p := partition(nums, left, right)
		if p == l {
			return nums[p]
		} else if p > l {
			right = p - 1
		} else {
			left = p + 1
		}
	}

	return -1
}

func partition(nums []int, left, right int) int {
	piviot := nums[right]
	i := left

	for j := left; j <= right-1; j++ {
		if nums[j] < piviot {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}

	nums[i], nums[right] = nums[right], nums[i]

	return i
}
