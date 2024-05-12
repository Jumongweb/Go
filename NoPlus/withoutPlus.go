package main

import "fmt"

func getSum(num1, num2 int) int {
	return num1 - (-num2)
}

func main() {
	fmt.Println(getSum(10, 20))
	fmt.Println(getSum(0, -5))
}
