package main

import (
	"fmt"
	"math"
)

func main1() {
	//Created Variables
	// var a = 10
	// investmentAmount, years := 1000, 5.5
	const inflationRate = 6.5 //Can't changed thereafter
	var investmentAmount float64
	var years float64
	expectedReturnRate := 5.5 //Shortcut of Autosetting of variable type

	fmt.Print("Investment Amount: ")
	//Scans the terminal for input
	fmt.Scan(&investmentAmount) //Using pointers

	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)
	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}

func main() {
	main1()
}
