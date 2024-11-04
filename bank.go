package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func displayUserOptions() {
	fmt.Println("\nWhat do you want to do?")
	fmt.Println("1. Check the balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Print("4. Exit\n\n")
}

func getUserInput() int {
	var choice int
	fmt.Print("Select your choice: ")
	fmt.Scan(&choice)

	return choice
}

func readBalanceFromFile() (float64, error) {
	defaultBalance := 0.0
	data, err := os.ReadFile(accountBalanceFile)

	if err != nil {
		return defaultBalance, errors.New("failed to find balance file")
	}

	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)

	if err != nil {
		return defaultBalance, errors.New("failed to parse stored balance")
	}

	return balance, nil
}

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)

	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
}

func getAccountBalance() float64 {
	accountBalance, err := readBalanceFromFile()

	if err != nil {
		fmt.Print("---------\n")
		fmt.Print("ERROR: ")
		fmt.Print(err, ". Setting the account balance to $0.0.\n")
		fmt.Print("---------\n\n")

		writeBalanceToFile(accountBalance)
	}

	return accountBalance
}

func updateBalance(amount float64) {
	fmt.Printf("Balance updated! New balance: $%.2f\n", amount)

	writeBalanceToFile(amount)
}

func checkBalance() {
	accountBalance := getAccountBalance()

	fmt.Printf("Your balance is $%.2f\n", accountBalance)
}

func depositMoney() error {
	accountBalance := getAccountBalance()
	var depositAmount float64

	fmt.Print("Your deposit: ")
	fmt.Scan(&depositAmount)

	if depositAmount <= 0 {
		return errors.New("invalid amount. Must be greater than 0")
	}

	accountBalance += depositAmount
	updateBalance(accountBalance)

	return nil
}

func withdrawMoney() error {
	accountBalance := getAccountBalance()
	var withdrawalAmount float64

	fmt.Print("Withdrawal amount: ")
	fmt.Scan(&withdrawalAmount)

	if withdrawalAmount <= 0 {
		return errors.New("invalid amount. Must be greater than 0")
	}

	if withdrawalAmount > accountBalance {
		return errors.New("invalid amount. You can't withdraw more than you have")
	}

	accountBalance -= withdrawalAmount
	updateBalance(accountBalance)

	return nil
}

func closeBank() {
	fmt.Print("\nThanks for using the Golang Bank\n")
	fmt.Println("Goodbye!")
}

func main() {
	fmt.Print("Welcome to the Golang Bank\n")

	for {
		displayUserOptions()

		choice := getUserInput()

		switch choice {
		case 1:
			checkBalance()
		case 2:
			err := depositMoney()

			if err != nil {
				fmt.Println(err)
			}
		case 3:
			err := withdrawMoney()

			if err != nil {
				fmt.Println(err)
			}
		default:
			closeBank()

			return
		}
	}
}
