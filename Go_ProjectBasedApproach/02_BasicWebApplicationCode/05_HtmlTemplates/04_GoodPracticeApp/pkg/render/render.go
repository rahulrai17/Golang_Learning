package render

import (
	"04_GoodPracticeApp/pkg/config"
	"04_GoodPracticeApp/pkg/models"
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

// app is a global variable which is of type *config.AppConfig (pointer since we need to use the value)
var app *config.AppConfig

// NewTemplates sets the config for the template package
func NewTemplates(a *config.AppConfig){
	app = a
}

func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData){
	var tc map[string]*template.Template

	if app.UseCache{
		// Get the template cache from the config.go
		println("Using Saved Cache")
		tc = app.TemplateCache
	}else {
		println("Creating New Template everytime")
		tc, _ = CreateTemplateCache()
	}	

	// get requested template from the cache
	t, ok := tc[tmpl]
	if !ok{
		log.Fatal("Could not get template from template cache")
	}

	// bytes.Buffer: A buffer to hold data temporarily in memory before writing it out.
	buf := new(bytes.Buffer)

	// t.Execute: Executes a template, injecting optional data.
	err := t.Execute(buf, td)
	if err != nil {
		log.Println(err)
	}

	// render template
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

// This is a function to Create Template cache that returns a value that is map which has key : template_name and value : rendered template and a error
func CreateTemplateCache() (map[string]*template.Template, error){

	// myCache := make(map[string]*template.Template) //creating map using make keyword
	myCache := map[string]*template.Template{} //this is creating and empty map without make, keyword both are same

	// get all of the files name *.page.tmpl from ./templates. 
	// filepath.Glob: Returns a list of files matching a glob pattern.
	pages, err := filepath.Glob("../../templates/*.page.tmpl")
	if err != nil{
		return myCache, err
	}

	// range through all files ending with *.page.tmpl
	for _, page := range pages{

		// filepath.Base: Extracts the base name of a file path.
		name := filepath.Base(page)

		// template.New: Creates a new template instance with a specific name.
		// ParseFiles: Parses one or more template files into a template instance.
		ts , err := template.New(name).ParseFiles(page)
		if err != nil{
			return myCache, err
		}

		matches, err := filepath.Glob("../../templates/*.layout.tmpl")
		if err != nil{
			return myCache, err
		}

		if len(matches) > 0 {
			//ParseGlob: Parses all template files matching a glob pattern into the template instance.
			ts, err = ts.ParseGlob("../../templates/*.layout.tmpl")
			if err != nil{
				return myCache, err
			}
		}

		myCache[name] = ts
	}

	return myCache, nil

}