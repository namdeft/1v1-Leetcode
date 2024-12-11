func canConstruct(ransomNote string, magazine string) bool {
	_map := make(map[rune]int)

	for _, v := range magazine {
		_map[v]++
	}

	for _, v := range ransomNote {
		_map[v]--

		if v, _ := _map[v]; v < 0 {
			return false
		}
	}

	return true
}

// aa -> aab