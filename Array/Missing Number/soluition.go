func missingNumber(nums []int) int {
	var result int
	var calculatedSum int
	var expectedSum int

	for i := 0; i <= len(nums); i++ {
		expectedSum += i
	}

	for _, val := range nums {
		calculatedSum += val
	}

	result = expectedSum - calculatedSum
	return result
}