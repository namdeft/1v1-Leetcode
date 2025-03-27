/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {
	var result int
	left := 0
	right := n - 1

	for left <= right {
		mid := (left + right) / 2

		if guess(mid+1) == -1 {
			right = mid - 1
		}

		if guess(mid+1) == 1 {
			left = mid + 1
		}

		if guess(mid+1) == 0 {
			result = mid + 1
			break
		}
	}

	return result
}