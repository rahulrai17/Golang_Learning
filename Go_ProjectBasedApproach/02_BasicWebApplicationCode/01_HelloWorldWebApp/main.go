package main

import (
	"fmt"      // Package for formatted I/O operations (like Println, Sprintf, Fprintf)
	"net/http" // Package for creating HTTP servers and handling HTTP requests and responses
)

func main() {

	// Define a route for the root URL path ("/")
	// HandleFunc registers a handler function for incoming HTTP requests to a specific path ("/" in this case).
	// When a request is received at this path, the anonymous function defined here will be called.
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		
		// Fprintf writes formatted data to the provided writer `w`, which is an http.ResponseWriter.
		// This sends the "Hello world!" message to the client (typically a browser).
		n, err := fmt.Fprintf(w, "Hello world!")
		
		// Check if there was an error when trying to write to the response writer
		if err != nil {
			// Print the error to the server's standard output if something went wrong
			fmt.Println(err)
		}

		// Print the number of bytes written to the response to the server's standard output
		// Sprintf formats the string, and Println outputs it.
		fmt.Println(fmt.Sprintf("bytes written: %d", n))
	})

	// ListenAndServe starts an HTTP server on port 8080 and listens for incoming requests.
	// It uses the default handler, which in this case is the function attached to "/".
	// If there is an error starting the server, it will be returned, though not handled here.
	http.ListenAndServe(":8080", nil)
}
