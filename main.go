package main

import (
	"fmt"
	"strconv"

	"example.com/booking/app/helper"
)

const conferenceTickets int = 50
var conferenceName  = "NezzyGroup Conference"
var remainingTickets uint = 50
var bookings = make([]map[string]string, 0)



func main(){
	//fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName )
	greetUsers()

	
	for {
	
	firstName, lastName, email, userTickets :=getUserInput()

	isValidName, isValidEmail, isValidTicketNumber :=helper.	ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets )
	

	if isValidName && isValidEmail && isValidTicketNumber{
		
		bookTicket(userTickets, firstName, lastName, email,)
		// call function print first names
		firstNames := printFirstName()
		fmt.Printf("The first names of bookings are: %v\n", firstNames)

	 if remainingTickets == 0 {
		//end the program
		fmt.Println("Our conference is booked out come back next year.")
		break
	 }
	
	}  else{
		if !isValidName{
			fmt.Println("first name or last name you enter is too short")
		}
		if  !isValidEmail{
			fmt.Println("email address you entered doesn't contain @sign")
		}
		if !isValidTicketNumber{
			fmt.Println("Number of tickets you entered is invalid ")}
	}
	}
}


func greetUsers() {
	fmt.Printf( "Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	println("Get your ticket here to attend")

}

func printFirstName() []string {

	firstNames := []string{}
	for _, booking := range bookings{
		firstNames = append(firstNames, booking["firstName"])
	}
	
	return firstNames
}


func getUserInput()(string, string, string, uint){
		
	var firstName string
	var lastName string
	var email string
	var userTickets uint	
	//ask user for their name
	fmt.Printf("Enter your first name:")
	fmt.Scan( &firstName)
	
	fmt.Printf("Enter your last name:")
	fmt.Scan( &lastName)

	fmt.Printf("Enter your email address:")
	fmt.Scan( &email)

	fmt.Printf("Enter the number of ticket:")
	fmt.Scan( &userTickets)

	return firstName, lastName, email, userTickets

}
func bookTicket( userTickets uint,  firstName string, lastName string, email string, ){
	remainingTickets = remainingTickets - userTickets

	//Create a map for the user
	var userData = make(map[string]string)
	userData["firstName"] = firstName
	userData["lastName"] = lastName
	userData["email"] = email
	userData["numberofTicket"] = strconv.FormatUint(uint64(userTickets), 10)

	bookings = append(bookings, userData)
	fmt.Printf("list of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	 fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName )
}
