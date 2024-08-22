package exercices

var memo = make(map[int]int)

func FibonacciMemoized(n int) int {

	if val, ok := memo[n]; ok {
		return val
	}
	if n == 0 {
		return 0
	}
	if n < 2 {
		return 1
	}

	memo[n] = FibonacciMemoized(n-1) + FibonacciMemoized(n-2)

	return memo[n]

}
