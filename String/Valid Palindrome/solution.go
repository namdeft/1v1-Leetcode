func isValid(s string) bool {
	stack := make([]rune, 0)
	mp := map[rune]rune{
		'}': '{',
		')': '(',
		']': '[',
	}

	for _, char := range s {
		if char == '(' || char == '{' || char == '[' {
			stack = append(stack, char)
		}

		if val, ok := mp[char]; ok {
			if len(stack) > 0 && stack[len(stack)-1] == val {
				stack = stack[0 : len(stack)-1]
			} else {
				return false
			}
		}
	}

	if len(stack) == 0 {
		return true
	}
	return false
}