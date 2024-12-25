package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = 50
	var userName string
	var lastName string
	var email string
	var userTickets int
	bookings := []string{}

	fmt.Println("Welcome to", conferenceName, "booking application")
	fmt.Println("We have a total of", conferenceTickets, "tickets and", remainingTickets, "are still available.")
	fmt.Println("Get your tickets here to attend")

	for {
		fmt.Print("Enter your first name: ")
		fmt.Scan(&userName)
		fmt.Print("Enter your last name: ")
		fmt.Scan(&lastName)
		fmt.Print("Enter your email: ")
		fmt.Scan(&email)
		fmt.Print("Enter the number of tickets: ")
		fmt.Scan(&userTickets)

		if userTickets < remainingTickets {
			remainingTickets -= userTickets
			bookings = append(bookings, userName+" "+lastName)

			fmt.Println("Thank you", userName, lastName, "for booking", userTickets, "tickets. You will receive a confirmation mail at", email)
			fmt.Println(remainingTickets, "tickets remaining for", conferenceName)

			// Display all first names of the bookings
			userNames := []string{}
			for _, booking := range bookings {
				names := strings.Fields(booking)
				userNames = append(userNames, names[0])
			}
			fmt.Println("These are all our bookings:", userNames)

			// Check if tickets are sold out
			if remainingTickets == 0 {
				fmt.Println("Our conference is fully booked. Come back next year!")
				break
			}
		} else {
			fmt.Printf("We only have %v tickets remaining, so you can't book %v tickets.\n", remainingTickets, userTickets)
			continue
		}
	}
}
