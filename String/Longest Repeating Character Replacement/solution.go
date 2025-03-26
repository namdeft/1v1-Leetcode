func characterReplacement(s string, k int) int {
	var result int
	mp := make(map[byte]int)
	left := 0
	maxCharFrequent := 0
	for right := 0; right < len(s); right++ {
		mp[s[right]]++
		maxCharFrequent = max(maxCharFrequent, mp[s[right]])

		if (right-left+1)-maxCharFrequent > k {
			mp[s[left]]--
			left++
		}

		result = max(result, right-left+1)
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}