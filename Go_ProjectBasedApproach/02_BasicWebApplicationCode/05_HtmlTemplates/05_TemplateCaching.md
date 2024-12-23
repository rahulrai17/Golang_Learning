## Template Caching in golang

Template caching in Go (Golang) is a common technique used to improve the performance of applications that render templates repeatedly. By caching parsed templates in memory, you avoid the overhead of reading and parsing template files from disk each time they are needed.

Go’s `text/template` and `html/template` packages are used for text and HTML templating, respectively. Both allow the creation of templates with logic like loops, conditions, and variables, making them versatile for rendering content dynamically.

---

### **What is Template Caching?**

Template caching involves preloading and storing parsed templates in memory (e.g., in a `map` or `sync.Map`) so they can be reused multiple times without needing to re-read or re-parse the templates. This is especially useful for applications that frequently render the same set of templates, such as web servers.

---

### **Why Use Template Caching?**

1. **Performance Optimization**: Parsing templates can be computationally expensive. By caching templates, you eliminate the need to re-parse them for each request.
2. **Centralized Management**: Caching templates allows centralized access to templates, which is helpful in complex applications.
3. **Scalability**: Reduces I/O and processing bottlenecks, allowing the application to scale more effectively.

---

### **Key Takeaways**

1. **Preloading vs. Lazy Loading**: Preloading templates is better for known templates, while lazy loading is useful when templates are dynamic or loaded on demand.
2. **Concurrency Safety**: Use `sync.RWMutex` or `sync.Map` for thread-safe template caching.
3. **Development Flexibility**: Use file watchers to reload templates dynamically during development.
4. **Testing and Debugging**: Cache invalidation or flushing mechanisms can be handy during testing or template updates.

By caching templates intelligently, you can achieve significant performance improvements while maintaining the flexibility needed for development and production environments.v

### **Template Caching with Standard Library**

The standard library in Go provides the `html/template` (or `text/template`) package for working with templates. Below is an example of template caching implemented entirely with the standard library:

---

#### **Standard Library Implementation**

```go
package main

import (
	"html/template"
	"log"
	"net/http"
	"sync"
)

// TemplateCache stores parsed templates with thread-safety
type TemplateCache struct {
	mu       sync.RWMutex
	templates map[string]*template.Template
}

// NewTemplateCache initializes the template cache
func NewTemplateCache() *TemplateCache {
	return &TemplateCache{
		templates: make(map[string]*template.Template),
	}
}

// Get retrieves a cached template or parses and caches it
func (tc *TemplateCache) Get(name, path string) (*template.Template, error) {
	tc.mu.RLock()
	tmpl, exists := tc.templates[name]
	tc.mu.RUnlock()

	if exists {
		return tmpl, nil
	}

	// Parse template and cache it
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		return nil, err
	}

	tc.mu.Lock()
	tc.templates[name] = tmpl
	tc.mu.Unlock()

	return tmpl, nil
}

func main() {
	// Initialize template cache
	cache := NewTemplateCache()

	// Example handler using the cached templates
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := cache.Get("index", "templates/index.html")
		if err != nil {
			http.Error(w, "Template not found", http.StatusInternalServerError)
			return
		}

		data := map[string]string{"Title": "Hello from Template Cache!"}
		if err := tmpl.Execute(w, data); err != nil {
			http.Error(w, "Failed to render template", http.StatusInternalServerError)
		}
	})

	log.Println("Server started at :8080")
	http.ListenAndServe(":8080", nil)
}
```

---

### **Template Caching with Popular Library**

One of the most popular libraries for web development in Go is **[gin-gonic/gin](https://github.com/gin-gonic/gin)**. Gin has built-in support for HTML rendering but doesn’t include template caching by default. We can extend Gin to add caching functionality.

#### **Gin Example with Template Caching**

```go
package main

import (
	"html/template"
	"log"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

// TemplateCache stores parsed templates with thread-safety
type TemplateCache struct {
	mu       sync.RWMutex
	templates map[string]*template.Template
}

// NewTemplateCache initializes the template cache
func NewTemplateCache() *TemplateCache {
	return &TemplateCache{
		templates: make(map[string]*template.Template),
	}
}

// Load parses and caches a template
func (tc *TemplateCache) Load(name, path string) error {
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		return err
	}

	tc.mu.Lock()
	tc.templates[name] = tmpl
	tc.mu.Unlock()
	return nil
}

// Get retrieves a cached template
func (tc *TemplateCache) Get(name string) (*template.Template, bool) {
	tc.mu.RLock()
	tmpl, exists := tc.templates[name]
	tc.mu.RUnlock()
	return tmpl, exists
}

func main() {
	// Initialize Gin and template cache
	r := gin.Default()
	cache := NewTemplateCache()

	// Preload templates
	err := cache.Load("index", "templates/index.html")
	if err != nil {
		log.Fatalf("Failed to load template: %v", err)
	}

	// Define a custom HTML render function
	r.GET("/", func(c *gin.Context) {
		tmpl, exists := cache.Get("index")
		if !exists {
			c.String(http.StatusInternalServerError, "Template not found")
			return
		}

		// Use Gin's response writer
		data := map[string]string{"Title": "Hello from Gin with Template Cache!"}
		if err := tmpl.Execute(c.Writer, data); err != nil {
			c.String(http.StatusInternalServerError, "Failed to render template")
		}
	})

	log.Println("Server started at :8080")
	r.Run(":8080")
}
```

---

### **Comparison**

| Feature         | Standard Library                                    | With Gin (`gin-gonic/gin`)                           |
| --------------- | --------------------------------------------------- | ---------------------------------------------------- |
| **Ease of Use** | Requires manual setup for HTTP handling and caching | Provides built-in routing; caching is added manually |
| **Performance** | Fast but requires boilerplate for caching           | Fast, with less boilerplate for routing              |
| **Flexibility** | Fully customizable                                  | Integrates well with Gin's ecosystem                 |
| **Use Case**    | Good for lightweight apps or custom frameworks      | Ideal for building REST APIs or web apps             |

Both implementations follow the same caching logic but differ in how they integrate with the routing and HTTP layer. For a lightweight application or custom server, the standard library works well. For a larger web application, using Gin simplifies development.
