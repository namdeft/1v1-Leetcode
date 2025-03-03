func topKFrequent(nums []int, k int) []int {
	var result []int
	mp := make(map[int]int)
	for _, val := range nums {
		mp[nums]++
	}

	bucket := make([][]int, len(nums))

	for k, val := range mp {
		bucket[val] = append(bucket[val], k)
	}

	for i := len(bucket) - 1; i > 0; i-- {
		for _, val := range bucket[i] {
			if k > 0 {
				result = append(result, val)
				k--
			}
		}
	}

	return result
}
