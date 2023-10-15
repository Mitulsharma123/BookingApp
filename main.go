package main

import (
	"fmt"
	"strings"
)

//define variable as package level
// note : package level variable are created with var keyword 
const conferenceTicket int = 50
var conferenceName = "Go conference" 
var remainingTickets uint = 50 
var bookings = []string{}

func main(){

	greetUser()

	for {
		
		firstName,lastName,email,userTicket := getUserInput()

		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTicket, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber{
			bookTicket(firstName,lastName,email, remainingTickets)
			firstNames := getFirstNames()
			fmt.Printf("First names of bookings are %v", firstNames)

			// when reached 0 we need to end the application code 
			if remainingTickets == 0{
				//exit the code
				fmt.Println("All ticket has been sold out, come next year")
				break
			}
		} else {
			if !isValidName{
				fmt.Println("Name too short")
			}
			if !isValidEmail{
				fmt.Println("Not a valid email id, missing '@' ")
			}
			if !isValidTicketNumber{
				fmt.Println("Number of tickt entered is invalid")
			}
			fmt.Println("Youe input data is invalid, please try again !")
			}	
	}
}

func greetUser(){
	fmt.Printf("Welcome to the %v booking application!!!\n", conferenceName)
	fmt.Printf("We've total of %v tickets and tickets available are %v.\n", conferenceTicket, remainingTickets)
	fmt.Println("Get your tickets here to attend....")
}


func getUserInput() (string, string, string, uint){
	var firstName,lastName,email string
	var userTicket uint

	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email details:")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets:")
	fmt.Scan(&userTicket)

	return firstName,lastName,email,userTicket
}

func validateUserInput(firstName string, lastName string, email string,userTicket uint, remainingTickets uint)(bool, bool, bool){
	isValidName := len(firstName) >= 2 && len(lastName) >= 2  
	isValidEmail := len(email) >= 2 && strings.Contains(email, "@") 
	isValidTicketNumber := userTicket > 0 && userTicket <= remainingTickets
	return isValidName, isValidEmail, isValidTicketNumber
}

func bookTicket(firstName string, lastName string, email string, userTicket uint){
	remainingTickets = remainingTickets - uint(userTicket)
	bookings = append(bookings, firstName + " " + lastName)

	fmt.Printf("Thanks You!! %v %v for booking %v tickets. You will recive a confirmation email at %v\n", firstName, lastName, userTicket, email)
	fmt.Printf("%v tickets left now\n", remainingTickets)
}

func getFirstNames() []string{
	var firstNames []string
	for _, booking := range bookings{
		var names = strings.Fields(booking) 
		firstNames = append(firstNames, names[0])
		// range provides the index and value of that element 
		// range iterate over elemnt for different data structure
		// blank identifier '_' for unused variable 
		//fmt.Printf("These are all our bookings:%v\n", booking)
	
		// split the string with white space as separator 
		// and return a slice with split elements
	}		
	return firstNames
}





