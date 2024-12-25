package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

var conferenceName = "Go Conference"

const conferenceTickets = 50

var remainingTickets = 50
var bookings []string

func main() {
	http.HandleFunc("/", renderForm)
	http.HandleFunc("/book", handleBooking)
	http.HandleFunc("/list", listBookings)
	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func renderForm(w http.ResponseWriter, r *http.Request) {
	page := `
	<!DOCTYPE html>
	<html>
	<head>
		<title>Go Conference Booking</title>
		<style>
			body {
				font-family: Arial, sans-serif;
				background-color: #f4f4f9;
				margin: 0;
				padding: 0;
				text-align: center;
				color: #333;
			}
			h1 {
				color: #444;
			}
			form {
				background: #fff;
				border: 1px solid #ccc;
				padding: 50px;
				margin: 20px auto;
				width: 300px;
				border-radius: 10px;
			}
			input, button {
				width: 100%;
				margin-bottom: 10px;
				padding: 8px;
				border: 1px solid #ccc;
				border-radius: 5px;
			}
			button {
				background: #5cb85c;
				color: white;
				border: none;
				cursor: pointer;
			}
			button:hover {
				background: #4cae4c;
			}
		</style>
	</head>
	<body>
		<h1>Welcome to {{.ConferenceName}} booking application</h1>
		<p>We have a total of {{.TotalTickets}} tickets, and {{.RemainingTickets}} are still available.</p>
		<a href="/list">View Bookings</a><br><br>
		<form method="POST" action="/book">
			<label>First Name:</label>
			<input type="text" name="firstName" required><br>
			<label>Last Name:</label>
			<input type="text" name="lastName" required><br>
			<label>Email:</label>
			<input type="email" name="email" required><br>
			<label>Number of Tickets:</label>
			<input type="number" name="tickets" required><br>
			<button type="submit">Book Tickets</button>
		</form>
	</body>
	</html>`

	tmpl, _ := template.New("form").Parse(page)
	tmpl.Execute(w, map[string]interface{}{
		"ConferenceName":   conferenceName,
		"TotalTickets":     conferenceTickets,
		"RemainingTickets": remainingTickets,
	})
}

func handleBooking(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		firstName := r.FormValue("firstName")
		lastName := r.FormValue("lastName")
		email := r.FormValue("email")
		tickets, err := strconv.Atoi(r.FormValue("tickets"))

		// Validation logic
		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicket := err == nil && tickets > 0 && tickets <= remainingTickets

		if isValidName && isValidEmail && isValidTicket {
			remainingTickets -= tickets
			bookings = append(bookings, firstName+" "+lastName)

			message := fmt.Sprintf(`
			<!DOCTYPE html>
			<html>
			<head>
				<title>Booking Confirmation</title>
			</head>
			<body>
				<h1>Thank you %s %s for booking %d tickets!</h1>
				<p>You will receive a confirmation email at %s.</p>
				<p>%d tickets remaining for %s.</p>
				<a href="/">Go back to home</a>
			</body>
			</html>`, firstName, lastName, tickets, email, remainingTickets, conferenceName)

			w.Write([]byte(message))

			// Check if tickets are sold out
			if remainingTickets == 0 {
				fmt.Println("All tickets are booked!")
			}
		} else {
			errorMessage := "<h1>Invalid Input:</h1><ul>"
			if !isValidName {
				errorMessage += "<li>First name or last name must have at least 2 characters.</li>"
			}
			if !isValidEmail {
				errorMessage += "<li>Email must contain '@'.</li>"
			}
			if !isValidTicket {
				errorMessage += fmt.Sprintf("<li>Invalid ticket count. Only %d tickets are remaining.</li>", remainingTickets)
			}
			errorMessage += "</ul><a href='/'>Go back to try again</a>"

			w.Write([]byte(errorMessage))
		}
	}
}

func listBookings(w http.ResponseWriter, r *http.Request) {
	page := "<!DOCTYPE html><html><head><title>Booking List</title></head><body><h1>Booking List</h1><ul>"
	for _, booking := range bookings {
		page += fmt.Sprintf("<li>%s</li>", booking)
	}
	page += "</ul><a href='/'>Go back to home</a></body></html>"

	w.Write([]byte(page))
}
