package main

import "fmt"

func main() {
	confrenceName := "Go Confrence"
	const confrenceTickets int = 50
	var remainingTickets = 50

	fmt.Printf("Welcome to %v booking application\n", confrenceName)
	fmt.Printf("We have total of %v tickets and %v Tickets are still available.\n", confrenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var userName string
	var userTickets int

	fmt.Scan(&userName)
	fmt.Scan(&userTickets)

	fmt.Printf("User %v booked %v tickets.\n", userName, userTickets)

}
