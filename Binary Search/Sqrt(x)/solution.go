func mySqrt(x int) int {
	var result int
	left := 0
	right := x

	for left <= right {
		mid := (left + right) / 2

		if mid*mid == x {
			result = mid
			break
		}

		if mid*mid > x {
			right = mid - 1
		}

		if mid*mid < x {
			result = mid
			left = mid + 1
		}
	}

	return result
}