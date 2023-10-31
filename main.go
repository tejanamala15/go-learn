package main

import "fmt"

func main(){
	conferenceName := "Go Conference"
	const conferenceTickets int= 50
	var remainingTickets uint = 50

	// fmt.Println(&conferenceName) -> Pointer

	fmt.Printf("conferenceName is %T, conferenceTickets is %T, remainingTickets is %T.\n", conferenceName, conferenceTickets, remainingTickets)

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")


    //array
	// var bookings [50] 
	
	//slice
	var bookings [] string
	

	var firstName string
	var lastName string
	var email string
	var userTickets int
	fmt.Println("Please enter your first name:")
	fmt.Scan(&firstName)

	fmt.Println("Please enter your last name:")
	fmt.Scan(&lastName)

	fmt.Println("Please enter your email address:")
	fmt.Scan(&email)

	fmt.Println("Please enter the number of tickers you need:")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - uint(userTickets)
	// bookings[0] = firstName + " " + lastName
	bookings = append(bookings, firstName + " " + lastName)

	fmt.Printf("The whole slice: %v\n", bookings)
	fmt.Printf("The first value: %v\n", bookings[0])
	fmt.Printf("Slice length: %v\n", len(bookings))



	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

	fmt.Printf("These are all our bookings %v\n", bookings)


}