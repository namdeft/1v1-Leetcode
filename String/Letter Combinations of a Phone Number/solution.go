func letterCombinations(digits string) []string {
	var results []string
	mp := map[byte][]byte{
		'2': []byte{'a', 'b', 'c'},
		'3': []byte{'d', 'e', 'f'},
		'4': []byte{'g', 'h', 'i'},
		'5': []byte{'j', 'k', 'l'},
		'6': []byte{'m', 'n', 'o'},
		'7': []byte{'p', 'q', 'r', 's'},
		'8': []byte{'t', 'u', 'v'},
		'9': []byte{'w', 'x', 'y', 'z'},
	}
	for len(digits) > 0 {
		letterArr := mp[digits[len(digits)-1]]
		digits = digits[:len(digits)-1]

		if len(results) > 0 {
			var tempArr []string
			for i := 0; i < len(letterArr); i++ {
				for j := 0; j < len(results); j++ {
					tempArr = append(tempArr, string(letterArr[i])+results[j])
				}
			}
			results = tempArr

		} else {
			for _, char := range letterArr {
				results = append(results, string(char))
			}
		}
	}

	return results
}