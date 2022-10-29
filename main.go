package main

import (
	"fmt"
)

func main() {
	var pin int
	var changedPin int = 0000 // default pin
	var userOption int
	var accountBalance int
	var withdrawalAmount int
	var depositAmount int
	var performAnotherTransaction int

	fmt.Println("Hello dear user, Welcome to the ATM machine")

	for {
		// ask user for a numeric pin
		fmt.Println("Please, kindly type your pin")
		fmt.Scanln(&pin)

		fmt.Println("Press 1 to change pin")
		fmt.Println("Press 2 to check account balance")
		fmt.Println("Press 3 to withdraw funds")
		fmt.Println("Press 4 to deposit funds")

		fmt.Scanln(&userOption)

		if userOption == 1 {
			fmt.Println("Type a four digit number as your new pin")
			fmt.Scanln(&changedPin)
			fmt.Printf("Your old pin %v is successfully changed to your new pin %v\n", pin, changedPin)
			pin = changedPin

			fmt.Println("Press 1 if you want to perform another transaction or press 0 if you want to exit.")

			fmt.Scanln(&performAnotherTransaction)
			if performAnotherTransaction == 1 {
				continue
			} else {
				break
			}
		}

		if userOption == 2 {
			fmt.Println("Your account balance is:", accountBalance)

			fmt.Println("Press 1 if you want to perform another transaction or press 0 if you want to exit.")
			fmt.Scanln(&performAnotherTransaction)
			if performAnotherTransaction == 1 {
				continue
			} else {
				break
			}
		}

		if userOption == 3 {
			fmt.Println("How much do you want to withdraw?")
			fmt.Scanln(&withdrawalAmount)
			accountBalance := accountBalance - withdrawalAmount
			fmt.Printf("Withdrawal successfully!. Your new balance is %v\n", accountBalance)

			fmt.Println("Press 1 if you want to perform another transaction or press 0 if you want to exit.")
			fmt.Scanln(&performAnotherTransaction)
			if performAnotherTransaction == 1 {
				continue
			} else {
				break
			}
		}

		if userOption == 4 {
			fmt.Println("How much do you want to deposit?")
			fmt.Scanln(&depositAmount)
			accountBalance = accountBalance + depositAmount
			fmt.Printf("Deposit successfully!. Your new balance is %v\n", accountBalance)

			fmt.Println("Press 1 if you want to perform another transaction or press 0 if you want to exit.")
			fmt.Scanln(&performAnotherTransaction)
			if performAnotherTransaction == 1 {
				continue
			} else {
				break
			}
		}

	}
}
