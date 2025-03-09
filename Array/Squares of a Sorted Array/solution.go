func sortedSquares(nums []int) []int {
	var result []int
	left := 0
	right := len(nums) - 1

	for left <= right {
		if squareNum(nums[left]) > squareNum(nums[right]) {
			result = append(result, squareNum(nums[left]))
			left++
		} else {
			result = append(result, squareNum(nums[right]))
			right--
		}
	}
	reverseSlice(result)

	return result
}

func squareNum(num int) int {
	return int(math.Pow(float64(num), 2))
}

func reverseSlice(slice []int) {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}