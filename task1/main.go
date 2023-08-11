package main

import "fmt"

type Car struct {
	Type         string
	FuelInLiters float64
}

func (c Car) EstimatedDistance() float64 {
	fuelEfficiency := 1.5 // L/mill
	return c.FuelInLiters / fuelEfficiency
}

func main() {
	myCar := Car{
		Type:         "Sedan",
		FuelInLiters: 30,
	}

	estimatedDistance := myCar.EstimatedDistance()
	fmt.Printf("Perkiraan jarak yang bisa ditempuh: %.2f mill\n", estimatedDistance)
}
