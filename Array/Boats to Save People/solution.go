func numRescueBoats(people []int, limit int) int {
	var result int
	sort.Slice(people, func(i, j int) bool {
		return people[i] < people[j]
	})
	left := 0
	right := len(people) - 1

	for left < right {
		if people[left]+people[right] <= limit {
			left++
		}
		result++
		right--
	}

	if left == right {
		result++
	}

	return result
}