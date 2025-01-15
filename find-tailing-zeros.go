package main

import (
	"fmt"
	"math"
)

func findTailingZeros(num int) int {
	iteration := int(math.Log(float64(num)) / math.Log(5))
	fmt.Println(iteration)

	result := 0
	for i := range iteration {
		result += (num / int(math.Pow(5, float64(i+1))))
	}

	return result
}
