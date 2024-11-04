package main

import "fmt"

func main() {
	fmt.Print("Welcome to the Bank Application\n\n")

	fmt.Println("What do you want to do?")
	fmt.Println("1. Check the balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")

	var choice int

	fmt.Print("Your choice: ")
	fmt.Scan(&choice)

	fmt.Println("Your choice:", choice)
}
