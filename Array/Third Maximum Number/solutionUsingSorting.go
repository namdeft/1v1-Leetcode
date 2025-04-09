func thirdMax(nums []int) int {
	var result int
	if len(nums) <= 1 {
		return nums[0]
	}

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	result = nums[len(nums)-1]
	time := 2
	for i := len(nums) - 2; i >= 0; i-- {
		if time == 0 {
			break
		}

		if nums[i] < result {
			result = nums[i]
			time--
		}
	}

	if time > 0 {
		return nums[len(nums)-1]
	}

	return result
}