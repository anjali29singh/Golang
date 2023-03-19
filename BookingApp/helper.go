package main

import "strings"

func validUserInput(fname string, lname string, email string, tickets uint, remainingTickets uint) bool {
	isValidName := len(fname) >= 2 && len(lname) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicket := tickets > 0 && remainingTickets >= tickets

	if isValidEmail && isValidTicket && isValidName {
		return true
	} else {
		return false
	}

}
