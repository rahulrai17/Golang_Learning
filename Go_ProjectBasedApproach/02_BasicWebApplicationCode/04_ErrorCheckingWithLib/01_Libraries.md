## Libraries for errors and when to use what

In Go, there are several libraries that extend Go's native error handling capabilities. Each library has unique features suited to different error-handling needs, such as error wrapping, stack tracing, custom error types, and structured logging. Here are some popular libraries, their key features, and recommendations for when to use each one.

### 1. **`errors` (Standard Library)**

- **Features**: Native Go error handling with simple functions for error creation (`errors.New`) and wrapping (`fmt.Errorf` with `%w`).
- **Best for**: Basic error handling and wrapping. It’s the most lightweight option and avoids any external dependencies.
- **Usage**: Ideal for most applications, especially if they have straightforward error-handling needs.

```go
import (
    "errors"
    "fmt"
)

func main() {
    err := fmt.Errorf("failed to connect to server: %w", errors.New("network unreachable"))
    fmt.Println(err)
}
```

### 2. **`pkg/errors` by `github.com/pkg/errors`**

- **Features**: Stack tracing, error wrapping with context, and enhanced error handling functions like `errors.Wrap`, `errors.Cause`, `errors.WithMessage`.
- **Pros**: Allows tracking the root cause of an error with stack traces, making it easier to debug complex applications.
- **Cons**: The Go team stopped using `pkg/errors` in favor of Go’s built-in error handling, and it’s no longer actively maintained.
- **Best for**: Applications that need stack tracing but are not using Go modules, or if you're working in Go versions prior to 1.13.

```go
import (
    "github.com/pkg/errors"
    "fmt"
)

func main() {
    err := errors.Wrap(errors.New("network unreachable"), "failed to connect to server")
    fmt.Println(err)
    fmt.Printf("Cause: %v\n", errors.Cause(err))
}
```

### 3. **`xerrors` (golang.org/x/xerrors)**

- **Features**: Enhanced error wrapping and formatting with `%w` for wrapping errors and detailed error messages.
- **Pros**: `xerrors` was designed to work with Go 1.13’s new error handling features and offers richer error formatting and stack traces.
- **Cons**: `xerrors` was meant as a transition tool to ease migration to Go's native error handling and is now largely replaced by Go’s standard library.
- **Best for**: Use only if working with an older Go version (pre-1.13). In modern Go versions, prefer the standard library.

```go
import (
    "golang.org/x/xerrors"
    "fmt"
)

func main() {
    err := xerrors.Errorf("failed to open file: %w", xerrors.New("file not found"))
    fmt.Println(err)
}
```

### 4. **`emperror` (github.com/emperror/errors)**

- **Features**: Comprehensive error handling with error wrapping, stack tracing, key-value attributes, and JSON encoding.
- **Pros**: Supports detailed error tracking, structured error logging, and integrates well with structured loggers.
- **Cons**: Slightly more complex and heavier than simpler error libraries, especially if only basic error handling is needed.
- **Best for**: Applications requiring structured error logging, or applications where errors need to be enriched with metadata or context.

```go
import (
    "github.com/emperror/errors"
    "log"
)

func main() {
    err := errors.WithMessage(errors.New("DB error"), "failed to fetch user data")
    log.Println(err)
}
```

### 5. **`errgo` (github.com/juju/errgo)**

- **Features**: Custom error wrapping and classification, with methods like `Notef`, `Cause`, and `Maskf` for controlling error visibility.
- **Pros**: Provides an organized way to classify errors, making it easy to understand the error type and propagate it with additional information.
- **Cons**: Limited community support; may require more initial learning compared to simpler libraries.
- **Best for**: Complex applications where error classification is a priority, such as when handling different types of errors (e.g., user vs. system errors).

```go
import (
    "github.com/juju/errgo"
    "fmt"
)

func main() {
    err := errgo.Notef(errgo.New("failed to connect"), "connection error")
    fmt.Println(errgo.Cause(err))
}
```

### 6. **`errors` (by `go-micro.dev/v4/errors`)**

- **Features**: Custom errors with structured details such as error code and status, useful for gRPC or HTTP-based microservices.
- **Pros**: Provides specific fields for service and status codes, which can be very useful in distributed applications.
- **Cons**: Adds complexity and is only needed in distributed or service-based applications.
- **Best for**: Microservices and distributed applications, especially where error codes and status fields are necessary for API communication.

```go
import (
    "go-micro.dev/v4/errors"
    "fmt"
)

func main() {
    err := errors.New("auth", "unauthorized access", 401)
    fmt.Println(err)
}
```

### 7. **`zerolog` (github.com/rs/zerolog)**

- **Features**: Primarily a structured logger, but integrates error handling with detailed logs and key-value pairs.
- **Pros**: Excellent for applications where logging and error handling are closely integrated; allows for structured JSON logs.
- **Cons**: Adds dependency on a structured logging library, which may be more than required for simpler applications.
- **Best for**: Production-grade applications requiring detailed logging and error context, or if structured logging is already in place.

```go
import (
    "github.com/rs/zerolog/log"
    "errors"
)

func main() {
    err := errors.New("DB connection failed")
    log.Error().Err(err).Msg("could not connect to database")
}
```

---

### Choosing the Right Library

| **Use Case**                      | **Recommended Library**     | **Rationale**                                                                                   |
| --------------------------------- | --------------------------- | ----------------------------------------------------------------------------------------------- |
| **Basic error handling**          | `errors` (standard library) | Lightweight and sufficient for basic error wrapping and checking.                               |
| **Stack tracing**                 | `pkg/errors`                | Allows tracing error origins and is great for debugging.                                        |
| **Go <1.13 compatibility**        | `xerrors`                   | Adds modern error wrapping support for legacy Go versions.                                      |
| **Structured error metadata**     | `emperror` or `zerolog`     | Provides additional metadata and integrates well with structured logging solutions.             |
| **Error codes and microservices** | `go-micro/errors`           | Useful for distributed systems requiring structured error fields, like status and service name. |
| **Error classification**          | `errgo`                     | Allows organizing and propagating categorized errors in complex systems.                        |

### Summary

1. **Basic needs**: Stick with the **standard `errors` package** unless you need specific functionality.
2. **Stack tracing**: Use **`pkg/errors`** for stack tracing or debugging.
3. **Legacy compatibility**: Use **`xerrors`** if you’re working with an older Go version.
4. **Structured data**: For structured error metadata, try **`emperror`** or **`zerolog`**.
5. **Distributed systems**: In microservices, **`go-micro/errors`** offers fields like error codes and service names.
6. **Complex classification**: Use **`errgo`** if you need to classify errors in detail.

Each library has distinct features and strengths, so choosing the best one depends on the error-handling requirements of your specific project.
