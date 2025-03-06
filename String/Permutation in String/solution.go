func checkInclusion(s1 string, s2 string) bool {
	mpS1Count := make(map[rune]int)
	for i, _ := range s1 {
		mpS1Count[s1[i]]++
	}
	mpS2Count := make(map[rune]int)
	left := 0
	matches := 0

	for right := 0; right < len(s2); right++ {
		char := s2[right]
		mpS2Count[char]++
		if mpS1Count[char] == mpS2Count[char] {
			matches++
		}

		if right-left == len(s1) {
			char = s2[left]
			if mpS1Count[char] == mpS2Count[char] {
				matches--
			}
			mpS2Count[char]--

			left++
		}

		if matches == len(mpS1Count) {
			return true
		}
	}

	return false
}