package main

import "fmt"

func main() {
	count := 0
	number := 0
	largest := 0
	secondLargest := 0

	for count < 10 {
		fmt.Print("Enter number: ")
		fmt.Scan(&number)

		if largest < number {
			secondLargest = largest
			largest = number
		} else if largest > secondLargest && number < largest {
			secondLargest = number
		}
		count++
	}
	fmt.Println("The largest number is: ", largest)
	fmt.Println("The secondLargest number is: ", secondLargest)
}
