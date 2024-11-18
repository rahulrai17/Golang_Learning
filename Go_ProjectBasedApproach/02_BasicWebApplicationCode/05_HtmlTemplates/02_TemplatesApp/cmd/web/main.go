package main

import (
	"02_TemplatesApp/pkg/handlers"
	"fmt"
	"net/http"
)

const portNumber = ":8080" 

func main() {
	http.HandleFunc("/home", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Starting application on port %s", portNumber)
	http.ListenAndServe(portNumber, nil)
}