package main

import (
	"fmt"
	"math"
)

func main() {
	lowest := math.MaxInt64
	highest := math.MinInt64

	fmt.Print("How many number do you want to enter: ")
	var inputCount int
	fmt.Scanln(&inputCount)

	for count := 0; count < inputCount; count++ {
		fmt.Print("Enter number: ")
		var input int
		fmt.Scanln(&input)
		if input < lowest {
			lowest = input
		}
		if input > highest {
			highest = input
		}
	}

	fmt.Print("The sum of extremes is ", lowest+highest)
}
