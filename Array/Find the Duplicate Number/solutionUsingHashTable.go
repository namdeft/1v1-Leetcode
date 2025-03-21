func findDuplicate(nums []int) int {
	var result int
	mp := make(map[int]bool)
	for _, val := range nums {
		if _, ok := mp[val]; ok {
			result = val
		}

		mp[val] = true
	}

	return result
}