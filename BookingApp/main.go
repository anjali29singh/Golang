package main

import "fmt"

func main() {

	var conferenceName string = "Go Conference"
	const conferenceTickets int = 60
	var remainingTickets int = 60
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var userName string
	var userTickets int
	//ask user for their name
	//for input use Scan()
	fmt.Scan(&userName)
	fmt.Scan(&userTickets)
	fmt.Printf("User %v booked %v tickets.\n", userName, userTickets)
}
