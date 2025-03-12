func maximumCount(nums []int) int {
	countNeg := binarySearch(nums, 0)
	countPos := len(nums) - binarySearch(nums, 1)

	return max(countNeg, countPos)
}

func binarySearch(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	result := len(nums)

	for left <= right {
		mid := (left + right) / 2
		if nums[mid] < target {
			left = mid + 1
		} else {
			result = mid
			right = mid - 1
		}
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}