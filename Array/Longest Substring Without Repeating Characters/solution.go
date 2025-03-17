func lengthOfLongestSubstring(s string) int {
	var result int
	mp := make(map[byte]bool)

	l := 0
	r := 0
	count = 0
	for r < len(s) {
		if _, ok := mp[s[r]]; !ok {
			mp[s[r]] = true
			count = r - l + 1
			result = max(result, count)
			r++
		} else {
			delete(mp, s[l])
			l++
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