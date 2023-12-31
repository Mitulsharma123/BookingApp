package validation

import "strings"

var MyVar = "some random value"

func ValidateUserInput(firstName string, lastName string, email string,userTicket uint, remainingTickets uint)(bool, bool, bool){
	isValidName := len(firstName) >= 2 && len(lastName) >= 2  
	isValidEmail := len(email) >= 2 && strings.Contains(email, "@") 
	isValidTicketNumber := userTicket > 0 && userTicket <= remainingTickets
	return isValidName, isValidEmail, isValidTicketNumber
}