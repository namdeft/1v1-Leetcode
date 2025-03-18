func calPoints(operations []string) int {
	var result int
	stack := make([]int, 0, len(operations))
	for _, val := range operations {
		if isNumber(val) {
			num, _ := strconv.Atoi(val)
			stack = append(stack, num)
		}
		if len(stack) > 0 && val == "C" {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 && val == "D" {
			stack = append(stack, stack[len(stack)-1]*2)
		}
		if len(stack) >= 2 && val == "+" {
			stack = append(stack, stack[len(stack)-1]+stack[len(stack)-2])
		}
	}

	for _, val := range stack {
		result += val
	}

	return result
}

func isNumber(char string) bool {
	if char == "C" || char == "D" || char == "+" {
		return false
	}

	return true
}