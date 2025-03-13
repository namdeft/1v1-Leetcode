func evalRPN(tokens []string) int {
	var result int
	stack := make([]int, 0, len(tokens))
	for _, val := range tokens {
		if !isOperator(val) {
			num, _ := strconv.Atoi(val)
			stack = append(stack, num)
		} else {
			output := 0
			num1 := stack[len(stack)-2]
			num2 := stack[len(stack)-1]

			switch val {
			case "+":
				output = num1 + num2
			case "-":
				output = num1 - num2
			case "*":
				output = num1 * num2
			case "/":
				output = num1 / num2
			}

			stack = stack[:len(stack)-2]
			stack = append(stack, output)
		}

		result = stack[0]
	}

	return result
}

func isOperator(operator string) bool {
	if operator == "+" || operator == "-" || operator == "*" || operator == "/" {
		return true
	}

	return false
}