package main

import (
	"fmt"
	"net/http"
)

//this is a custom middle ware function i am creating, also its pretty common to name your param as next for middleware.
func writeToConsole(next http.Handler) http.Handler {
	// Return an http.HandlerFunc which is a function that implements the http.Handler interface
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Print a message to the console
			fmt.Println("From writeToConsole()-middleware : i will take you to the page")
			// Call the next handler in the chain, passing along the ResponseWriter and Request
			next.ServeHTTP(w, r)
	})
}