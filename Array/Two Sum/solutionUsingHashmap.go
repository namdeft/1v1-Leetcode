package main

func twoSum(nums []int, target int) []int {
	var result []int
	hashMap := make(map[int]int)

	for i, num := range nums {
		found := target - num
		if val, ok := hashMap[found]; ok {
			result = append(result, i, val)
		}
		hashMap[num] = i
	}

	return result
}
