package main

import (
	"fmt"
)

func main() {
	var digit int

	fmt.Print("Enter five digits: ")
	fmt.Scanln(&digit)
	for digit < 9999 || digit > 99999 {
		fmt.Print("Five digit only: ")
		fmt.Scanln(&digit)
	}

	firstDigit := digit / 10000
	secondDigit := digit / 1000 % 10
	//thirdDigit := digit / 100 % 10
	fourthDigit := digit / 10 % 10
	fifthDigit := digit % 10

	if firstDigit == fifthDigit && secondDigit == fourthDigit {
		fmt.Println("This is a palindrome")
	} else {
		fmt.Println("This is not a palindrome")
	}
	
}
