func longestConsecutive(nums []int) int {
	var result int
	mp := make(map[int]bool)
	for _, val := range nums {
		mp[val] = true
	}

	for val, _ := range mp {
		if !mp[val-1] {
			len := 1
			currentNum := val

			for mp[currentNum+1] {
				len++
				currentNum++
			}

			if len >= result {
				result = len
			}
		}
	}

	return result
}