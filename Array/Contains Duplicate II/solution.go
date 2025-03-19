func containsNearbyDuplicate(nums []int, k int) bool {
	mp := make(map[int]int)
	l := 0
	for r := 0; r < len(nums); r++ {
		mp[nums[r]]++

		if mp[nums[r]] > 1 {
			return true
		}

		if r-l == k {
			mp[nums[l]]--
			l++
		}
	}

	return false
}