func findPoisonedDuration(timeSeries []int, duration int) int {
	var result int
	if len(timeSeries) == 1 {
		return duration
	}

	for i := 1; i < len(timeSeries); i++ {
		timeDiff := timeSeries[i] - timeSeries[i-1]

		if timeDiff < duration {
			result += timeDiff
		} else {
			result += duration
		}
	}
	result += duration

	return result
}