## Layouts in Golang

In Go (Golang), when building web applications, layouts are commonly used to define reusable structures for templates. Layouts allow you to share common components like headers, footers, and navigation menus across multiple pages while still allowing the pages to have their unique content.

Go's standard library includes the `html/template` package, which provides robust support for templates and layouts. Here's a breakdown of how to implement layouts in Go templates.

---

### **Step 1: Define the Layout**

The layout is a base template containing the shared structure. Use `{{define}}` to create named template blocks and `{{template}}` to include content.

**Example: `layout.html`**

```html
<!DOCTYPE html>
<html>
  <head>
    <title>{{block "title" .}}Default Title{{end}}</title>
  </head>
  <body>
    <header>
      <h1>My Website</h1>
    </header>
    <main>{{block "content" .}}{{end}}</main>
    <footer>
      <p>&copy; 2024 My Website</p>
    </footer>
  </body>
</html>
```

---

### **Step 2: Create Specific Templates**

Other templates can override blocks defined in the layout.

**Example: `home.html`**

```html
{{define "title"}}Home Page{{end}} {{define "content"}}
<h2>Welcome to the Home Page</h2>
<p>This is the content of the home page.</p>
{{end}}
```

**Example: `about.html`**

```html
{{define "title"}}About Us{{end}} {{define "content"}}
<h2>About Us</h2>
<p>This page provides information about us.</p>
{{end}}
```

---

### **Step 3: Parse Templates**

In your Go application, you need to parse the layout and content templates together. The `html/template` package allows you to parse multiple templates and render them.

**Example Code**

```go
package main

import (
	"html/template"
	"net/http"
)

var tmpl = template.Must(template.ParseFiles("layout.html", "home.html", "about.html"))

func homeHandler(w http.ResponseWriter, r *http.Request) {
	err := tmpl.ExecuteTemplate(w, "layout", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	err := tmpl.ExecuteTemplate(w, "layout", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/about", aboutHandler)

	http.ListenAndServe(":8080", nil)
}
```

---

### **Step 4: Render the Appropriate Page**

Use the `ExecuteTemplate` method to render a specific named template. For example:

- In `homeHandler`, you can render the `"home.html"` template by calling:
  ```go
  tmpl.ExecuteTemplate(w, "home", nil)
  ```
- Similarly, render `"about.html"` in `aboutHandler`:
  ```go
  tmpl.ExecuteTemplate(w, "about", nil)
  ```

---

### **Improved Example: Dynamic Data**

Pass dynamic data to templates by using structs or maps.

```go
type PageData struct {
	Title string
	Body  string
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	data := PageData{
		Title: "Home Page",
		Body:  "Welcome to the Home Page",
	}
	err := tmpl.ExecuteTemplate(w, "layout", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
```

---

### **Tips for Better Layout Management**

1. **Use Directory Structure**:

   - Store layout and page-specific templates in separate folders.
   - Example:
     ```
     templates/
       layout.html
       pages/
         home.html
         about.html
     ```

2. **Use Template Caching**:

   - Pre-parse all templates at startup to avoid repeated file reads.

3. **Embed Templates**:
   - Use the `embed` package (Go 1.16+) to embed templates into the binary for easier deployment.

By following this approach, you can efficiently manage layouts and create a consistent structure for your web application.
