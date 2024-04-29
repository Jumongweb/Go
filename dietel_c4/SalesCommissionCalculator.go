package main

import "fmt"

func main() {
	var sale int
	var salary float64
	totalSales := 0.0
	for {
		if sale == -1 {
			break
		}
		fmt.Println("Enter weekly sales: ")
		fmt.Scan(&sale)
		//totalSales += sale
	}
	salary = (0.09 * totalSales) + 200
	fmt.Println("Your salary for this week is: ", salary)
}
