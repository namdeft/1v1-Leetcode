func isHappy(n int) bool {
	mp := make(map[int]bool)

	for !mp[n] {
		digit := 0
		sum := 0
		mp[n] = true

		for n > 0 {
			digit = n % 10
			sum += digit * digit
			n = n / 10
		}

		if sum == 1 {
			return true
		}
		n = sum
	}

	return false
}