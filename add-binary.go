package main

func addBinary(a string, b string) string {
	result := ""
	addValue := '0'

	for len(a) != 0 || len(b) != 0 {
		lastA, lastB := '0', '0'

		if len(a) != 0 {
			lastA = rune(a[len(a)-1])
		}

		if len(b) != 0 {
			lastB = rune(b[len(b)-1])
		}

		curResult1, curAddValue1 := addition(lastA, lastB)
		curResult2, _ := addition(curResult1, addValue)

		result = string(curResult2) + result
		addValue = curAddValue1

		if len(a) != 0 {
			a = a[:len(a)-1]
		}

		if len(b) != 0 {
			b = b[:len(b)-1]
		}
	}

	return result
}

func addition(i, j rune) (result, addValue rune) {
	if i == '1' && j == '1' {
		return '0', '1'
	}

	if i == '1' || j == '1' {
		return '1', '0'
	}

	return '0', '0'
}
