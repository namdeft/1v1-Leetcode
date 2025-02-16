func isPalindrome(s string) bool {
	if len(s) == 0 {
		return true
	}

	pointer1 := 0
	pointer2 := len(s) - 1
	for pointer1 < pointer2 {
		for pointer1 < pointer2 && !isAlphabet(s[pointer1]) {
			pointer1++
		}
		for pointer1 < pointer2 && !isAlphabet(s[pointer2]) {
			pointer2--
		}
		if unicode.ToLower(rune(s[pointer1])) != unicode.ToLower(rune(s[pointer2])) {
			return false
		}

		pointer1++
		pointer2--
	}

	return true
}

func isAlphabet(char byte) bool {
	if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9') {
		return true
	}

	return false
}