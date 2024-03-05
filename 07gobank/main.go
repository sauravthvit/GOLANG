package main

import (
	"fmt"
)

func main() {

	accountBalance := 1000.0

	fmt.Println("Welcome to Go Bank!")
	fmt.Println("What do you want to do?")
	fmt.Println("1. Check Balance")
	fmt.Println("2. Deposit Money")
	fmt.Println("3. Withdraw Money")
	fmt.Println("4. Exit")

	var choice int
	fmt.Print("Your choice: ")
	fmt.Scan(&choice)

	if choice == 1 {
		fmt.Println("Your balance is:", accountBalance)
	} else if choice == 2 {
		fmt.Println("Your deposit amount: ")
		var depositAmount float64
		fmt.Scan(&depositAmount)

		if depositAmount <= 0.0 {
			fmt.Println("Invalid deposit amount")
			return
		}

		accountBalance += depositAmount
		fmt.Println("Balance Updated!\nNew amount: ", accountBalance)
	} else if choice == 3 {
		fmt.Println("Your withdrwawl amount: ")
		var withdrawlamount float64
		fmt.Scan(&withdrawlamount)

		if withdrawlamount <= 0.0 || withdrawlamount > accountBalance {
			fmt.Println("Invalid amount. You can't withdraw more than you have.")
			return
		}

		accountBalance -= withdrawlamount
		fmt.Println("Balance Updated!\nNew amount: ", accountBalance)
	} else {
		fmt.Println("Thank you! Visit Again!")
	}
}
