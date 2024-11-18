## Error Checking in Golang

In Go, error handling is explicit, which makes it clear and reliable but requires careful handling. Errors in Go are values, so rather than using exceptions or error codes, Go encourages handling errors explicitly at each point of failure. Here’s a comprehensive guide on different ways to handle errors in Go, from basic error checking to best practices.

### 1. **Basic Error Checking**

The simplest way to handle errors in Go is to check if an error is `nil` and act accordingly.

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    file, err := os.Open("nonexistent.txt")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer file.Close()
    fmt.Println("File opened successfully")
}
```

- **Explanation**: Here, we check if `err` is `nil`. If not, we print the error and return.
- **Use case**: Suitable for straightforward cases where a single operation might fail.

### 2. **Custom Error Messages**

If the standard error message is not descriptive enough, you can create a custom error message.

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    _, err := os.Open("nonexistent.txt")
    if err != nil {
        fmt.Printf("Failed to open file: %v\n", err)
        return
    }
}
```

- **Explanation**: We add context to the error message, which helps clarify where the error occurred and why.
- **Use case**: Helpful when you want to add specific details to an error, making it easier to troubleshoot.

### 3. **Custom Error Types**

Go allows you to create custom error types by implementing the `error` interface, which has a single method `Error() string`.

```go
package main

import (
    "errors"
    "fmt"
)

// CustomError is a user-defined error type
type CustomError struct {
    Code int
    Msg  string
}

func (e *CustomError) Error() string {
    return fmt.Sprintf("Error %d: %s", e.Code, e.Msg)
}

func mayFail() error {
    return &CustomError{Code: 404, Msg: "Resource not found"}
}

func main() {
    err := mayFail()
    if err != nil {
        fmt.Println("Error:", err)
    }
}
```

- **Explanation**: `CustomError` is a struct with additional information about the error. By implementing the `Error()` method, it satisfies the `error` interface.
- **Use case**: When you need more detail than a simple message, or when you want to provide additional context to errors, such as error codes or metadata.

### 4. **Using `errors.New` and `fmt.Errorf`**

The `errors.New` function and `fmt.Errorf` function are simple ways to create errors.

#### Using `errors.New`

```go
package main

import (
    "errors"
    "fmt"
)

func mayFail() error {
    return errors.New("something went wrong")
}

func main() {
    err := mayFail()
    if err != nil {
        fmt.Println("Error:", err)
    }
}
```

#### Using `fmt.Errorf`

```go
package main

import (
    "fmt"
)

func mayFail() error {
    return fmt.Errorf("operation failed: %w", errors.New("network error"))
}

func main() {
    err := mayFail()
    if err != nil {
        fmt.Println("Error:", err)
    }
}
```

- **Explanation**: `errors.New` creates a simple error, while `fmt.Errorf` allows adding formatted strings and wrapping errors with `%w`.
- **Use case**: When you want quick, descriptive errors, or when you want to wrap an error with more context.

### 5. **Error Wrapping with `%w` and `errors.Is`/`errors.As`**

Go 1.13 introduced error wrapping, which allows you to wrap an error within another error, preserving the original error information.

```go
package main

import (
    "errors"
    "fmt"
)

var ErrNotFound = errors.New("resource not found")

func mayFail() error {
    return fmt.Errorf("failed to fetch resource: %w", ErrNotFound)
}

func main() {
    err := mayFail()
    if errors.Is(err, ErrNotFound) {
        fmt.Println("The error was due to missing resource.")
    } else {
        fmt.Println("Another error occurred:", err)
    }
}
```

- **Explanation**: The `fmt.Errorf` function wraps `ErrNotFound` within a new error, preserving the original. `errors.Is` checks if `err` is of type `ErrNotFound`.
- **Use case**: Useful for error checking without losing context about the root cause.

### 6. **Using `errors.As` to Check Error Types**

`errors.As` is useful when working with custom error types. It allows you to check if an error is of a specific type and extract it.

```go
package main

import (
    "errors"
    "fmt"
)

type CustomError struct {
    Code int
    Msg  string
}

func (e *CustomError) Error() string {
    return fmt.Sprintf("Error %d: %s", e.Code, e.Msg)
}

func mayFail() error {
    return &CustomError{Code: 403, Msg: "forbidden access"}
}

func main() {
    err := mayFail()
    var customErr *CustomError
    if errors.As(err, &customErr) {
        fmt.Println("Custom Error:", customErr.Code, customErr.Msg)
    } else {
        fmt.Println("Other error:", err)
    }
}
```

- **Explanation**: `errors.As` checks if `err` is of type `CustomError` and, if so, assigns it to `customErr`.
- **Use case**: Useful when working with multiple error types, allowing you to handle each type appropriately.

### 7. **Deferring Error Handling**

