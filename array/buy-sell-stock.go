package array

func BuyAndSellStock(prices []int) int {
	result, curResult, left := 0, 0, 0

	for right := 1; right < len(prices); right++ {
		if prices[right] > prices[right-1] {
			curResult = prices[right] - prices[left]
		} else {
			result += curResult
			left = right
			curResult = 0
		}
	}

	result += curResult

	return result
}
