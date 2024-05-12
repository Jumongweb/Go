package main

import "fmt"

func getSumOfDivisorsOf(number, length int) int {
	count := 1
	sum := 0

	for count < length {
		if count%number == 0 {
			sum += count
		}
	}
	return sum
}

func main() {
	fmt.Print("hello")
	getSumOfDivisorsOf(3, 30)
}
