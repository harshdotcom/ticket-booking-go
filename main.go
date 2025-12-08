package main

import (
	"fmt"
	"strings"
)

func main() {
	confrenceName := "Go Confrence"
	const confrenceTickets int = 50
	var remainingTickets uint = 50
	var bookings []string

	fmt.Printf("Welcome to %v booking application\n", confrenceName)
	fmt.Printf("We have total of %v tickets and %v Tickets are still available.\n", confrenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint

		fmt.Println("Enter Your First Name")
		fmt.Scan(&firstName)

		fmt.Println("Enter Your Last Name")
		fmt.Scan(&lastName)

		fmt.Println("Enter Your  Email")
		fmt.Scan(&email)

		fmt.Println("How many tickets you want to book")
		fmt.Scan(&userTickets)

		remainingTickets = remainingTickets - userTickets
		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, confrenceName)

		firstNames := []string{}
		for _, booking := range bookings {
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}

		fmt.Printf("First Names of Bookings %v\n", firstNames)

		if remainingTickets == 0 {
			// end the program
			fmt.Println("Our Confrence is booked out. Come back next year.")
			break
		}
	}

}
