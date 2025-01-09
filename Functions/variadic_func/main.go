package main

import "fmt"

func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func main() {
	fmt.Println(sum(1, 2, 3, 4)) // Output: 10
}


// Key Points:
// The ... before the type allows the function to take multiple arguments of that type.
// You can pass a slice using func(slice...).