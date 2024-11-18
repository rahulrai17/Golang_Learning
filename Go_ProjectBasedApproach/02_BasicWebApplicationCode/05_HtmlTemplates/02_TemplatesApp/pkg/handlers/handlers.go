package handlers

// Since both the file are in package main we can use the function or varible from one file to other

import (
	"02_TemplatesApp/pkg/render"
	"net/http"
)

// "H" in home is capital so that it can be accessed from other packages also
func Home(w http.ResponseWriter, r *http.Request){
	render.RenderTemplate(w, "home.html")
}

// "w" send replies to the user of webpage , "r" keeps the request values from the user.
func About(w http.ResponseWriter, r *http.Request){
	render.RenderTemplate(w, "about.page.tmpl")
}


