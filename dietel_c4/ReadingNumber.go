package main

import "fmt"

func main() {
	sum := 0
	number := 0

	fmt.Print("Enter specified sum: ")
	var specifiedSum int
	fmt.Scanln(&specifiedSum)

	for sum < specifiedSum {
		fmt.Print("Enter number: ")
		fmt.Scanln(&number)
		sum += number
	}

	fmt.Println("The total number entered is : ", sum)

}
