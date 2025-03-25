func maxSubArray(nums []int) int {
	var result int
	currentSum := nums[0]
	result = nums[0]
	for i := 1; i < len(nums); i++ {
		currentSum = max(nums[i], currentSum+nums[i])
		result = max(result, currentSum)
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}