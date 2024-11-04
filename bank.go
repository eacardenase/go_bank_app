package main

import "fmt"

func main() {
	accountBalance := 1000.0

	fmt.Print("Welcome to the Bank Application\n\n")

	fmt.Println("What do you want to do?")
	fmt.Println("1. Check the balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")

	var choice int

	fmt.Print("Select your choice: ")
	fmt.Scan(&choice)

	if choice == 1 {
		fmt.Printf("Your balance is $%.2f\n", accountBalance)
	} else if choice == 2 {
		var depositAmount float64

		fmt.Print("Your deposit: ")
		fmt.Scan(&depositAmount)

		if depositAmount <= 0 {
			fmt.Println("Invalid amount. Must be greater than 0.")

			return
		}

		accountBalance += depositAmount

		fmt.Printf("Balance updated! New amount: $%.2f\n", accountBalance)
	} else if choice == 3 {
		var withdrawalAmount float64

		fmt.Print("Withdrawal amount: ")
		fmt.Scan(&withdrawalAmount)

		if withdrawalAmount <= 0 {
			fmt.Println("Invalid amount. Must be greater than 0.")

			return
		}

		if withdrawalAmount > accountBalance {
			fmt.Println("Invalid amount. You can't withdraw more than you have.")

			return
		}

		accountBalance -= withdrawalAmount

		fmt.Printf("Balance updated! New amount: $%.2f\n", accountBalance)
	} else {
		fmt.Println("Goodbye!")
	}

	fmt.Println("Your choice:", choice)
}
