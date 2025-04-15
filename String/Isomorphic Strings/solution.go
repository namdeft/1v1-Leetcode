func isIsomorphic(s string, t string) bool {
	mpS := make(map[byte]byte)
	mpT := make(map[byte]byte)

	for i := 0; i < len(s); i++ {
		if val, ok := mpS[s[i]]; ok {
			if val != t[i] {
				return false
			}
		}
		mpS[s[i]] = t[i]

		if val, ok := mpT[t[i]]; ok {
			if val != s[i] {
				return false
			}
		}
		mpT[t[i]] = s[i]
	}

	return true
}