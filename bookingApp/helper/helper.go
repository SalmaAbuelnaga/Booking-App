package helper

import (
	"fmt"
	"strings"
)

func ValidateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTickets := userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTickets
}
func Validation(isValidEmail bool, isValidName bool, isValidTickets bool, remainingTickets uint) {
	if !isValidEmail {
		fmt.Println("Please enter a valid email")
	}
	if !isValidName {
		fmt.Println("Names have to be 2 or more characters")
	}
	if !isValidTickets {
		fmt.Printf("Tickets need to be between 0 and %v tickets\n", remainingTickets)
	}
}
