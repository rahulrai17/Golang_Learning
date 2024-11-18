## Serve Dynamic HTML Using Templates

Serving dynamic HTML using templates in Go is a powerful way to create web pages with content that can change based on the data you pass to it. Go’s standard library provides the `html/template` package to help you render HTML templates with data injected into them. Below, I’ll walk you through a detailed explanation, including multiple examples to illustrate key concepts.

### How Templates Work

The `html/template` package uses **template files** that contain HTML markup along with placeholders, called **actions** (e.g., `{{ .Title }}`, `{{ .Content }}`). These placeholders are replaced with actual data when the template is rendered.

1. **Actions** are marked by `{{ ... }}` syntax.
2. **Data** is passed to the template as a struct, map, or another type containing fields to replace the placeholders.

### Basic Setup for Dynamic HTML Templates

#### Step 1: Create a Template File

Let’s start with a simple HTML template. Create a file named `index.html` in a `templates` directory:

```html
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>{{ .Title }}</title>
  </head>
  <body>
    <h1>{{ .Heading }}</h1>
    <p>{{ .Content }}</p>
  </body>
</html>
```

Here, `{{ .Title }}`, `{{ .Heading }}`, and `{{ .Content }}` are placeholders for dynamic content.

#### Step 2: Define a Struct to Pass Data

Create a Go struct to hold the data that will be passed to the template.

```go
type PageData struct {
    Title   string
    Heading string
    Content string
}
```

#### Step 3: Load and Execute the Template

Here’s how to load the template file, populate it with data, and serve it in response to HTTP requests:

```go
package main

import (
    "html/template"
    "net/http"
    "log"
)

type PageData struct {
    Title   string
    Heading string
    Content string
}

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        tmpl, err := template.ParseFiles("templates/index.html")
        if err != nil {
            log.Fatal(err)
        }

        data := PageData{
            Title:   "Welcome Page",
            Heading: "Hello, World!",
            Content: "Welcome to our website!",
        }

        tmpl.Execute(w, data)
    })

    log.Println("Starting server on :8080")
    http.ListenAndServe(":8080", nil)
}
```

In this example:

1. `template.ParseFiles("templates/index.html")` loads the HTML template.
2. `tmpl.Execute(w, data)` injects `PageData` values into the template, replacing placeholders.

### Example 1: Passing Dynamic Lists to Templates

You can pass more complex data types like slices (lists) to templates. Suppose we want to display a list of articles.

1. Update the `index.html` template to include a section for displaying articles:

   ```html
   <!DOCTYPE html>
   <html lang="en">
     <head>
       <meta charset="UTF-8" />
       <meta name="viewport" content="width=device-width, initial-scale=1.0" />
       <title>{{ .Title }}</title>
     </head>
     <body>
       <h1>{{ .Heading }}</h1>
       <p>{{ .Content }}</p>

       <h2>Articles</h2>
       <ul>
         {{ range .Articles }}
         <li>{{ . }}</li>
         {{ end }}
       </ul>
     </body>
   </html>
   ```

2. Modify the `PageData` struct and update the handler function to include a list of articles:

   ```go
   type PageData struct {
       Title    string
       Heading  string
       Content  string
       Articles []string
   }

   func main() {
       http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
           tmpl, err := template.ParseFiles("templates/index.html")
           if err != nil {
               log.Fatal(err)
           }

           data := PageData{
               Title:   "Welcome to the Blog",
               Heading: "Hello, Reader!",
               Content: "Here are our latest articles.",
               Articles: []string{"Article 1", "Article 2", "Article 3"},
           }

           tmpl.Execute(w, data)
       })

       log.Println("Starting server on :8080")
       http.ListenAndServe(":8080", nil)
   }
   ```

In this example, `{{ range .Articles }}{{ . }}{{ end }}` iterates over each item in `Articles` and displays it in an unordered list.

### Example 2: Conditional Rendering

