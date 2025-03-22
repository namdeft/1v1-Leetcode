func findDuplicate(nums []int) int {
	var result int
	slow := nums[0]
	fast := nums[0]
	for {
		slow = nums[slow]
		fast = nums[nums[fast]]

		if slow == fast {
			break
		}
	}

	fast = nums[0]
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}

	result = slow
	return result
}