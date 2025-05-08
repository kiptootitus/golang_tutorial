package main
import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50

	fmt.Printf("conferenceTickers %T and remainingTickers is %T types.\n", conferenceName, remainingTickets)


	fmt.Printf("We have these %v tickers\n", remainingTickets)
	fmt.Println("Welcome to our conference booking",conferenceName ,"Application")
	fmt.Println("We have a total of these",conferenceTickets,"tickets and ",remainingTickets,"are available")

	// Data types

	var userName string
	var noTickets uint
	var firstName string
	var lastName string
	var emailAddr string
	var phoneNumber int

	
	//ask user to enter the name 
	fmt.Print("Please Enter your Username: ")
	fmt.Scan(&userName)

	fmt.Print("Please Enter your First Name: ")
	fmt.Scan(&firstName)

	fmt.Print("Please Enter your Last Name: ")
	fmt.Scan(&lastName)

	fmt.Print("Please Enter your Email Adress: ")
	fmt.Scan(&emailAddr)

	fmt.Print("Please Enter your Phone Number: ")
	fmt.Scan(&phoneNumber)


	// to allow user input we must use a pointer & to point to a memory address of the variable representing the value
	fmt.Print("Enter number of tickers you want to book: ")
	fmt.Scan(&noTickets)

	remainingTickets = remainingTickets - noTickets

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive confirmation in %v and SMS notification in %v. Thank you %v for your time.\n", firstName, lastName, noTickets, emailAddr, phoneNumber, firstName)
	
	// lets print the remaining tickets after booking

	fmt.Printf("The remaining tickets are %v for %v\n", remainingTickets, conferenceName)

}
