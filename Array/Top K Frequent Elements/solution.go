func topKFrequent(nums []int, k int) []int {
	var result []int
	var mp map[int]int
	for _, num := range nums {
		mp[num]++
	}

	var tempArr []Pair
	for k, v := range mp {
		tempArr = append(tempArr, Pair{num: k, count: v})
	}
	sort.Slice(tempArr, func(i, j int) bool {
		return tempArr[i].Count > tempArr[j].Count
	})

	for i := 0; i < k; i++ {
		result = append(result, tempArr[i].Num)
	}

	return result
}

type Pair struct {
	Num   int
	Count int
}