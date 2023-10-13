package main

import "fmt"

func main(){
	var conferenceName = "Go conference" 
	const conferenceTicket uint = 50
	var remainingTickets int = 50 
	var bookings []string

	fmt.Println("welcome to the", conferenceName, "booking application")
	fmt.Println("Number of tickets available is", remainingTickets)

	for {
	var firstName,lastName,email string
	var userTicket int
	
	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email details:")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets:")
	fmt.Scan(&userTicket)

	remainingTickets = remainingTickets - userTicket
	//bookings[0] = firstName + " " + lastName 
	bookings = append(bookings, firstName + " " + lastName)

	fmt.Printf("Thanks You!! %v %v for booking %v tickets. You will recive a confirmation email at %v\n", firstName, lastName, userTicket, email)

	fmt.Printf("%v tickets left now\n", remainingTickets)

	fmt.Printf("These are all our bookings: %v \n", bookings)
	}
	
}