func thirdMax(nums []int) int {
	var result int
	mp := make(map[int]bool)
	for _, val := range nums {
		mp[val] = true
	}

	arr := make([]int, 0, len(nums))
	for k, _ := range mp {
		arr = append(arr, k)
	}

	k := 3
	if len(arr) < k {
		return quickSelect(arr, 0, len(arr)-1, len(arr)-1)
	}
	result = quickSelect(arr, 0, len(arr)-1, len(arr)-k)
	return result
}

func quickSelect(arr []int, i, j int, k int) int {
	if i >= j {
		return arr[i]
	}
	index := partition(arr, i, j)
	if k == index {
		return arr[k]
	} else if k < index {
		return quickSelect(arr, i, index-1, k)
	}

	return quickSelect(arr, index+1, j, k)
}

func partition(arr []int, i, j int) int {
	var index int
	index = i
	pivot := arr[j]
	for left := i; left < len(arr); left++ {
		if arr[left] < pivot {
			arr[left], arr[index] = arr[index], arr[left]
			index++
		}
	}
	arr[index], arr[j] = arr[j], arr[index]

	return index
}