package render

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

// This is a variable that we will use to store rendered Template.This will a package level variable since it will be called from other functions too.
// We are using "*template.Template" because this is the value returned by template.ParseFiles function
var tempCache = make(map[string]*template.Template)

func RenderTemplateAdv(w http.ResponseWriter, t string) {
	var tmpl *template.Template
	var err error

	// check to see if we already have the template in our cache
	// map returns two values (value_of_key , if_key_available_in_map), so here we don't need the first value therefore we are ignoring it and using the second value which will be true or false .
	_, inMap := tempCache[t]
	if !inMap {
		// need to create the template
		log.Println("Creating template and adding to cache...")
		err = createTemplateCache(t)
		if err != nil{
			log.Println(err)
		}
	} else {
		// we have the template in the cache
		log.Println("Using cached template...")
	}

	// get the template and store it in tmpl
	tmpl = tempCache[t]

	// execute the template
	err = tmpl.Execute(w, nil)
	if err != nil{
		log.Println(err)
	}

}

func createTemplateCache(t string) error {

	// Creating a string with the templates locations to render
	templates := []string{
		fmt.Sprintf("../../templates/" + t), "../../templates/base.layout.tmpl",
	}

	// parse the template
	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}

	// add template to cache (map)
	tempCache[t] = tmpl

	return nil
}                                                                                                                        


// Summary (Template Caching System)
// This is a basic advance method to render html templates you can use it for beginner level applications.
// This method renders a template the it stores it in map then it uses it till the application ends.
// Drawback : you need to specify the dependencies for exmaple to render a page we neeed the base.layout.tmpl so we added it, therefore as more dependency increases it will be an headche
// This is good method , but there a more complex method which is renderMethodThree do check it if you need to more good render mehtod.