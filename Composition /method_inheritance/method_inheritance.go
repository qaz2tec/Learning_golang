package main

import "fmt"

type Address struct {
	City string
}

func (a Address) PrintCity() {
	fmt.Println("City:", a.City)
}

type Person struct {
	Name    string
	Address // Embedded struct
}

func main() {
	p := Person{
		Name: "Dhruv",
		Address: Address{
			City: "Delhi",
		},
	}

	p.PrintCity() // Call method from Address through Person
}
