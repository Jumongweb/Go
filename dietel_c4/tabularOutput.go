package main

import "fmt"

func main() {
	fmt.Println("N1\tN2\tN3\tN4")
	count := 1
	for count != 6 {

		fmt.Print(count, "\t")
		fmt.Print(count*count, "\t")
		fmt.Print(count*count*count, "\t")
		fmt.Print(count*count*count*count, "\t")
		fmt.Print("\n")
		count++
	}
}
