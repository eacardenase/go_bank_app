package main

import (
	"fmt"

	"github.com/Pallinder/go-randomdata"
)

func main() {
	fmt.Print("Welcome to the Golang Bank\n")
	fmt.Println("Reach us 24/7 at", randomdata.PhoneNumber())

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
