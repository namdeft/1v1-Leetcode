func maxProfit(prices []int) int {
	var result int
	if len(prices) <= 1 {
		return 0
	}

	for i := 1; i < len(prices); i++ {
		currentProfit := prices[i] - prices[i-1]
		if currentProfit > 0 {
			result += currentProfit
		}
	}

	return result
}