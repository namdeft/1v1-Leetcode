func majorityElement(nums []int) int {
	mp := make(map[int]int)
	n := len(nums) / 2
	var result int

	for _, num := range nums {
		mp[num]++
	}

	for k, v := range mp {
		if v > n {
			result = k
		}
	}

	return result
}