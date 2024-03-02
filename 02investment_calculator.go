package main

import (
	"fmt"
	"math"
)

func main() {
	//Created Variables
	// var a = 10
	// investmentAmount, years := 1000, 5.5
	const inflationRate = 6.5
	investmentAmount := 1000.0
	expectedReturnRate := 5.5 //Shortcut of Autosetting of variable type
	years := 10.0

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)
	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}
