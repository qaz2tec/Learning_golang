package main

import "fmt"

func apply(fn func(int) int, num int) int {
	return fn(num)
}

func square(x int) int {
	return x * x
}

func main() {
	fmt.Println(apply(square, 5)) // Output: 25
}
