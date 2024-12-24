package main
import "fmt"
func main(){
	var conferenceName ="Go Conference"
	const conferenceTickets = 50
	var remainingTickets = 50
	var userName string
	var lastName string
	var email string
	var userTickets int
	println("Welcome to", conferenceName, "booking application")
	println("we have total of", conferenceTickets,"tickets and", remainingTickets,"are still available.")
	println("Get your tickets here to attend")
	print("Enter your first name: ") 
	fmt.Scan(&userName)
	print("Enter your last name: ") 
	fmt.Scan(&lastName)
	print("Enter your email: ") 
	fmt.Scan(&email)
	print("Enter the no.of tickets: ") 
	fmt.Scan(&userTickets)
	fmt.Println("thank you",userName, lastName, "for booking", userTickets, "tickets. You will recive a confermation mail at", email)
	
	
	
}
