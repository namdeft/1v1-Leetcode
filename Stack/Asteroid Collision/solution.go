func asteroidCollision(asteroids []int) []int {
	var results []int
	for _, val := range asteroids {
		for {
			if len(results) == 0 || !willColide(results[len(results)-1], val) {
				results = append(results, val)
				break
			}

			last := len(results) - 1

			if abs(results[last]) == abs(val) {
				results = results[:last]
				break
			}

			if abs(results[last]) > abs(val) {
				break
			}

			if abs(results[last]) < abs(val) {
				results = results[:last]
			}
		}
	}

	return results
}

func willColide(a, b int) bool {
	if a > 0 && b < 0 {
		return true
	}

	return false
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}