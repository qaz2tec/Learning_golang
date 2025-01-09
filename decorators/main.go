package main

import "fmt"

func logDecorator(fn func(string)) func(string) {
	return func(msg string) {
		fmt.Println("Before function call")
		fn(msg)
		fmt.Println("After function call")
	}
}

func printMessage(msg string) {
	fmt.Println(msg)
}

func main() {
	decorated := logDecorator(printMessage)
	decorated("Hello, Go!")
	// Output:
	// Before function call
	// Hello, Go!
	// After function call
}
