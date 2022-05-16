package main

import (
	"bookingApp/helper"
	"fmt"
	"sync"
	"time"
)

const conferenceTickets = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings = make([]userData, 0)

type userData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUsers()

	firstName, lastName, email, userTickets := getUserInput()

	bookTicket(firstName, lastName, email, userTickets)
	wg.Add(1)
	go sendTicket(firstName, lastName, email, userTickets)
	wg.Wait()
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v tickets are still available!\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)

	}
	return firstNames
}
func bookTicket(firstName string, lastName string, email string, userTickets uint) {
	isValidEmail, isValidName, isValidTickets := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)
	if isValidEmail && isValidName && isValidTickets {
		remainingTickets -= userTickets

		var userData = userData{
			firstName:       firstName,
			lastName:        lastName,
			email:           email,
			numberOfTickets: userTickets,
		}
		bookings = append(bookings, userData)
		fmt.Printf("List of bookings is %v\n", bookings)
		fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve a confirmation email at %v.\n", firstName, lastName, userTickets, email)
		fmt.Printf("Remaining tickets now are %v\n", remainingTickets)
		firstNames := getFirstNames()
		fmt.Printf("The first names of bookings are: %v\n", firstNames)
	} else {
		helper.Validation(isValidEmail, isValidName, isValidTickets, remainingTickets)
	}
}
func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	fmt.Println("Let's book your tickets!")
	fmt.Println("Please enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Please enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Please enter your email: ")
	fmt.Scan(&email)

	fmt.Println("Please enter number of tickets: ")
	fmt.Scan(&userTickets)
	return firstName, lastName, email, userTickets
}
func sendTicket(firstName string, lastName string, email string, userTickets uint) {
	time.Sleep(50 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v\n", userTickets, firstName, lastName)
	fmt.Println("####################")
	fmt.Printf("Sending ticket: \n %v to email address:  %v\n", ticket, email)
	fmt.Println("####################")
	wg.Done()
}
