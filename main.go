package main

import (
	"fmt"
	"strings"
)

func main(){
	var conferenceName = "Go conference" 
	var remainingTickets uint = 50 
	var bookings []string

	fmt.Println("welcome to the", conferenceName, "booking application")
	fmt.Println("Number of tickets available is", remainingTickets)

	for {
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

		//user input validation
		isValidName := len(firstName) >= 2 && len(lastName) >= 2  
		isValidEmail := len(email) >= 2 && strings.Contains(email, "@") 
		isValidTicketNumber := userTicket > 0 && userTicket <= remainingTickets
		//isValidCity := city == "Singapore" || city == "USA"

	if isValidName && isValidEmail && isValidTicketNumber{
		if userTicket > remainingTickets {
			fmt.Printf("We've only %v ticket left, you can not book %v tickets.\n", remainingTickets, userTicket)
		} else {
			remainingTickets = remainingTickets - uint(userTicket)
			//bookings[0] = firstName + " " + lastName 
			bookings = append(bookings, firstName + " " + lastName)

			fmt.Printf("Thanks You!! %v %v for booking %v tickets. You will recive a confirmation email at %v\n", firstName, lastName, userTicket, email)

			fmt.Printf("%v tickets left now\n", remainingTickets)

			var firstNames []string

			for _, booking := range bookings{
				// range provides the index and value of that element 
				// range iterate over elemnt for different data structure
				// blank identifier '_' for unused variable 
				fmt.Printf("These are all our bookings:%v\n", booking)
				var names = strings.Fields(booking) 
				// split the string with white space as separator 
				// and return a slice with split elements
				firstNames = append(firstNames, names[0])
			}
			
			fmt.Printf("These are all list of firstName from our bookings: %v\n", firstNames)

			// when reached 0 we need to end the application code 
			if remainingTickets == 0{
				//exit the code
				fmt.Println("All ticket has been sold out, come next year")
				break
		}
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


		