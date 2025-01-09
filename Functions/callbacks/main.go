package main

import "fmt"

func process(numbers []int, callback func(int)) {
	for _, n := range numbers {
		callback(n)
	}
}

func main() {
	process([]int{1, 2, 3}, func(num int) {
		fmt.Println(num * 2) // Output: 2, 4, 6
	})
}
