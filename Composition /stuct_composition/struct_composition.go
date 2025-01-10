package main

import "fmt"

// Base struct
type Address struct {
	City  string
	State string
}

// Composing struct
type Person struct {
	Name    string
	Age     int
	Address // Embedded struct
}

func main() {
	p := Person{
		Name: "Arjun",
		Age:  25,
		Address: Address{
			City:  "Pune",
			State: "Maharashtra",
		},
	}

	// Access embedded fields directly
	fmt.Println("Name:", p.Name)
	fmt.Println("City:", p.City)   // Access field from Address
	fmt.Println("State:", p.State) // Access field from Address
	fmt.Println("State:", p.Age) 
}
