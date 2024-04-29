package main

import "fmt"

func main() {
	largest := 0
	count := 0

	for count < 10 {
		fmt.Printf("Enter number %d: ", count)
		var number int
		fmt.Scanln(&number)
		if largest < number {
			largest = number
		}
		count++
	}
	fmt.Println("The largest number is: ", largest)
}
