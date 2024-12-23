# Template rendering application flow(Part - 1) :

## config.go

1. Step one is to define the `AppConfig` Sturct which will have all the `application related configurations` stored here.
2. Here we have declare `AppConfig` which has a `TemplateCache` of type map inside it, which we will use to load all the templates in the starting of the application only.

```go
package config

import "html/template"

// AppConfig holds the application config
type AppConfig struct{
  TemplateCache map[string]*template.Template
}
```

## main.go

1. `main.go` is the starting point of any application.
2. Here we first create the variable `app` of type `AppConfig` from `config.go`.
3. we call the `CreateTemplateCache()` function from `render.go`, which will then create and return the `rendered template` here.
4. Next step is to Store the `rendered template` in the `TemplateCache` varible of `AppConfig` for the use of whole application.
5. Next we will pass the refrence of `app` of type `AppConfig` to the `render.NewTemplates()` function.

```go
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
	// This will create a variable app of type Appconfig from config.go file
	var app config.AppConfig

	// This will help in creation of template in the starting of the application and store it in the tc variable
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}

	// This will store the value in the TemplateCache variable
	app.TemplateCache = tc

	// This will pass reference to the AppConfig struct
	render.NewTemplates(&app)

	http.HandleFunc("/home", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Starting application on port %s", portNumber)
	http.ListenAndServe(portNumber, nil)
}
```

## render.go

1. Here we are creating `app` again but, here it is a pointer which means it will hold the memory of some AppConfig variable already created in application.
2. When `NewTemplates` function is called it gets a `refrence(which means address to some variable's memory)` as a value to its parameter.
3. The `refrence` here we have passed is of `app` from `main.go`. So now this enables to do what we please with the actual `app` varible and its value stored in memory.
4. Why using pointer, without pointer we be creating new varibles in the memory and making any changes in that won't reflect in the acutal memory.
5. Now here we assing the passed `value(address)` to the `app(refrence variable)` here.
6. This process allowed us to `create Template cache` when the application starts and use it till the end of the application.
7. Now the flow will go something like this when application starts it will create template cache, then when a enpoint `eg: \home` is hit it will `handler method`, which will then call `RenderTemplate` from `render.go`.
8. Here we will get TemplateCache using the `app(refrence variable)`.
9. Search the template called by the endpoint `eg: \home` will call something like `home.page.tmpl or home.html`.
10. Then there we will hold data in buffer after executing it(optional step).
11. Then the webpage will be shown.

```go
package render

import (
	"04_GoodPracticeApp/pkg/config"
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

func RenderTemplate(w http.ResponseWriter, tmpl string){
	// Get the template cache from the config.go
	tc := app.TemplateCache

	// get requested template from the cache
	t, ok := tc[tmpl]
	if !ok{
		log.Fatal("Could not get template from template cache")
	}

	// bytes.Buffer: A buffer to hold data temporarily in memory before writing it out.
	buf := new(bytes.Buffer)

	// t.Execute: Executes a template, injecting optional data.
	err := t.Execute(buf, nil)
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
```

## handler.go

1. In the handler.go we handle the requests for API endpoints.

```go
package handlers

// Since both the file are in package main we can use the function or varible from one file to other

import (
	"02_TemplatesApp/pkg/render"
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

```

# Make it more Efficient(Part - 2 )

Till now we have covered the best way of handling template rendering(i.e, `Template caching`). Now lets cover Template caching and rendering in `different modes`, such as `development`, `staging`, and `production`, is a strategy used in web development to optimize performance and streamline the development process. Here's an explanation of the concept behind it:

### The Role of Templates in Web Development

Templates are pre-defined structures (HTML, XML, or other markup formats) used by web applications to dynamically generate views based on data. When rendering a template:

    - The server combines the template with data to produce a final output (e.g., an HTML page).
    - This process can be computationally expensive if done repeatedly for similar requests.

To optimise this process we use template caching. Template caching refers to storing a pre-processed or pre-compiled version of a template in memory or storage to avoid repetitive processing:

    - Without caching: Each request triggers the parsing, processing, and rendering of the template.
    - With caching: The server skips repetitive processing and retrieves the ready-to-use template.
    - Caching improves : Performance: Reduces response times by serving templates faster. Efficiency: Lowers CPU and memory usage by avoiding redundant processing.

### Different Modes: Development vs. Production

