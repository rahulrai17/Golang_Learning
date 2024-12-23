# Passing Data to the Template

The best way to pass data to a template in Go depends on the complexity of the data and the structure of your application. Here’s a breakdown of the most effective approaches:

---

### **1. Use a `struct` to Pass Data**

The most common and structured way to pass data to a template is to define a `struct`. A `struct` helps you organize related data in a clear and reusable way.

#### **Example**:

```go
type TemplateData struct {
    StringMap map[string]string       // A map for storing string-based key-value pairs.
    IntMap    map[string]int          // A map for integer-based key-value pairs.
    FloatMap  map[string]float64      // A map for float-based key-value pairs.
    Data      map[string]interface{}  // A generic map for holding data of any type.
    CSRFToken string                  // A security token for protecting against CSRF attacks.
    Flash     string                  // A message shown to the user, typically after a successful action.
    Warning   string                  // A warning message to display in the UI.
    Error     string                  // An error message to display in the UI.
}

```

- You can add fields to this `TemplateData` struct to hold all the dynamic data you want to pass to the template.
- Use it in your handlers to populate data and pass it to the `RenderTemplate` function.
- This is a solid practice for larger applications where you need flexibility and scalability.
- For smaller apps or templates with very specific data needs, creating separate structs for each template might be simpler and more explicit.
- The choice depends on your app’s complexity and how reusable you want your data-passing mechanism to be.

#### **Usage**:

```go
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
    data := TemplateData{
        StringMap: map[string]string{"welcome": "Welcome to our site!"},
    }
    render.RenderTemplate(w, "home.page.tmpl", &data)
}
```

---

### **2. Use a Centralized TemplateData Struct**

If your app frequently passes similar types of data to templates (like flash messages, user session data, or CSRF tokens), create a helper function that adds these common data fields.

#### **Example**:

```go
func AddDefaultData(td *TemplateData, r *http.Request) *TemplateData {
    td.CSRFToken = "example-csrf-token" // Add CSRF token
    td.Flash = "Sample flash message"
    td.Error = "Sample error message"
    td.Warning = "Sample warning"
    return td
}
```

- Call this function in your handlers before rendering the template.

#### **Usage**:

```go
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
    data := TemplateData{
        StringMap: map[string]string{"title": "About Us"},
    }
    data = *AddDefaultData(&data, r) // Add default fields
    render.RenderTemplate(w, "about.page.tmpl", &data)
}
```

---

### **3. Use `map[string]interface{}` for Simplicity**

If the data structure is simple or one-off, you can use a `map[string]interface{}` to pass dynamic data to the template.

#### **Example**:

```go
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
    data := map[string]interface{}{
        "Title":   "Welcome Home",
        "Message": "Hello, World!",
    }
    render.RenderTemplate(w, "home.page.tmpl", data)
}
```

#### **When to Use This**:

- Use `map[string]interface{}` for quick, simple, or one-off data passing.
- For more complex applications, prefer a well-structured `struct`.

---

### **4. Combine Static and Dynamic Data in Templates**

In some cases, templates may include static and dynamic content. You can use the `TemplateData` struct for dynamic content while keeping static elements in the template file itself.

#### **Example**:

```html
<!-- home.page.tmpl -->
<h1>{{.StringMap.welcome}}</h1>
<p>This is static content written directly in the template.</p>
```

---

### **5. Use a Template Rendering Helper**

To streamline rendering templates, you can write a helper function that abstracts the `template.Execute` logic.

#### **Example**:

```go
package render

import (
    "html/template"
    "net/http"
)

func RenderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
    parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
    err := parsedTemplate.Execute(w, data)
    if err != nil {
        http.Error(w, "Unable to render template", http.StatusInternalServerError)
    }
}
```

#### **Usage in Handlers**:

```go
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
    data := TemplateData{
        StringMap: map[string]string{"welcome": "Welcome Home"},
    }
    render.RenderTemplate(w, "home.page.tmpl", &data)
}
```

---

### **6. Leverage Template Caching for Performance**

To improve performance, especially for production environments, cache your templates instead of parsing them on every request.

#### **Example Template Caching**:

```go
var templateCache = make(map[string]*template.Template)

func LoadTemplates() {
    templates := []string{"home.page.tmpl", "about.page.tmpl"}
    for _, tmpl := range templates {
        parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
        templateCache[tmpl] = parsedTemplate
    }
}

func RenderTemplateFromCache(w http.ResponseWriter, tmpl string, data interface{}) {
    tmplToRender := templateCache[tmpl]
    tmplToRender.Execute(w, data)
}
```

#### **Usage**:

```go
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
    data := TemplateData{
        StringMap: map[string]string{"title": "About Us"},
    }
    render.RenderTemplateFromCache(w, "about.page.tmpl", &data)
}
```

---

### **Best Practices**

1. **Use a `struct` for Reusability**:

   - Define a `TemplateData` struct for consistency and scalability.
   - Use helper functions to populate shared data like session info or CSRF tokens.

2. **Avoid Parsing Templates on Every Request**:

   - Implement template caching for better performance in production.

3. **Keep Handlers Focused**:

   - Let handlers focus on preparing data. Keep rendering logic in a separate package (like `render`).

4. **Validate Template Fields**:
   - Ensure that the fields passed to templates match the expected fields in the HTML files to avoid runtime errors.

# Important Note

- Suppose you are created this TemplateData inside the `handlers.go` this will make
  - `handler.go` imports `render.go` because it needs its functionality.
  - `render.go` tries to import `handler.go`, creating a loop where:
    - handler.go depends on render.go.
    - render.go depends on handler.go.
  - This creates a circular reference that Go cannot resolve. Go does not allow circular imports, and this will result in a compile-time error.
  - Why does this happen? In Go, the import system is single-pass and doesn’t allow cycles. When one file imports another, it expects all its dependencies to be resolved without loops.
  - In simple terms since `handlers.go` already imports some features of `render.go`. Go will not allow you render.go to import `handlers.go`, therefore to resolve this issue we are needed to creates a sperate file so we will be able to use the `TemplateData` struct.
  - The easiest way to resolve this will be creating a `models` package and making `templateDate.go` to make this work .
  - There are many other ways also :
    - `Use Interfaces` : If render depends on a specific function or type in handler, you can define an interface in render that handler can implement. This avoids direct imports between the two files.
    - you can research other methods also.
