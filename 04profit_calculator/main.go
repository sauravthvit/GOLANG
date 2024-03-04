package main

import (
	"fmt"
)

func main() {

	var taxRate float64
	var revenue float64
	var expenses float64

	fmt.Print("Revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Tax Rate: ")
	fmt.Scan(&taxRate)

	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	fmt.Println("EBT:", ebt)
	fmt.Println("Profit:", profit)
	fmt.Println("Ratio:", ratio)
	fmt.Printf("EBT: %v\nProfit: %v\n", ebt, profit) //You can use printf and multiple placeholders and write the variables in correct order.
}
