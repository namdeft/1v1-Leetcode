func isAnagram(s string, t string) bool {
	charMaps := make(map[rune]int)

	for _, v := range s {
		charMaps[v]++
	}

	for _, v := range t {
		charMaps[v]--
	}

	for _, v := range charMaps {
		if v != 0 {
			return false
		}
	}

	return true
}