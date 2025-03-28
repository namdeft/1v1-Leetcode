func findLengthOfLCIS(nums []int) int {
	var result int
	if len(nums) <= 1 {
		return 1
	}

	currentLongest := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			currentLongest++
		} else {
			currentLongest = 1
		}

		result = max(result, currentLongest)
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}