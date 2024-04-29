package main

import "fmt"

func main() {
	var name string
	var earning int

	for {
		fmt.Print("Enter citizen name: ")
		fmt.Scan(&name)
		fmt.Print("Enter citizen earnings: ")
		fmt.Scan(&earning)

		var tax float64
		if earning == -1 {
			break
		}
		if earning <= 30000 {
			tax = float64(earning) * 0.15
			fmt.Printf("%s your tax is $%.2f\n", name, tax)
		} else {
			tax = float64(earning) * 0.2
			fmt.Printf("%s your tax is $%.2f\n", name, tax)
		}
	}

}
