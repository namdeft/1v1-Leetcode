func twoSum(nums []int, target int) []int {
	var results []int

	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return append(results, i, j)
			}
		}
	}

	return results
}