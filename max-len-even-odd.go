package main

import (
	"math"
)

func maxLenEvenOdd(arr []int) int {
	res, curRes := 1, 1

	for i := 1; i < len(arr); i++ {
		firstIsEven := (arr[i-1] % 2) == 0
		secondIsEven := (arr[i] % 2) == 0

		if firstIsEven != secondIsEven {
			curRes++
			res = int(math.Max(float64(res), float64(curRes)))
		} else {
			curRes = 1
		}
	}

	return res
}