When handling multiple errors in a function, you can use `defer` to simplify error handling, especially for cleanup actions.

```go
package main

import (
    "fmt"
    "os"
)

func readFile(filename string) (err error) {
    file, err := os.Open(filename)
    if err != nil {
        return fmt.Errorf("failed to open file: %w", err)
    }
    defer func() {
        if cerr := file.Close(); cerr != nil {
            err = fmt.Errorf("failed to close file: %v, original error: %w", cerr, err)
        }
    }()
    // ...read file contents
    return nil
}
```

- **Explanation**: `defer` allows handling errors on cleanup actions, like closing files, and ensures that all errors are handled appropriately.
- **Use case**: When dealing with resources that require cleanup, such as files or network connections.

### 8. **Logging Errors**

Error logging is crucial for debugging and maintaining applications. Use Go's `log` package or custom logging libraries like `zap` or `logrus` for structured logging.

```go
package main

import (
    "log"
    "os"
)

func main() {
    file, err := os.Open("nonexistent.txt")
    if err != nil {
        log.Printf("Error opening file: %v\n", err)
        return
    }
    defer file.Close()
    log.Println("File opened successfully")
}
```

- **Explanation**: `log.Printf` outputs the error with context, which is useful for production environments.
- **Use case**: Essential for capturing error details in logs, especially for services running in production.

### 9. **Best Practices for Error Handling in Go**

- **Check errors immediately**: Always check if an error occurred right after calling a function.
- **Use descriptive error messages**: Provide meaningful context with `fmt.Errorf`.
- **Avoid panics**: Reserve `panic` for truly exceptional situations where the program cannot continue.
- **Wrap errors**: Use `%w` to wrap errors and add context.
- **Use `errors.Is` and `errors.As`**: For error comparison and extracting custom error types.
- **Log errors in production code**: Logging allows for monitoring and troubleshooting in production.
- **Return errors rather than logging in libraries**: Let the calling code decide how to handle and log the error.

By following these practices, you can handle errors in a way that is both readable and robust, providing valuable context for debugging and maintaining your Go applications.

## Best way to check error in golang web development

In web development with Go, effective error handling is critical for building robust applications. Here are the best practices for handling errors specifically in the context of web development, along with techniques to improve error handling in APIs, middleware, and HTTP responses.

### 1. **Return and Log Detailed Errors**

In a web context, the HTTP handlers should return errors that include specific details, and any sensitive information should be omitted from the response but logged for later troubleshooting.

#### Example

```go
package main

import (
    "log"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    if err := doSomething(); err != nil {
        log.Printf("internal error: %v", err) // Log detailed error message
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }
    w.Write([]byte("Success"))
}

func doSomething() error {
    return fmt.Errorf("database connection failed")
}
```

- **Explanation**: In this example, `http.Error` sends a generic error message to the user, while the detailed error is logged. This helps keep sensitive information safe while giving developers enough information to debug.
- **Best for**: Errors that occur in API handlers or where sensitive information must not be revealed.

### 2. **Using `errors.Is` and `errors.As` for Custom Error Types**

When using custom error types (such as for validation, authorization, or specific business logic errors), use `errors.Is` and `errors.As` to check and respond based on error types.

#### Example

```go
package main

import (
    "errors"
    "fmt"
    "net/http"
)

var ErrUnauthorized = errors.New("unauthorized access")

func handler(w http.ResponseWriter, r *http.Request) {
    if err := authenticateUser(r); err != nil {
        if errors.Is(err, ErrUnauthorized) {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }
    w.Write([]byte("Welcome!"))
}

func authenticateUser(r *http.Request) error {
    // Simulate an authorization failure
    return ErrUnauthorized
}
```

- **Explanation**: `errors.Is` allows us to identify the `ErrUnauthorized` error and respond with `http.StatusUnauthorized`. This approach is ideal for error types that map directly to specific HTTP statuses.
- **Best for**: Errors that represent specific HTTP statuses like `404 Not Found`, `401 Unauthorized`, or `400 Bad Request`.

### 3. **Using Middleware for Centralized Error Handling**

Middleware can be used to handle and log errors in a centralized manner. This is helpful for logging all errors consistently or transforming errors into a standard response format.

#### Example

```go
package main

import (
    "log"
    "net/http"
)

func main() {
    http.Handle("/", errorHandlingMiddleware(http.HandlerFunc(handler)))
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func errorHandlingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        defer func() {
            if err := recover(); err != nil {
                log.Printf("Recovered from panic: %v", err)
                http.Error(w, "Internal Server Error", http.StatusInternalServerError)
            }
        }()
        next.ServeHTTP(w, r)
    })
}

func handler(w http.ResponseWriter, r *http.Request) {
    panic("something went wrong!") // Simulating a panic for demonstration
}
```

