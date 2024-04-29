package main

import "fmt"

func factorial(number int) int {
	for {
		if number < 1 {
			return 0
		} else {
			for count := number - 1; count > 0; count-- {
				number *= count
			}
			return number
		}
	}
}

func main() {
	fmt.Print("Enter a number: ")
	var number int
	fmt.Scanln(&number)
	fmt.Println(factorial(number))
}
