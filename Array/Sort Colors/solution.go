func sortColors(nums []int) {
	quickSort(nums, 0, len(nums)-1)
}

func quickSort(nums []int, i, j int) {
	if i >= j {
		return
	}

	index := partition(nums, i, j)
	quickSort(nums, i, index-1)
	quickSort(nums, index+1, j)
}

func partition(nums []int, i, j int) int {
	var index int
	index = i
	pivot := nums[j]
	for left := i; left < j; left++ {
		if nums[left] < pivot {
			nums[index], nums[left] = nums[left], nums[index]
			index++
		}
	}
	nums[index], nums[j] = nums[j], nums[index]

	return index
}