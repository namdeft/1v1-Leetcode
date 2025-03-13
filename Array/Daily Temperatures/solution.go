func dailyTemperatures(temperatures []int) []int {
	results := make([]int, len(temperatures))
	stack := make([]int, 0)
	for i, val := range temperatures {
		for len(stack) > 0 && temperatures[stack[len(stack)-1]] < val {
			index := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			results[index] = i - index
		}

		stack = append(stack, i)
	}

	return results
}