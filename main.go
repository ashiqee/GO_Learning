package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickers = 50
	var remainingTickets = 50

	fmt.Println("Welcome to", conferenceName, "booking application")
	fmt.Println("We have total of", conferenceTickers, "tickets and", remainingTickets, "are still available")
	fmt.Println("Get your tickets here to attend")

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickers, remainingTickets)

}
