func maxProfit(prices []int) int {
	var result int
	if len(prices) <= 1 {
		return 0
	}

	left := 0
	right := 1
	maxProfit := 0
	for right < len(prices) {
		if prices[left] > prices[right] {
			left = right
		} else {
			currentProfit := prices[right] - prices[left]
			maxProfit = max(maxProfit, currentProfit)
		}

		right++
	}

	result = maxProfit
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}