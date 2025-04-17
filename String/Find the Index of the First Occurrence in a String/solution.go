func strStr(haystack string, needle string) int {
	var result int
	result = -1
	var str string
	l := 0
	for r := len(needle) - 1; r < len(haystack); r++ {
		str = haystack[l : r+1]
		if needle == str {
			return l
		}

		l++
	}

	return result
}