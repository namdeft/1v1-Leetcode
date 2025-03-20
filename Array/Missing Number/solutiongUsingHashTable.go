func missingNumber(nums []int) int {
	var result int
	mp := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		mp[nums[i]]++
	}

	for j := 0; j <= len(nums); j++ {
		if _, ok := mp[j]; !ok {
			result = j
		}
	}

	return result
}