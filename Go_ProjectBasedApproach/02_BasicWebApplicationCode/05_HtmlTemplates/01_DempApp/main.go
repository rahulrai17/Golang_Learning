package main

import (
	"fmt"
	"net/http"
	"text/template"
)

//defining 8080 as a string because for ListenAndServe first argument is string in will be uncompatible
const portNumber = ":8080"

func main() {
	  // Api's declared with function mapping 
		http.HandleFunc("/home", Home)
		http.HandleFunc("/about",About)
		
		fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
		http.ListenAndServe(portNumber , nil)
}

// "H" in home is capital so that it can be accessed from other packages also
func Home(w http.ResponseWriter, r *http.Request){
	renderTemplate(w, "home.html")
}

// "w" send replies to the user of webpage , "r" keeps the request values from the user.
func About(w http.ResponseWriter, r *http.Request){
	renderTemplate(w, "about.page.tmpl")
}


// this is function created so that we will not need do teh render template thing multiple times 
func renderTemplate(w http.ResponseWriter, tmpl string){
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl) 
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template:", err)
	}
}