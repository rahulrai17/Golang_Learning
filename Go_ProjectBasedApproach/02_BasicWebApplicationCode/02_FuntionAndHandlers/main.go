package main

import (
	"fmt"
	"net/http"
)

var portNumber = ":8080" // The port number the server will listen on

// Home is the handler function for the root URL ("/")
func Home(w http.ResponseWriter, r *http.Request) {

	// Respond with "This is the home page"
	fmt.Fprintf(w, "This is the home page")
}

// About is the handler function for the "/about" URL
func About(w http.ResponseWriter, r *http.Request) {

	// Call addValues function to add 2 and 2
	sum := addValues(2, 2)

	// Respond with a message that includes the sum (2 + 2 = 4)
	fmt.Fprintf(w, "This is the about page and 2 + 2 is %d", sum)
}

// addValues is a helper function that adds two integers
func addValues(x, y int) int {

	return x + y // Returns the sum of x and y
}

// The main function where the server is set up
func main() {

	// Register the Home function to handle requests at "/"
	http.HandleFunc("/", Home)

	// Register the About function to handle requests at "/about"
	http.HandleFunc("/about", About)

	// Start the server and listen on port 8080
	http.ListenAndServe(portNumber, nil)
}


// Summary
// This Go program sets up a simple HTTP server that listens on port 8080. It has two route handlers:

// Home handler (at "/"): responds with the text "This is the home page".
// About handler (at "/about"): calls addValues (which adds 2 + 2), then responds with "This is the about page and 2 + 2 is 4".

// The main function registers these handlers and starts the server. When a user visits http://localhost:8080/, they'll see the home page, and http://localhost:8080/about will show the about page with the calculated result.