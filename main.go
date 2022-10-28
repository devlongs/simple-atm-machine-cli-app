package main

import (
	"fmt"
)

func main(){
	var pin int
	var changedPin int
	var userOption int

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

}

// func getUserInput() (string, string, string, uint) {
// 	var firstName string
// 	var lastName string
// 	var email string
// 	var userTickets uint

// 	fmt.Println("Enter Your First Name: ")
// 	fmt.Scanln(&firstName)

// 	fmt.Println("Enter Your Last Name: ")
// 	fmt.Scanln(&lastName)

// 	fmt.Println("Enter Your Email: ")
// 	fmt.Scanln(&email)

// 	fmt.Println("Enter number of tickets: ")
// 	fmt.Scanln(&userTickets)

// 	return firstName, lastName, email, userTickets
// }