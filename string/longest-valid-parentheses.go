package string

import "math"

func LongestValidParentheses(s string) int {
	result := 0
	curResult := 0
	counter := 0
	streak := true

	for _, v := range s {
		if v == rune('(') {
			counter++
		} else {
			counter--
		}

		if counter == 0 {
			if streak {
				curResult += 2
			} else {
				curResult = 2
			}
		}
		if counter < 0 {
			result = int(math.Max(float64(result), float64(curResult)))
			streak = false
		}

	}

	result = int(math.Max(float64(result), float64(curResult)))

	return result
}
