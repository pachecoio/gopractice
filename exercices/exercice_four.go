package exercices

import (
	"sync"
)

func fibonacci(n int) []int {
	fib := make([]int, n)
	fib[0] = 0
	fib[1] = 1
	for i := 2; i < n; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}
	return fib
}

func FibonacciMemoized(n int) int {
	var wg sync.WaitGroup

	wg.Add(1)
	var fibonacciSequence []int
	go func() {
		fibonacciSequence = fibonacci(n)
		wg.Done()
	}()

	wg.Wait()

	return fibonacciSequence[n-1] + fibonacciSequence[n-2]

}
