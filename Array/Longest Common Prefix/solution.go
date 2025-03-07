func longestCommonPrefix(strs []string) string {
	var result string
	sort.Slice(strs, func(i, j int) bool {
		return strs[i] < strs[j]
	})

	first := strs[0]
	last := strs[len(strs)-1]
	for i := 0; i < len(first); i++ {
		if first[i] != last[i] {
			break
		} else {
			result += string(first[i])
		}
	}

	return result
}