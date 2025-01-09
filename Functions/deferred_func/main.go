package main

import "fmt"

func main() {
	defer fmt.Println("World")
	fmt.Println("Hello")
	// Output:
	// Hello
	// World
}
