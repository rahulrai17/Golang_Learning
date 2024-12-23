package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rahulrai17/porject/pkg/config"
	"github.com/rahulrai17/porject/pkg/handlers"
	"github.com/rahulrai17/porject/pkg/render"
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

	// This is a better way to define start the server.
	srv := &http.Server{
		Addr: portNumber,
		Handler: routes(),
	}
	
	fmt.Printf("Starting server at %s\n", portNumber)
	
	//http.ListenAndServe(portNumber, routes())
  err = srv.ListenAndServe();
	log.Fatal(err)
}