The way templates are handled differs between development and production modes to balance performance with flexibility.

1. Development Mode:

   - `Goal`: Enable rapid changes and debugging.
   - `Features`:
     - `No caching or minimal caching`: Templates are recompiled on every request to reflect the latest changes immediately.
     - `Debug-friendly rendering`: Provides verbose error messages, stack traces, and unoptimized outputs to help developers identify issues.
     - `Frequent reloads`: Ensures the development server picks up file changes automatically.

2. Production Mode

   - `Goal`: Maximize efficiency and performance for end users.
   - `Features`:
     - `Aggressive caching`: Templates are pre-compiled and stored in memory or on disk. Changes require a manual reload or restart.
     - `Minimized assets`: HTML, CSS, and JavaScript in templates are minified to reduce load times.
     - `Optimized rendering`: Templates are rendered using faster, optimized pathways with minimal debugging overhead.

### To Acheive all we can simply update our code:

## Config.go

1. Lets Start with `config.go`
2. Here up have added a `UseCache` we will simply use depending on the use case for example if we need to use the template cache or relaod the template with every change.

```go
package config

import "html/template"

// This is a file that is accessed by whole project. This is also know as (application wide configuration).
// We only add most useful and valid configuration here.
// This file is imported by whole application
// Remenber this should be program in a way that this doesn't import anything from the application.
// It is only allowed to import form standard libraries only (This precaution is taken to avoid the import cycle. Import cycle doesn't allow the application to be compiled).

// AppConfig holds the application config
type AppConfig struct{
  UseCache bool
  TemplateCache map[string]*template.Template
}
```

## main.go

1. Now here we have used the `UseCache` variable and set its value to true, the values here will mean
   - true : We will use the saved templates.
   - false : We will reload template with every new request(i.e, it will call the `CreateTemplateCache()` on every API request).
2. Here we have also called a `NewRepo()`, `NewHandlers()` function, This will be used in passing the AppConfig values to the `handler.go` code.

```go
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

	// will store the value in the TemplateCache variable
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
```

## handlers.go

1. Now here we can see that the line of code has increased.
2. What have implemented the repository pattern here.

   - The `Repository pattern` implemented here provides a structured and centralized way to manage dependencies and configurations used by your application. Let’s break down its purpose, functionality, and benefits

     - The `Repository` struct serves as a centralized holder for shared resources or dependencies (like the AppConfig) that are required by the handlers (e.g., Home and About). By creating a repository, you decouple the application's handlers from the rest of your application logic, making the code cleaner, more modular, and easier to maintain.

     ### Why Do Handlers Need a Repository?

     - `Access Shared Data or Settings`: Your app likely has global configurations or resources that multiple handlers (like Home and About) need. For example:
       - A cache of templates.
       - Session management.
       - Feature toggles (like UseCache).
     - Instead of passing these settings or resources manually to every handler, the repository stores them centrally. Handlers can then grab what they need without extra code.
     - Without using repository this code will look something like:

     ```go
     // we will need to pass appConfig (and maybe other resources) to every handler.
     // If you add more shared resources, you’d have to update every handler function signature.
     // But using repository if we need to add new values or any other resources we can simply update it in the Repository stuct and the handler will not need to worry about it anymore
     func Home(w http.ResponseWriter, r *http.Request, appConfig *config.AppConfig) {
     		render.RenderTemplate(w, "home.page.tmpl", appConfig.TemplateCache)
     }

     ```

```go
package handlers

// Since both the file are in package main we can use the function or varible from one file to other

import (
	"04_GoodPracticeApp/pkg/config"
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
	render.RenderTemplate(w, "home.page.tmpl")
}

// "w" send replies to the user of webpage , "r" keeps the request values from the user.
func (m *Repository) About(w http.ResponseWriter, r *http.Request){
	render.RenderTemplate(w, "about.page.tmpl")
}

```

## render.go

1. Last we have `render.go`. Where we have the logic updated for using the `Template Cache`

```go
package render

import (
	"04_GoodPracticeApp/pkg/config"
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

func RenderTemplate(w http.ResponseWriter, tmpl string){
	var tc map[string]*template.Template

	if app.UseCache{
		// Get the template cache from the config.go
		println("I am using saved template Cache")
		tc = app.TemplateCache
	}else {
		println("I am Creating template with every request")
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
	err := t.Execute(buf, nil)
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
```
