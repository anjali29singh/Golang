package main

import "fmt"

func main() {

	var conferenceName string = "Go Conference"
	const conferenceTickets uint = 60
	var remainingTickets uint = 60
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var firstName string
	var lastName string
	var email string
	var userTickets int

	//arrays

	// var bookings = [60]string{""}
	//slices
	var bookings []string

	//ask user for their name
	//for input use Scan()
	fmt.Println("enter your first name:")
	fmt.Scan(&firstName)

	fmt.Println("enter your last name:")
	fmt.Scan(&lastName)

	fmt.Println("enter your email id:")
	fmt.Scan(&email)

	fmt.Println("enter number of tickets:")
	fmt.Scan(&userTickets)
	// bookings[0] = firstName + " " + lastName //in case of array
	//in case of slice
	bookings = append(bookings, firstName+" "+lastName)

	remainingTickets = remainingTickets - uint(userTickets)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve a confirmation at email at %v.\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v.\n", remainingTickets, conferenceName)

	fmt.Println(bookings, len(bookings))

}
