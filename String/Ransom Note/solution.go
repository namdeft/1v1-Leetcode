func canConstruct(ransomNote string, magazine string) bool {
	slice := make([]int, 26)

	for _, v := range magazine {
		slice[v-'a']++
	}

	for _, v := range ransomNote {
		slice[v-'a']--

		if slice[v-'a'] < 0 {
			return false
		}
	}

	return true
}