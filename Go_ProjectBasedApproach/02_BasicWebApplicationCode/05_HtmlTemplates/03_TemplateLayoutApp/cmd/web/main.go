package main

import (
	"03_TemplateLayoutApp/pkg/handlers"
	"fmt"
	"net/http"
)

const portNumber = ":8080" 

func main() {
	http.HandleFunc("/home", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	http.HandleFunc("/hometwo", handlers.HomeTwo)
	http.HandleFunc("/homethree", handlers.HomeThree)

	fmt.Printf("Starting application on port %s", portNumber)
	http.ListenAndServe(portNumber, nil)
}