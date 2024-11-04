package main

import (
	"fmt"
	"os"
)

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)

	os.WriteFile("balance.txt", []byte(balanceText), 0644)
}

func main() {
	accountBalance := 1000.0

	fmt.Print("Welcome to the Golang Bank\n\n")

loop:
	for {

		fmt.Println("What do you want to do?")
		fmt.Println("1. Check the balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

		var choice int

		fmt.Print("Select your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Printf("Your balance is $%.2f\n", accountBalance)
		case 2:
			var depositAmount float64

			fmt.Print("Your deposit: ")
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")

				continue
			}

			accountBalance += depositAmount

			fmt.Printf("Balance updated! New amount: $%.2f\n", accountBalance)

			writeBalanceToFile(accountBalance)
		case 3:
			var withdrawalAmount float64

			fmt.Print("Withdrawal amount: ")
			fmt.Scan(&withdrawalAmount)

			if withdrawalAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")

				continue
			}

			if withdrawalAmount > accountBalance {
				fmt.Println("Invalid amount. You can't withdraw more than you have.")

				continue
			}

			accountBalance -= withdrawalAmount

			fmt.Printf("Balance updated! New amount: $%.2f\n", accountBalance)

			writeBalanceToFile(accountBalance)
		default:
			fmt.Println("Goodbye!")

			break loop
		}

		fmt.Println("Your choice:", choice)
	}

	fmt.Println("Thanks for using the Golang Bank!")
}
