package string

func KMP(text, pattern string) []int {
	n, m := len(text), len(pattern)
	res := []int{}

	lps := findLps(pattern, m)
	i, j := 0, 0

	for i < n {
		if text[i] == text[j] {
			i++
			j++

			if j == m {
				res = append(res, i-m)
				j = lps[j-1]
			}
		} else {
			if j == 0 {
				i++
			} else {
				j = lps[j-1]
			}
		}
	}

	return res
}

func findLps(pattern string, n int) []int {
	lps := make([]int, n)
	len := 0
	i := 1

	for i < n {
		if pattern[len] == pattern[i] {
			len++
			lps[i] = len
			i++
		} else {
			if len == 0 {
				lps[i] = 0
				i++
			} else {
				len = lps[len-1]
			}
		}
	}

	return lps
}
