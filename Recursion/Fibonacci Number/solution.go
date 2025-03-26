func fib(n int) int {
	var result int
	if n <= 1 {
		return n
	}

	if val, ok := mp[n]; ok {
		return val
	}

	mp[n] = fib(n-1) + fib(n-2)
	result = mp[n]
	return result
}

var mp = make(map[int]int)