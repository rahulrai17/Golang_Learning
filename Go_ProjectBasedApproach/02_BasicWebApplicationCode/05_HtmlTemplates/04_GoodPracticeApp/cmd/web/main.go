package main

import (
	"04_GoodPracticeApp/pkg/config"
	"04_GoodPracticeApp/pkg/handlers"
	"04_GoodPracticeApp/pkg/render"
	"fmt"
	"log"
	"net/http"
)

const portNumber = ":8080" 

func main() {
	//this will create a variable app of type Appconfig from config.go file 
	var app config.AppConfig

	//This will help in creation of template in the starting of the application and store it in the tc variable 
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}

	// will will store the value in the TemplateCache variable  
	app.TemplateCache = tc
	app.UseCache = true

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	// this will pass reference to the AppConfig struct
	render.NewTemplates(&app)
	


	http.HandleFunc("/home", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)


	fmt.Printf("Starting application on port %s", portNumber)
	http.ListenAndServe(portNumber, nil)
}