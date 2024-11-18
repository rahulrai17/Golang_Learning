In Go, you can serve HTML pages in a variety of ways, depending on the complexity and structure of your application. Here are three common approaches:

### 1. **Serve Static HTML Files**

If you have basic HTML files (e.g., `index.html`, `about.html`), you can use Go's `http.ServeFile` to serve them directly from the filesystem. This is useful for simple sites with static content.

```go
package main

import (
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "static/index.html")
    })

    http.ListenAndServe(":8080", nil)
}
```

In this example, `static/index.html` will be served when users navigate to the root path (`/`). To serve more files, you could create additional handlers or use `http.FileServer`.

### 2. **Serve Dynamic HTML Using Templates**

If you want to serve HTML with dynamic content, you can use Go's `html/template` package. This allows you to define templates with placeholders for dynamic data and render them in response to HTTP requests.

**Example with `html/template`:**

1. Create a template file `template.html`:

   ```html
   <!DOCTYPE html>
   <html lang="en">
     <head>
       <meta charset="UTF-8" />
       <meta name="viewport" content="width=device-width, initial-scale=1.0" />
       <title>{{ .Title }}</title>
     </head>
     <body>
       <h1>{{ .Content }}</h1>
     </body>
   </html>
   ```

2. Serve the template with Go:

   ```go
   package main

   import (
       "html/template"
       "net/http"
   )

   type PageData struct {
       Title   string
       Content string
   }

   func main() {
       http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
           tmpl := template.Must(template.ParseFiles("template.html"))
           data := PageData{
               Title:   "Hello, World!",
               Content: "Welcome to my website!",
           }
           tmpl.Execute(w, data)
       })

       http.ListenAndServe(":8080", nil)
   }
   ```

In this example, Go reads `template.html`, injects the `Title` and `Content` fields from `PageData`, and serves the generated HTML to the user.

### 3. **Serve HTML Pages with a Framework (e.g., Gin)**

For larger applications, you might want to use a web framework like [Gin](https://github.com/gin-gonic/gin) to handle routing and templating in a more organized way. Gin simplifies routing and also has support for serving HTML templates.

**Example using Gin:**

1. Install Gin:

   ```bash
   go get -u github.com/gin-gonic/gin
   ```

2. Serve a template with Gin:

   ```go
   package main

   import (
       "github.com/gin-gonic/gin"
   )

   func main() {
       r := gin.Default()
       r.LoadHTMLGlob("templates/*.html")

       r.GET("/", func(c *gin.Context) {
           c.HTML(200, "index.html", gin.H{
               "title":   "Hello, World!",
               "content": "Welcome to my Gin-powered site!",
           })
       })

       r.Run(":8080")
   }
   ```

In this example, all templates inside the `templates` folder are available for rendering. The `gin.H` map passes dynamic data to the `index.html` template, which can be rendered similarly to the earlier example.

These methods provide flexibility for different application needs:

- Simple static serving,
- Dynamic templating with `html/template`,
- and full-fledged frameworks like Gin for more complex applications.
