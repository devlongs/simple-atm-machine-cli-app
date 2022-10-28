package main

import (
	"fmt"
)

func main(){
	var pin int
	fmt.Println("Hello dear user, Welcome to the ATM machine")
	
	fmt.Println("Please, kindly type your pin")
	fmt.Scanln(&pin)

	fmt.Println(pin)

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