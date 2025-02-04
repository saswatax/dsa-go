package main

import "math"

func trap(height []int) int {
	n := len(height)
	left, right := make([]int, n), make([]int, n)

	left[0] = height[0]
	for i := 1; i < n; i++ {
		left[i] = int(math.Max(float64(height[i]), float64(left[i-1])))
	}

	right[n-1] = height[n-1]
	for i := n - 2; i >= 0; i-- {
		right[i] = int(math.Max(float64(height[i]), float64(right[i+1])))
	}

	result := 0

	for i := 0; i < n; i++ {
		result += int(math.Min(float64(left[i]), float64(right[i]))) - height[i]
	}

	return result
}
