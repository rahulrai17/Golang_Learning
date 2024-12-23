package render

import (
	"fmt"
	"html/template"
	"net/http"
)

// this is function created so that we will not need do the render template thing multiple times
func RenderTemplate(w http.ResponseWriter, tmpl string){
	// remember "../../templates/" --> this is the way to directing back in folder in golang
	parsedTemplate, err := template.ParseFiles("../../templates/" + tmpl, "../../templates/base.layout.tmpl" ) 
  if err != nil {
		fmt.Printf("error parsing template: %s,\nerror is: %s\n", tmpl, err)
		return
	}

	//since err is first declared above we can use it again by updating its value here
	err = parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Printf("error executing template: %s, \nerror is: %s\n", tmpl, err)
		return
	}
}