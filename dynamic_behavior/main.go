package main

import "fmt"

func dynamicFunc(op string) func(int, int) int {
	if op == "add" {
		return func(a, b int) int { return a + b }
	}
	return func(a, b int) int { return a - b }
}

func main() {
	add := dynamicFunc("add")
	fmt.Println(add(3, 4)) // Output: 7
}
