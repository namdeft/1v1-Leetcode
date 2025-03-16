func merge(intervals [][]int) [][]int {
	var results [][]int
	if len(intervals) == 1 {
		return intervals
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	results = append(results, intervals[0])

	j := 0
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= results[j][1] {
			results[j][0] = min(intervals[i][0], results[j][0])
			results[j][1] = max(intervals[i][1], results[j][1])
		} else {
			results = append(results, intervals[i])
			j++
		}
	}

	return results
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

// Input: intervals = [[1,3],[2,6],[8,10],[15,18]]
// Output: [[1,6],[8,10],[15,18]]
// Explanation: Since intervals [1,3] and [2,6] overlap, merge them into [1,6].

// Input: intervals = [[1,4],[4,5]]
// Output: [[1,5]]
// Explanation: Intervals [1,4] and [4,5] are considered overlapping.