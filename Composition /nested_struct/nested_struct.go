package main

import "fmt"

type Engine struct {
	Horsepower int
}

type Wheels struct {
	Count int
}

type Car struct {
	Brand  string
	Engine // Embedded struct
	Wheels // Embedded struct
}

func main() {
	myCar := Car{
		Brand: "Tesla",
		Engine: Engine{
			Horsepower: 500,
		},
		Wheels: Wheels{
			Count: 4,
		},
	}

	fmt.Printf("%s car with %d HP and %d wheels.\n", myCar.Brand, myCar.Horsepower, myCar.Count)
}
