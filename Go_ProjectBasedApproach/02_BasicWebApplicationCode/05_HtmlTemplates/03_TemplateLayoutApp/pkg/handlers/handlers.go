package handlers

// Since both the file are in package main we can use the function or varible from one file to other

import (
	"03_TemplateLayoutApp/pkg/render"
	"net/http"
)

// "H" in home is capital so that it can be accessed from other packages also
func Home(w http.ResponseWriter, r *http.Request){
	render.RenderTemplate(w, "home.page.tmpl")
}

// "w" send replies to the user of webpage , "r" keeps the request values from the user.
func About(w http.ResponseWriter, r *http.Request){
	render.RenderTemplate(w, "about.page.tmpl")
}

// This is endpoint defined to check check second way of rendering using renderMethodTwo/RenderTemplateAdv
func HomeTwo(w http.ResponseWriter, r *http.Request){
	render.RenderTemplateAdv(w, "homeTwo.page.tmpl")
}

// This is endpoint defined to check check second way of rendering using renderMethodTwo/RenderTemplateAdv
func HomeThree(w http.ResponseWriter, r *http.Request){
	render.RenderTemplateComplex(w, "homeThree.page.tmpl")
}


