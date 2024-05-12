package main

import (
	"fmt"
	"strings"
	"time"
)

type PurchaseDetails struct {
	ItemName       string
	QuantityOfItem int
	AmountPerUnit  float32
	Total          float32
}

//func (p *PurchaseDetails) calculateTotal() float32 {
//	return p.AmountPerUnit * float32(p.QuantityOfItem)
//}

var customerName string
var itemName string
var quantity int
var price float32
var total float32
var cashierName string
var subTotal float32
var discount float32
var amountPaid float32

func main() {

	var itemsDetails []PurchaseDetails

	doubleLine := strings.Repeat("=", 60)
	singleLine := strings.Repeat("-", 60)

	fmt.Println("What is the customer's name?")
	fmt.Scanln(&customerName)

	for {
		fmt.Println("What did the customer buy?")
		fmt.Scanln(&itemName)

		fmt.Println("How many pieces?")
		fmt.Scanln(&quantity)

		fmt.Println("How much per unit?")
		fmt.Scanln(&price)

		fmt.Println("Add more items")
		var addMoreItems string
		fmt.Scanln(&addMoreItems)

		total = price * float32(quantity)
		subTotal += total
		purchaseDetails := PurchaseDetails{ItemName: itemName, QuantityOfItem: quantity, AmountPerUnit: price, Total: total}

		itemsDetails = append(itemsDetails, purchaseDetails)

		if strings.ToLower(addMoreItems) == "no" {
			break
		}
	}

	fmt.Println("What is your name?")
	fmt.Scanln(&cashierName)

	fmt.Println("How much discount will he get?")
	var discount float32
	fmt.Scanln(&discount)

	currentTime := time.Now()

	fmt.Println(itemsDetails)

	fmt.Println()
	fmt.Println("SEMICOLON STORES")
	fmt.Println("MAIN BRANCH")
	fmt.Println("LOCATION: 312, HERBERT MERCAULY WAY, SABO YABA, LAGOS.")
	fmt.Println("Tel: 03293828343")
	fmt.Println("Date:", currentTime)
	fmt.Println("Cashier:", cashierName)
	fmt.Println("Customer:", customerName)

	fmt.Println(doubleLine)
	fmt.Printf("%s\t%10s\t%10s\t%10s\n", "ITEM", "QTY", "PRICE", "TOTAL(NGN)")
	fmt.Println(singleLine)

	for index := 0; index < len(itemsDetails); index++ {
		fmt.Printf("%v\t%10v\t%10v\t%10v\n", itemsDetails[index].ItemName, itemsDetails[index].QuantityOfItem, itemsDetails[index].AmountPerUnit, itemsDetails[index].Total)
	}
	fmt.Println(singleLine)
	discount = subTotal * (discount / 100)
	vat := subTotal * (17.50 / 100)
	billTotal := subTotal - discount + vat

	fmt.Printf("%40s %.2f\n", "Subtotal: ", subTotal)
	fmt.Printf("%40s %.2f\n", "Discount: ", discount)
	fmt.Printf("%40s %.2f\n", "VAT @ 17.50%%: ", vat)
	fmt.Println(doubleLine)
	fmt.Printf("%40s %.2f\n", "Bill Total: ", billTotal)
	fmt.Println(doubleLine)
	fmt.Println("THIS IS NOT A RECEIPT KINDLY PAY ", billTotal)
	fmt.Println(doubleLine)

	fmt.Println("How much did the customer give to you? ")
	fmt.Scanln(&amountPaid)

	fmt.Println()
	fmt.Println("SEMICOLON STORES")
	fmt.Println("MAIN BRANCH")
	fmt.Println("LOCATION: 312, HERBERT MERCAULY WAY, SABO YABA, LAGOS.")
	fmt.Println("Tel: 03293828343")
	fmt.Println("Date:", currentTime)
	fmt.Println("Cashier:", cashierName)
	fmt.Println("Customer:", customerName)

	fmt.Println(doubleLine)
	fmt.Printf("%s\t%10s\t%10s\t%10s\n", "ITEM", "QTY", "PRICE", "TOTAL(NGN)")
	fmt.Println(singleLine)

	for index := 0; index < len(itemsDetails); index++ {
		fmt.Printf("%v\t%10v\t%10v\t%10v\n", itemsDetails[index].ItemName, itemsDetails[index].QuantityOfItem, itemsDetails[index].AmountPerUnit, itemsDetails[index].Total)
	}
	fmt.Println(singleLine)
	fmt.Printf("%40s %.2f\n", "Subtotal: ", subTotal)
	fmt.Printf("%40s %.2f\n", "Discount: ", discount)
	fmt.Printf("%40s %.2f\n", "VAT @ 17.50%%: ", vat)
	fmt.Println(doubleLine)
	fmt.Printf("%40s %.2f\n", "Bill Total: ", billTotal)
	fmt.Printf("%40s %.2f\n", "Amount Paid", amountPaid)
	var balance float32
	balance = amountPaid - billTotal

	fmt.Printf("%40s %.2f\n", "Balance:", balance)
	fmt.Println(doubleLine)
	fmt.Println("\t\t\tTHANK YOU FOR YOU PATRONAGE")
	fmt.Println(doubleLine)
}
