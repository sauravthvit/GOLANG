package main

import (
	"fmt"
)

func main2() {

	taxRate := 6.4
	var revenue float64
	var expenses float64

	fmt.Print("Revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Expenses: ")
	fmt.Scan(&expenses)

	ebt := revenue - expenses
	profit := ebt - taxRate
	ratio := ebt / profit
	fmt.Println(ebt)
	fmt.Println(profit)
	fmt.Println(ratio)
}

func main() {
	main2()
}
