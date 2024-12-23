# Project for Practicing Routing and Middleware

This project involves creating a web application with routing and middleware using the `chi` package. The steps are documented here for future reference.

## Steps to Create This Project

- `STEP 1`: Initialize the project using `go mod init <name_of_your_project>`. This will create a `go.mod` file for the project.
- `STEP 2`: Run the command `go mod tidy` to install all the necessary packages.
- `STEP 3`: Structure the project folders. Start with a basic structure and expand it as the project grows.

  ```plaintext
  myproject/                   # Root of your project
    ├── cmd/                   # Entry points for your application(s)
    │   └── web/               # Main application folder (web server)
    │       └── main.go        # Main entry point for the web server
    ├── pkg/                   # Reusable packages for your project
    │   ├── config/            # Configuration-related code
    │   │   └── config.go      # Define configuration structures and load settings
    │   ├── handler/           # HTTP handler functions for routes
    │   ├── middleware/        # Middleware functions
    │   │   └── middleware.go  # Define middleware functions
    │   ├── model/             # Data structures and database models
    │   └── render/            # Rendering templates and views
    ├── static/                # Static files served by the web server
    │   ├── css/               # CSS files
    │   │   └── styles.css     # Example CSS file
    │   ├── js/                # JavaScript files
    │   │   └── scripts.js     # Example JavaScript file
    │   └── img/               # Images
    ├── templates/             # HTML templates for rendering views
    │   ├── layout/            # Layout templates
    │   │   └── base.html      # Base layout template
    │   ├── pages/             # Page-specific templates
    │   │   └── index.html     # Example template for the homepage
    │   └── partials/          # Reusable partial templates
    │       └── header.html    # Example partial for the header
    ├── go.mod                 # Go module definition
    ├── go.sum                 # Go module dependencies
    └── README.md              # Project documentation
  ```

- `STEP 4`: Set up `main.go` to start a server using the `chi` package for routing and middleware.

  1. Install `chi` using the command from the official `chi` repository on GitHub.
     ```sh
     go get -u github.com/go-chi/chi/v5
     ```
  2. Create a `routes.go` file to define your routes.

     ```go
     // routes.go
     package main

     import (
         "net/http"
         "github.com/go-chi/chi/v5"
         "github.com/go-chi/chi/v5/middleware"
     )

     // routes function sets up the router with necessary routes and middleware
     func routes() *chi.Mux {
         mux := chi.NewRouter()
         // Using built-in middleware for logging and recovering from panics
         mux.Use(middleware.Logger)
         mux.Use(middleware.Recoverer)
         // Defining a simple route for the homepage
         mux.Get("/", func(w http.ResponseWriter, r *http.Request) {
             w.Write([]byte("Welcome to the homepage!"))
         })
         return mux
     }
     ```

     **Summary**: The `routes.go` file sets up the router with necessary routes and middleware using the `chi` package. It includes built-in middleware for logging and recovering from panics and defines a simple route for the homepage.

  3. Create a `middleware.go` file to define custom middleware.

     ```go
     // middleware.go
     package main

     import (
         "net/http"
         "time"
     )

     // ExampleMiddleware logs the duration of each request
     func ExampleMiddleware(next http.Handler) http.Handler {
         return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
             start := time.Now()
             next.ServeHTTP(w, r)
             duration := time.Since(start)
             println("Request took", duration.String())
         })
     }
     ```

     **Summary**: The `middleware.go` file defines custom middleware functions. In this example, `ExampleMiddleware` logs the duration of each request.

  4. Update `main.go` to start the HTTP server and use the previously created router and middleware.

     ```go
     // main.go
     package main

     import (
         "net/http"
         "log"
     )

     // main function starts the HTTP server on port 8080
     func main() {
         mux := routes()
         // Adding custom middleware to the router
         mux.Use(ExampleMiddleware)
         portNumber := ":8080"
         log.Printf("Starting server on %s", portNumber)
         err := http.ListenAndServe(portNumber, mux)
         if err != nil {
             log.Fatal(err)
         }
     }
     ```

     **Summary**: The `main.go` file starts the HTTP server on port 8080. It uses the router and custom middleware defined in `routes.go` and `middleware.go`.

By following these steps, you will have a basic web application set up with routing and middleware using the `chi` package. The `routes.go` file defines the routes, and the `middleware.go` file contains custom middleware functions. This structure helps in organizing the code and making it more maintainable.
