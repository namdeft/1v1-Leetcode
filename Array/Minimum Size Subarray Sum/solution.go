func minSubArrayLen(target int, nums []int) int {
	var result int
	result = math.MaxInt64
	left := 0
	sum := 0
	for right := 0; right < len(nums); right++ {
		sum += nums[right]

		for sum >= target {
			sum -= nums[left]
			result = min(result, right-left+1)
			left++
		}

	}

	if result == math.MaxInt64 {
		return 0
	}

	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}