func twoSum(nums []int, target int) []int {
	numsLength := len(nums)

	for i := 0; i < numsLength-1; i++ {
		for j := i + 1; j < numsLength; j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return []int{}
}.