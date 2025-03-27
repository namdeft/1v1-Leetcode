func longestPalindrome(s string) int {
	var result int
	mp := make(map[rune]int)
	for _, val := range s {
		mp[val]++
	}

	hasOdd := false
	for _, val := range mp {
		result += (val / 2) * 2

		if val%2 != 0 {
			hasOdd = true
		}
	}
	if hasOdd {
		result += 1
	}

	return result
}