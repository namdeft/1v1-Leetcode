func twoSum(nums []int, target int) []int {
	var results []int
	numsMap := make(map[int]int)

	for index, value := range nums {
		numNeeds := target - value

		if pairNumNeeds, ok := numsMap[numNeeds]; ok {
			return append(results, index, pairNumNeeds)
		}

		numsMap[value] = index
	}

	return results
}