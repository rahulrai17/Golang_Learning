package handlers

// Since both the file are in package main we can use the function or varible from one file to other

import (
	"04_GoodPracticeApp/pkg/config"
	"04_GoodPracticeApp/pkg/models"
	"04_GoodPracticeApp/pkg/render"
	"net/http"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type 
type Repository struct{
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository{
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// "H" in home is capital so that it can be accessed from other packages also
func (m *Repository) Home(w http.ResponseWriter, r *http.Request){
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// "w" send replies to the user of webpage , "r" keeps the request values from the user.
func (m *Repository) About(w http.ResponseWriter, r *http.Request){
	// creating a map with data
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again"

	// passing the map with data by matching the fields
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap, 
	})
}


