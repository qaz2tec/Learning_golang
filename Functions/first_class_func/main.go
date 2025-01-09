package main

import "fmt"

func greet(name string) string {
	return "Hello, " + name
}

func main() {
	f := greet
	fmt.Println(f("Go")) // Output: Hello, Go
}
