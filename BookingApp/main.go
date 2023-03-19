package main

import (
	"fmt"
)

func greetUser(fname string) {
	fmt.Printf("Welcome to our conference %v", fname)
}

func main() {

	var conferenceName string = "Go Conference"
	const conferenceTickets uint = 60
	var remainingTickets uint = 60

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Printf("Get your tickets here to attend.\n")

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

	/*isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicket := userTickets > 0 && userTickets <= int(remainingTickets)
	*/
	// bookings[0] = firstName + " " + lastName //in case of array

	isValidInput := validUserInput(firstName, lastName, email, uint(userTickets), remainingTickets)

	//in case of slice
	if isValidInput {
		bookings = append(bookings, firstName+" "+lastName)

		remainingTickets = remainingTickets - uint(userTickets)

		fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve a confirmation at email at %v.\n", firstName, lastName, userTickets, email)
		fmt.Printf("%v tickets remaining for %v.\n", remainingTickets, conferenceName)

		fmt.Println(bookings, len(bookings))

		//loops
		firstNames := []string{}
		for index, booking := range bookings {
			fmt.Println(index, booking)
			firstNames = append(firstNames, booking)
		}
		//if do not want to use some variable then _
		var _ = "anjali"
		//if else
		if remainingTickets == 0 {
			fmt.Printf("Our conference is booked out.Come back next year....\n")
		} else {
			fmt.Printf("do you want to buy more tickets.\n")
		}

		//encapsulation - creating function for each task to make code readable and clear

		greetUser("anjali")
	} else {
		fmt.Println("Please enter valid inputs")
	}

	//switch statements
	/*city := "London"

	switch city {
	case "India":
		//execute code for booking India conference tickets


	case "Berlin":
		//execute code for Berlin

	case "Hong Kong":

	default:
		//if none of the above case matches

	}
	*/

	//packages

}
