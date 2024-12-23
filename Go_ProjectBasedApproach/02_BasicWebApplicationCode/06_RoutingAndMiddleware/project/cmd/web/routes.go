package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rahulrai17/porject/pkg/handlers"
)

func routes() http.Handler {
	// Create a new router
	mux := chi.NewRouter()

	// Use the Recoverer middleware to recover from panics
	mux.Use(middleware.Recoverer)
	// Use a custom middleware to write to console, note: both these middlewares will be automatically applied in all routes since we are using routers here(eg: chi in this case).
	mux.Use(writeToConsole)
	// This will be called twice once before the api request and once after.
	
	// Define the route for the home page
	mux.Get("/home", http.HandlerFunc(handlers.Repo.Home))
	// Define the route for the about page
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	// Return the configured router
	return mux
}