# BookingApp

Readable code 
- Greet the user
- Get the usetr input
- Validate the user input
- If all entry is valid 
- Book the ticket
- Get the first names
- Print the fist names
- If no tickets left, exit the program

Creating a slice with make
- alternative way to create a slice, define the initial size of the slice
 
Define a structure that can hold mixed data type
- which fields of the user type
- struct is a light weight class which e.g. doesn't support inheritance

Scope of variable
- Local 
* Declare within function like firstNames, can be used only within that function
* Delcare within if-else block like names, can be used only within that block

- Package
* Declaration outside all functions, can be used everywhere in the same function like conferenceName can be used within main package
 
- Global
* Declaration outside all the functions & uppercase letters, like Myvar used in validation.go ; can be used across all packages 

time
- sleep function stops or blocks the current thread (go routine) execution for the defined duration. 

go keyword
- go routine is light weight thread managed by go runtime 

Problem we have to fix:
main go routine exited before "send ticket" had time to start and execute the code
By default, main go routine does not wait for other goroutine

Solution :
use wait group that wait for launch goroutine to finish 
- Add() : set the number of goroutine to wait for (use before the go routine)
- wait() : block until the counter reaches 0 (use at last line of main.go code)
- done(): decrement the waitgroup counter by 1 , indicate the it is finished (used inside the goroutine function)

