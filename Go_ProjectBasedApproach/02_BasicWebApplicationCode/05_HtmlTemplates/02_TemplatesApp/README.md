## In this we will build a mini project to learn more about templates and How to Structure Code properly

### the project structure will be

```go
TemplatesApp/
    ├── cmd/
    │   └── web/
    │       └── main.go
    ├── pkg/
    │   ├── handler/
    │   │   └── handlers.go
    │   └── render/
    │       └── render.go
    ├── templates/
    │   └── home.html
    └── go.mod
```

### **Flow Diagram**

```
+------------+      +------------+      +-------------------+      +--------------+
|   Browser  | ---> |  main.go   | ---> |(handlers.go)      | ---> | render.go    |
+------------+      +------------+      +-------------------+      +--------------+
     |                    |                     |                       |
     |        Request     |      Call handler   |   Call RenderTemplate |
     |(eg- /home, /about) |    (e.g., /home)    |      (home.html)      |
     |                    |                     |                       |
     +------------------> |                     |                       |
                          |                     |                       |
                          +-------------------> |                       |
                                                |                       |
                                                +-----------------------> |
                                                                      Render HTML template and return to browser
```

### **Process Breakdown**:

1. **Browser**: The user sends a request (e.g., `/home` or `/about`).
2. **`main.go`**: The `main` function sets up the web server and listens for requests.
3. **Handler (`handlers.go`)**: The handler functions are triggered by specific routes and call the `RenderTemplate` function.
4. **`render.go`**: The `RenderTemplate` function loads and renders the appropriate HTML template (like `home.html`) and returns the response to the browser.

---

### How to run this code

### Step 1

```go
go mod init <application_name/folder_name>
```

### Step 1

```go
go mod tidy //this will install all dependencies that are needed
```

### step 3 - go to the location of main.go, then

```go
go run main.go
```

### Summary

This is a basic application the flow of the application is
