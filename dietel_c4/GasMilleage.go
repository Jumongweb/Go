package main

import (
	"fmt"
)

func main() {
	var trip int
	totalMilesDriven := 0.0
	totalGallonsUsed := 0.0

	for trip != -1 {
		fmt.Print("Enter numbers of miles: ")
		var miles float64
		fmt.Scanln(&miles)

		fmt.Print("Enter number of gallons used: ")
		var gallons float64
		fmt.Scanln(&gallons)

		totalMilesDriven += miles
		totalGallonsUsed += gallons

		milesPerGallon := miles / gallons
		fmt.Println(milesPerGallon)

		fmt.Print("Enter number of trip (Enter -1 to exit): ")
		fmt.Scanln(&trip)

	}

	totalMilesPerGallon := totalMilesDriven / totalGallonsUsed
	fmt.Printf("Total miles per gallon used is: %.2f\n\n", totalMilesPerGallon)
}
