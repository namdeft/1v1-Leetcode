func wordPattern(pattern string, s string) bool {
	arr := strings.Split(s, " ")
	if len(arr) != len(pattern) {
		return false
	}
	mpPattern := make(map[byte]string)
	mpString := make(map[string]byte)

	for i := 0; i < len(pattern); i++ {
		if val, ok := mpPattern[pattern[i]]; ok {
			if val != arr[i] {
				return false
			}
		}
		mpPattern[pattern[i]] = arr[i]

		if val, ok := mpString[arr[i]]; ok {
			if val != pattern[i] {
				return false
			}
		}
		mpString[arr[i]] = pattern[i]
	}

	return true
}