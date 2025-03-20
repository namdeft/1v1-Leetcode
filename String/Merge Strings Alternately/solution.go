func mergeAlternately(word1 string, word2 string) string {
	var result string

	j := 0
	for i := 0; i < len(word1); i++ {
		result += string(word1[i])

		if j < len(word2) {
			result += string(word2[j])
			j++
		}
	}

	var lenDiff int
	if len(word2) > len(word1) {
		lenDiff = len(word2) - len(word1)
	}

	for i := len(word2) - lenDiff; i < len(word2); i++ {
		result += string(word2[i])
	}

	return result
}