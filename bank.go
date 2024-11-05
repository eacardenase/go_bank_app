package main

import (
	"fmt"
)

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