- **Explanation**: The middleware catches and logs panics. A similar approach can be applied for logging or transforming responses for every request.
- **Best for**: Global error handling, logging, and transforming errors into standardized HTTP responses.

### 4. **Wrap Errors with Context Using `%w` in `fmt.Errorf`**

For more complex applications, wrap errors with additional context using `fmt.Errorf` and `%w`. This provides a traceable error chain, which helps when debugging complex issues.

#### Example

```go
package main

import (
    "fmt"
    "log"
)

func main() {
    err := handler()
    if err != nil {
        log.Printf("Error: %v", err)
    }
}

func handler() error {
    if err := doSomething(); err != nil {
        return fmt.Errorf("handler failed: %w", err)
    }
    return nil
}

func doSomething() error {
    return fmt.Errorf("database connection failed: %w", ErrDatabase)
}
```

- **Explanation**: Wrapping errors allows you to track the error’s source through each layer. This is useful for troubleshooting layered applications with multiple components, like databases, caching, or external APIs.
- **Best for**: Complex applications with multiple dependencies or service layers.

### 5. **Structured Error Responses for APIs**

When building REST APIs, standardize error responses in JSON format so that clients can parse and handle them effectively.

#### Example

```go
package main

import (
    "encoding/json"
    "net/http"
)

type ErrorResponse struct {
    Message string `json:"message"`
    Code    int    `json:"code"`
}

func writeJSONError(w http.ResponseWriter, message string, code int) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(code)
    json.NewEncoder(w).Encode(ErrorResponse{Message: message, Code: code})
}

func handler(w http.ResponseWriter, r *http.Request) {
    if err := doSomething(); err != nil {
        writeJSONError(w, "Failed to process request", http.StatusInternalServerError)
        return
    }
    w.Write([]byte("Success"))
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}
```

- **Explanation**: `writeJSONError` function standardizes the error response format, allowing clients to handle errors consistently.
- **Best for**: REST API development where clients expect JSON responses.

### 6. **Using `log.Fatal` and `log.Printf` for Startup and Fatal Errors**

Use `log.Fatal` or `log.Fatalf` to handle startup errors that prevent the application from running (like configuration errors). For non-fatal errors, use `log.Printf` or another logging library.

#### Example

```go
package main

import (
    "log"
    "net/http"
)

func main() {
    if err := startServer(); err != nil {
        log.Fatalf("Server startup failed: %v", err)
    }
}

func startServer() error {
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        return fmt.Errorf("failed to start server: %w", err)
    }
    return nil
}
```

- **Explanation**: Using `log.Fatal` ensures the application stops if a fatal error occurs at startup.
- **Best for**: Startup errors or configuration issues that require the application to stop.

### 7. **Retry Mechanism for Temporary Errors**

For operations that may fail temporarily (like network or database calls), implement a retry mechanism. Use libraries like `go-retryablehttp` or write a custom retry function.

#### Example

```go
package main

import (
    "fmt"
    "net/http"
    "time"
)

func retry(attempts int, sleep time.Duration, fn func() error) error {
    for i := 0; i < attempts; i++ {
        if err := fn(); err != nil {
            if i == attempts-1 {
                return fmt.Errorf("after %d attempts, last error: %w", attempts, err)
            }
            time.Sleep(sleep)
            continue
        }
        return nil
    }
    return fmt.Errorf("failed after %d attempts", attempts)
}

func main() {
    err := retry(3, 2*time.Second, func() error {
        resp, err := http.Get("http://example.com")
        if err != nil {
            return err
        }
        defer resp.Body.Close()
        return nil
    })
    if err != nil {
        fmt.Println("Operation failed:", err)
    }
}
```

- **Explanation**: The `retry` function retries the operation up to a specified number of attempts.
- **Best for**: Transient errors, such as network or database connection issues.

### Summary of Best Practices for Error Handling in Web Development

| Practice                                  | Use Case                                                                         |
| ----------------------------------------- | -------------------------------------------------------------------------------- |
| **Return and Log Detailed Errors**        | Keep sensitive details out of responses; log full errors internally.             |
| **Use `errors.Is` and `errors.As`**       | For matching and handling specific errors (e.g., validation, auth errors).       |
| **Centralized Error Handling Middleware** | Simplifies handling and logging errors at the middleware level.                  |
| **Error Wrapping with `%w`**              | Trace errors across layers for debugging.                                        |
| **Structured Error Responses**            | For standardized, parseable API responses.                                       |
| **Use `log.Fatal` and `log.Printf`**      | Use for startup issues or non-fatal errors in logs for monitoring.               |
| **Retry Mechanism for Temporary Errors**  | Handles transient failures gracefully, especially in network-heavy applications. |

Applying these practices will help ensure that your web applications are reliable, secure, and easier to debug.
