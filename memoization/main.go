package main

import "fmt"

func memoize(fn func(int) int) func(int) int {
	cache := make(map[int]int)
	return func(n int) int {
		if result, exists := cache[n]; exists {
			return result
		}
		result := fn(n)
		cache[n] = result
		return result
	}
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	memoizedFactorial := memoize(factorial)
	fmt.Println(memoizedFactorial(5)) // Output: 120
	fmt.Println(memoizedFactorial(5)) // Cached result: 120
}
