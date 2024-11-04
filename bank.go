package main

import "fmt"

func main() {
	accountBalance := 100.0

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

		accountBalance += depositAmount

		fmt.Printf("Balance updated! New amount: $%.2f\n", accountBalance)
	}

	fmt.Println("Your choice:", choice)
}
