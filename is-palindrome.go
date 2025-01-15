package main

import "math"

func isPalindrome(num int) bool {
	digitCount := int(math.Floor(math.Log10(float64(num))) + 1)

	for range digitCount / 2 {
		curDigitCount := int(math.Floor(math.Log10(float64(num))) + 1)

		headDigit := num / int(math.Pow10(curDigitCount-1))
		tailDigit := num % 10

		num = num % int(math.Pow10(curDigitCount-1))
		num = num / 10

		if headDigit != tailDigit {
			return false
		}
	}

	return true
}
