func romanToInt(s string) int {
	var result int
	mp := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	for i := 0; i < len(s); i++ {
		if i == len(s)-1 {
			result += mp[s[i]]
			continue
		}

		if mp[s[i]] < mp[s[i+1]] {
			result += mp[s[i+1]] - mp[s[i]]
			i++
		} else {
			result += mp[s[i]]
		}
	}

	return result
}