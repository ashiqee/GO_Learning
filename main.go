package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickers = 50
	var remainingTickets = 50

	fmt.Printf("ConfernceTickets is %T, remaining Tickets is %T, ConferenceName is %T\n", conferenceTickers, remainingTickets, conferenceName)

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickers, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var userName string
	var userTickets int
	// ask user for their name

	userName = "Abid"
	userTickets = 2
	fmt.Printf("User %v booked %v tickets.\n", userName, userTickets)

}