Templates also support conditionals with `if` actions. Here’s an example that displays a different message if there are no articles.

1. Update `index.html` with an `if` action:

   ```html
   <!DOCTYPE html>
   <html lang="en">
     <head>
       <meta charset="UTF-8" />
       <meta name="viewport" content="width=device-width, initial-scale=1.0" />
       <title>{{ .Title }}</title>
     </head>
     <body>
       <h1>{{ .Heading }}</h1>
       <p>{{ .Content }}</p>

       <h2>Articles</h2>
       {{ if .Articles }}
       <ul>
         {{ range .Articles }}
         <li>{{ . }}</li>
         {{ end }}
       </ul>
       {{ else }}
       <p>No articles available at the moment.</p>
       {{ end }}
     </body>
   </html>
   ```

2. Adjust the Go code to test both cases (with and without articles):

   ```go
   type PageData struct {
       Title    string
       Heading  string
       Content  string
       Articles []string
   }

   func main() {
       http.HandleFunc("/with-articles", func(w http.ResponseWriter, r *http.Request) {
           tmpl, err := template.ParseFiles("templates/index.html")
           if err != nil {
               log.Fatal(err)
           }

           data := PageData{
               Title:    "Articles Page",
               Heading:  "Available Articles",
               Content:  "Here are our articles:",
               Articles: []string{"Article 1", "Article 2", "Article 3"},
           }

           tmpl.Execute(w, data)
       })

       http.HandleFunc("/no-articles", func(w http.ResponseWriter, r *http.Request) {
           tmpl, err := template.ParseFiles("templates/index.html")
           if err != nil {
               log.Fatal(err)
           }

           data := PageData{
               Title:    "Articles Page",
               Heading:  "Available Articles",
               Content:  "Here are our articles:",
               Articles: nil,  // Empty slice or nil to simulate no articles
           }

           tmpl.Execute(w, data)
       })

       log.Println("Starting server on :8080")
       http.ListenAndServe(":8080", nil)
   }
   ```

In this example, if there are no articles (i.e., `Articles` is `nil` or empty), the template will show "No articles available at the moment."

### Example 3: Using Nested Templates

You can split templates into smaller files and include them within each other. For instance, you might have a `header.html` and `footer.html` included in multiple pages.

1. Create a `base.html` file:

   ```html
   <!DOCTYPE html>
   <html lang="en">
     <head>
       <meta charset="UTF-8" />
       <meta name="viewport" content="width=device-width, initial-scale=1.0" />
       <title>{{ .Title }}</title>
     </head>
     <body>
       {{ template "header" . }} {{ template "content" . }} {{ template "footer"
       . }}
     </body>
   </html>
   ```

2. Define `header.html` and `footer.html`:

   ```html
   {{ define "header" }}
   <header>
     <h1>Website Header</h1>
   </header>
   {{ end }}
   ```

   ```html
   {{ define "footer" }}
   <footer>
     <p>Website Footer</p>
   </footer>
   {{ end }}
   ```

3. Create the main content template, `content.html`:

   ```html
   {{ define "content" }}
   <h2>{{ .Heading }}</h2>
   <p>{{ .Content }}</p>
   {{ end }}
   ```

4. Update your Go code to parse all templates and use them:

   ```go
   func main() {
       http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
           tmpl, err := template.ParseFiles(
               "templates/base.html",
               "templates/header.html",
               "templates/footer.html",
               "templates/content.html",
           )
           if err != nil {
               log.Fatal(err)
           }

           data := PageData{
               Title:   "Welcome",
               Heading: "Dynamic Content with Go Templates",
               Content: "This is the content of the page.",
           }

           tmpl.ExecuteTemplate(w, "base.html", data)
       })

       log.Println("Starting server on :8080")
       http.ListenAndServe(":8080", nil)
   }
   ```

In this example, the `base.html` template includes the `header`, `content`, and `footer`

templates, allowing you to modularize and reuse templates across pages.
