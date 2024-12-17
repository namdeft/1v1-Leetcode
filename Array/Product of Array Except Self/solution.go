func productExceptSelf(nums []int) []int {
	result := make([]int, len(nums))
	product := 1
	var numberOfZero int

	for _, v := range nums {
		if v == 0 {
			numberOfZero++
			continue
		}

		product *= v
	}

	for i, v := range nums {
		if numberOfZero == 0 {
			result[i] = product / v
		} else if numberOfZero == 1 {
			result[i] = product
		}
	}

	return result
}