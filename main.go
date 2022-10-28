package main

import (
	"fmt"
)

func main(){
	var pin int
	var changedPin int
	var userOption int
	var accountBalance int

	fmt.Println("Hello dear user, Welcome to the ATM machine")
	
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
	}

	if userOption == 2 {
		fmt.Println("Your account balance is:", accountBalance)
	}

}
