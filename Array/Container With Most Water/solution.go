func maxArea(height []int) int {
	var area int
	left, right := 0, len(height)-1

	for left < right {
		width := right - left
		if height[left] < height[right] {
			area = max(area, height[left]*width)
			left++
		} else {
			area = max(area, height[right]*width)
			right--
		}
	}

	return area
}

func max(a, b int) int {
	if a >= b {
		return a
	}

	return b
}