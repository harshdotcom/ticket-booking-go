package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

const confrenceTickets int = 50

var confrenceName = "Go Confrence"
var remainingTickets uint = 50
var bookings = make([]UserData, 0)

// struct can we compare as classes in Java
type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUsers()

	for {

		firstName, lastName, email, userTickets := getUsersInput()
		isValidEmail, isValidName, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidEmail && isValidName && isValidTicketNumber {

			bookTickets(remainingTickets, userTickets, firstName, lastName, email, confrenceName)
			wg.Add(1)
			go sendTicket(userTickets, firstName, lastName, email)

			// Prient First Name
			firstNames := printFirstName()

			fmt.Printf("First Names of Bookings %v\n", firstNames)

			if remainingTickets <= 0 {
				// end the program
				fmt.Println("Our Confrence is booked out. Come back next year.")
				break
			}
		} else {
			fmt.Printf("We only have %v tickets remaining, so you can't book %v tickets\n", remainingTickets, userTickets)
		}

	}

}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", confrenceName)
	fmt.Printf("We have total of %v tickets and %v Tickets are still available.\n", confrenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func printFirstName() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUsersInput() (string, string, string, uint) {
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
	return firstName, lastName, email, userTickets
}

func bookTickets(remainingTickets uint, userTickets uint, firstName string, lastName string, email string, confrenceName string) {
	remainingTickets = remainingTickets - userTickets

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of Booking is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, confrenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(5 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("###############")
	fmt.Printf("Sending tickets:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("###############")
	wg.Done()
}
