package main

import (
	"bookingapp/validation" //import package <modulename/packagename>
	"fmt"
	"time"
)

//define variable as package level
// note : package level variable are created with var keyword
const conferenceTicket int = 50
var conferenceName = "Go conference" 
var remainingTickets uint = 50 
//create an empty list of maps
var bookings = make([]UserData, 0)

//define a structure for mixed data type
type UserData struct{
	firstName string
	lastName string
	email string
	userTicket uint
}

func main(){

	greetUser()

	for {
		
		firstName,lastName,email,userTicket := getUserInput()

		isValidName, isValidEmail, isValidTicketNumber := validation.ValidateUserInput(firstName, lastName, email, userTicket, remainingTickets)
		//calling function from package <packagename.Functionname>
		// now functionName should start with Capital Latter as calling from package 
		if isValidName && isValidEmail && isValidTicketNumber{
			bookTicket(firstName,lastName,email, userTicket)
			go sendTicket(userTicket, firstName, lastName, email)
			// it blocks the execution of code for 10 seconds; to overcome create a separate thread
			// go keyword creates a new go routine
			firstNames := getFirstNames()
			fmt.Printf("First names of bookings are %v.\n", firstNames)

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

func bookTicket(firstName string, lastName string, email string, userTicket uint){
	remainingTickets = remainingTickets - uint(userTicket)
	//create a emtry map
	var userData = UserData{
		firstName: firstName,
		lastName: lastName,
		email: email,
		userTicket: userTicket, 
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of Bookings is %v\n", bookings)

	fmt.Printf("Thanks You!! %v %v for booking %v tickets. You will recive a confirmation email at %v\n", firstName, lastName, userTicket, email)
	fmt.Printf("%v tickets left now\n", remainingTickets)
}

func getFirstNames() []string{
	var firstNames []string
	for _, booking := range bookings{
		firstNames = append(firstNames, booking.firstName)
		// range provides the index and value of that element 
		// range iterate over elemnt for different data structure
		// blank identifier '_' for unused variable 
		//fmt.Printf("These are all our bookings:%v\n", booking)
	
		// split the string with white space as separator 
		// and return a slice with split elements
	}		
	return firstNames
}

// generate a ticket and send it to the email
func sendTicket(userTicket uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v ticket for %v %v", userTicket, firstName, lastName)
	fmt.Println("###################################################")
	fmt.Printf("Sending %v to the email address %v\n", ticket, email)
	fmt.Println("####################################################")
}



