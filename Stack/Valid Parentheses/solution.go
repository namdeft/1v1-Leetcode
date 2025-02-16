func isValid(s string) bool {
	stack := make([]rune, 0)
	mp := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, char := range s {
		if val, ok := mp[char]; ok {
			if len(stack) == 0 || val != stack[len(stack)-1] {
				return false
			} else {
				return true
			}
		} else {
			stack = append(stack, char)
		}
	}

	if len(stack) == 0 {
		return true
	} else {
		return false
	}
}