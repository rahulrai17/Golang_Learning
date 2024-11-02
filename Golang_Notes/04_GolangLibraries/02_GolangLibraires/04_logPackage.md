## Basic Utilites

1. **log package**
   The `log` package in Go provides a simple and flexible way to log messages. It offers various functionalities for logging at different levels of severity and allows you to configure the output format, destination, and more. Below is an in-depth look at the `log` package, including its key functions, constants, and examples.

### Key Types and Constants

#### 1. **Loggers**

- **`log.Logger`**
  - The main type provided by the `log` package. You create a logger instance that handles writing log messages to an `io.Writer`.

#### 2. **Log Flags**

Log flags determine what information is included in the log messages. You can combine multiple flags using a bitwise OR operation. Common flags include:

- **`log.Ldate`** - Date in the local time zone: `2009/01/23`
- **`log.Ltime`** - Time in the local time zone: `01:23:23`
- **`log.Lmicroseconds`** - Microsecond resolution: `01:23:23.123123`
- **`log.Llongfile`** - Full file name and line number: `/a/b/c/d.go:23`
- **`log.Lshortfile`** - Final file name element and line number: `d.go:23`
- **`log.LstdFlags`** - Standard logger flags: `Ldate | Ltime`

### Key Functions

#### 1. **Creating Loggers**

- **`log.New(w io.Writer, prefix string, flag int) *Logger`**
  - Creates a new logger with a specified output writer, prefix, and log flags.
  ```go
  logger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)
  logger.Println("This is an informational message.")
  ```

#### 2. **Logging Functions**

- **`log.Print(v ...interface{})`**

  - Logs a message at the default log level without any prefix.

  ```go
  log.Print("This is a simple log message.")
  ```

- **`log.Println(v ...interface{})`**

  - Logs a message with a newline at the end.

  ```go
  log.Println("This is a log message with a newline.")
  ```

- **`log.Printf(format string, v ...interface{})`**

  - Logs a formatted message.

  ```go
  log.Printf("Value: %d, String: %s", 42, "example")
  ```

- **`log.Fatal(v ...interface{})`**

  - Logs a message and then calls `os.Exit(1)`.

  ```go
  log.Fatal("This is a fatal error message.")
  ```

- **`log.Fatalln(v ...interface{})`**

  - Logs a message with a newline and then calls `os.Exit(1)`.

  ```go
  log.Fatalln("This is a fatal error message with a newline.")
  ```

- **`log.Fatalf(format string, v ...interface{})`**

  - Logs a formatted message and then calls `os.Exit(1)`.

  ```go
  log.Fatalf("Error occurred: %v", err)
  ```

- **`log.Panic(v ...interface{})`**

  - Logs a message and then panics.

  ```go
  log.Panic("This is a panic message.")
  ```

- **`log.Panicln(v ...interface{})`**

  - Logs a message with a newline and then panics.

  ```go
  log.Panicln("This is a panic message with a newline.")
  ```

- **`log.Panicf(format string, v ...interface{})`**
  - Logs a formatted message and then panics.
  ```go
  log.Panicf("Critical error: %v", err)
  ```

#### 3. **Setting Output**

- **`log.SetOutput(w io.Writer)`**

  - Sets the output destination for the default logger.

  ```go
  log.SetOutput(os.Stdout)  // Logs to standard output
  ```

- **`log.SetFlags(flag int)`**

  - Sets the log flags for the default logger.

  ```go
  log.SetFlags(log.LstdFlags | log.Lshortfile)
  ```

- **`log.SetPrefix(prefix string)`**
  - Sets the prefix for the default logger.
  ```go
  log.SetPrefix("DEBUG: ")
  ```

### Example Usage

Here are some complete examples demonstrating various functionalities of the `log` package:

#### Example 1: Basic Logging

```go
package main

import (
    "log"
    "os"
)

func main() {
    // Create a new logger
    logger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)

    // Log messages at different levels
    logger.Println("This is an informational message.")
    logger.Printf("Value: %d, String: %s\n", 42, "example")
}
```

#### Example 2: Logging Errors

```go
package main

import (
    "log"
    "os"
)

func main() {
    // Log a fatal error
    _, err := os.Open("nonexistentfile.txt")
    if err != nil {
        log.Fatalf("Error opening file: %v\n", err)  // Logs the error and exits
    }
}
```

#### Example 3: Custom Logger

```go
package main

import (
    "log"
    "os"
)

func main() {
    // Create a custom logger
    errorLog := log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

    // Log an error message
    errorLog.Println("This is an error message.")
}
```

#### Example 4: Using Panic and Fatal

```go
package main

import (
    "log"
)

func main() {
    defer func() {
        if r := recover(); r != nil {
            log.Println("Recovered from panic:", r)
        }
    }()

    log.Println("Normal log message.")
    log.Panic("This will panic and log a message.")
    log.Println("This line will not be executed.")  // This won't be reached
}
```

### Summary

The `log` package in Go is a powerful and flexible logging tool that allows developers to log messages at various levels of severity and customize the logging format and output. It is essential for debugging and monitoring applications. Whether you need simple logs or structured logging with prefixes and flags, the `log` package provides the necessary functionality.

If you have any specific questions or need further examples, feel free to ask!
