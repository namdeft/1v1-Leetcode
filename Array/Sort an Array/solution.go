func sortArray(nums []int) []int {
	var results []int
	if len(nums) <= 1 {
		return nums
	}

	mid := len(nums) / 2
	leftArr := sortArray(nums[:mid])
	rightArr := sortArray(nums[mid:])

	results = merge(leftArr, rightArr)

	return results
}

func merge(leftArr []int, rightArr []int) []int {
	var results []int
	l, r := 0, 0

	for l < len(leftArr) && r < len(rightArr) {
		if leftArr[l] < rightArr[r] {
			results = append(results, leftArr[l])
			l++
		} else {
			results = append(results, rightArr[r])
			r++
		}
	}

	results = append(results, leftArr[l:]...)
	results = append(results, rightArr[r:]...)

	return results
}