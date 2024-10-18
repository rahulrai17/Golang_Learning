## Go Keywords

### List of Go Keywords

Here are the 25 reserved keywords in Go:

| Keyword       | Description                                                                |
| ------------- | -------------------------------------------------------------------------- |
| `break`       | Exits a loop or switch statement.                                          |
| `case`        | Specifies a case in a switch statement.                                    |
| `chan`        | Defines a channel type for goroutines.                                     |
| `const`       | Declares a constant value.                                                 |
| `continue`    | Skips the current iteration of a loop and proceeds to the next iteration.  |
| `default`     | Specifies the default case in a switch statement.                          |
| `defer`       | Delays the execution of a function until the surrounding function returns. |
| `else`        | Executes an alternative branch in an if statement.                         |
| `fallthrough` | Allows execution to continue to the next case in a switch statement.       |
| `for`         | Starts a loop or an iteration construct.                                   |
| `func`        | Declares a function.                                                       |
| `go`          | Starts a new goroutine (concurrent function execution).                    |
| `goto`        | Jumps to a labeled statement.                                              |
| `if`          | Executes a block of code based on a condition.                             |
| `import`      | Includes other packages in the current package.                            |
| `interface`   | Defines an interface, specifying method signatures.                        |
| `map`         | Defines a map type (a collection of key-value pairs).                      |
| `package`     | Declares a package name.                                                   |
| `range`       | Iterates over elements in a variety of data structures.                    |
| `return`      | Exits a function and optionally returns a value.                           |
| `select`      | Waits on multiple channel operations.                                      |
| `struct`      | Defines a composite data type (structure).                                 |
| `switch`      | Executes one block of code among multiple options based on a value.        |
| `type`        | Defines a new type or a type alias.                                        |
| `var`         | Declares a variable.                                                       |

---

### Detailed Explanation of Each Keyword

#### 1. **`break`**

- **Purpose**: Exits the nearest enclosing loop or switch statement.
- **Example**:
  ```go
  for i := 0; i < 10; i++ {
      if i == 5 {
          break // Exit the loop when i is 5
      }
      fmt.Println(i)
  }
  ```

#### 2. **`case`**

- **Purpose**: Specifies a branch in a switch statement.
- **Example**:
  ```go
  switch x := 2; x {
  case 1:
      fmt.Println("One")
  case 2:
      fmt.Println("Two")
  default:
      fmt.Println("Unknown")
  }
  ```

#### 3. **`chan`**

### **Purpose**

The `chan` keyword is used to create channels, which are essential for communication between goroutines in Go. Channels allow goroutines to send and receive messages, helping synchronize their execution and manage data flow.

### **Types of Channels**

- **Unbuffered Channels**: The sender and receiver must be ready at the same time. This provides synchronization but can lead to blocking if one side is not ready.
- **Buffered Channels**: Can hold a fixed number of values. Sending to a buffered channel only blocks when the buffer is full, and receiving blocks when the buffer is empty.

### **Syntax**

```go
// Create an unbuffered channel
ch := make(chan int)

// Create a buffered channel with a capacity of 2
chBuffered := make(chan int, 2)
```

### **Example**

Here’s a simple example demonstrating how to use channels:

```go
package main

import (
    "fmt"
)

func sendData(ch chan int) {
    for i := 1; i <= 5; i++ {
        ch <- i // Send data to the channel
    }
    close(ch) // Close the channel when done
}

func main() {
    ch := make(chan int) // Unbuffered channel
    go sendData(ch) // Start the goroutine

    // Receive data from the channel
    for value := range ch {
        fmt.Println("Received:", value)
    }
}
```

### **Real-World Use Case**

Channels are commonly used in applications where concurrent processing is required, such as:

- **Web servers**: Handling requests concurrently using goroutines, with channels managing communication between handlers.
- **Data processing pipelines**: Where data flows through various processing stages, with each stage running in its own goroutine.

---

#### 4. **`const`**

- **Purpose**: Declares a constant, a value that cannot be changed after it's declared.
- **Example**:
  ```go
  const Pi = 3.14
  ```

#### 5. **`continue`**

- **Purpose**: Skips the current iteration of a loop and moves to the next iteration.
- **Example**:
  ```go
  for i := 0; i < 10; i++ {
      if i%2 == 0 {
          continue // Skip even numbers
      }
      fmt.Println(i) // Prints only odd numbers
  }
  ```

#### 6. **`default`**

- **Purpose**: Specifies the default case in a switch statement.
- **Example**:
  ```go
  switch x := 2; x {
  case 1:
      fmt.Println("One")
  default:
      fmt.Println("Not One")
  }
  ```

#### 7. **`defer`**

### **Purpose**

The `defer` keyword is used to delay the execution of a function until the surrounding function returns. This is useful for resource management, such as closing files, unlocking mutexes, or printing debug information at the end of a function.

### **Syntax**

```go
defer functionName(parameters)
```

### **Example**

```go
package main

import (
    "fmt"
)

func main() {
    defer fmt.Println("This will be printed last.")
    fmt.Println("This will be printed first.")
}
```

### **Real-World Use Case**

`defer` is often used for resource cleanup:

- **File Handling**: Ensuring that files are closed properly after their operations are completed.

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    file, err := os.Open("test.txt")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer file.Close() // Ensures file is closed when the function exits

    // Perform file operations...
}
```

- **Database Connections**: Closing connections after querying the database.

---

#### 8. **`else`**

- **Purpose**: Provides an alternative block of code to execute if the condition in an if statement is false.
- **Example**:
  ```go
  if x > 10 {
      fmt.Println("Greater than 10")
  } else {
      fmt.Println("10 or less")
  }
  ```

#### 9. **`fallthrough`**

### **Purpose**

The `fallthrough` keyword is used within a `switch` statement to allow execution to continue to the next case, even if the next case’s condition is not true.

### **Syntax**

```go
switch variable {
case value1:
    // code for case 1
    fallthrough
case value2:
    // code for case 2
}
```

### **Example**

```go
package main

import (
    "fmt"
)

func main() {
    num := 2
    switch num {
    case 1:
        fmt.Println("One")
    case 2:
        fmt.Println("Two")
        fallthrough
    case 3:
        fmt.Println("Three") // This will also be executed
    default:
        fmt.Println("Default case")
    }
}
```

### **Real-World Use Case**

`fallthrough` can be useful when multiple cases should execute the same code:

- **Status Codes**: Handling multiple HTTP status codes that share the same response logic.

```go
func getHTTPStatusMessage(code int) string {
    switch code {
    case 200:
        return "OK"
    case 404:
        return "Not Found"
    case 500:
        fallthrough // continue to 501
    case 501:
        return "Not Implemented"
    default:
        return "Unknown Status"
    }
}
```

#### 10. **`for`**

- **Purpose**: Starts a loop construct; the only looping construct in Go.
- **Example**:
  ```go
  for i := 0; i < 5; i++ {
      fmt.Println(i) // Prints numbers from 0 to 4
  }
  ```

#### 11. **`func`**

- **Purpose**: Declares a function.
- **Example**:
  ```go
  func greet(name string) {
      fmt.Println("Hello, " + name)
  }
  ```

#### 12. **`go`**

- **Purpose**: Starts a new goroutine, allowing concurrent execution.
- **Example**:
  ```go
  go func() {
      fmt.Println("This runs concurrently")
  }()
  ```

#### 13. **`goto`**

### **Purpose**

The `goto` keyword is used to jump to a labeled statement within the same function. While its use can lead to less readable code, it can be helpful in certain scenarios for handling complex flow control.

### **Syntax**

```go
labelName:
    // some code
goto labelName // Jumps to labelName
```

### **Example**

```go
package main

import (
    "fmt"
)

func main() {
    i := 0
    Loop:
        if i < 5 {
            fmt.Println(i)
            i++
            goto Loop // Jump back to Loop label
        }
}
```

### **Real-World Use Case**

`goto` might be useful in:

- **Error Handling**: Jumping to cleanup code in case of an error without deeply nested if statements.

```go
package main

import (
    "fmt"
)

func main() {
    defer fmt.Println("Cleanup done.")

    for i := 0; i < 3; i++ {
        if i == 2 {
            goto Error // Jump to error handling
        }
        fmt.Println(i)
    }

Error:
    fmt.Println("An error occurred.")
}
```

### **Caution**

While `goto` can simplify error handling in some cases, it's often discouraged because it can lead to spaghetti code. Prefer structured error handling with functions, channels, or other control structures when possible.

---

#### 14. **`if`**

- **Purpose**: Executes a block of code based on a condition.
- **Example**:
  ```go
  if x > 10 {
      fmt.Println("x is greater than 10")
  }
  ```

#### 15. **`import`**

- **Purpose**: Includes other packages in the current package.
- **Example**:
  ```go
  import "fmt" // Import the fmt package for formatted I/O
  ```

#### 16. **`interface`**

- **Purpose**: Defines a contract for types, specifying method signatures.
- **Example**:
  ```go
  type Shape interface {
      Area() float64
  }
  ```

#### 17. **`map`**

- **Purpose**: Defines a map type, a collection of key-value pairs.
- **Example**:
  ```go
  myMap := make(map[string]int)
  myMap["Alice"] = 25
  ```

#### 18. **`package`**

- **Purpose**: Declares the package name for the current file.
- **Example**:
  ```go
  package main // The main package, entry point of the program
  ```

#### 19. **`range`**

- **Purpose**: Iterates over elements in a variety of data structures, including slices and maps.
- **Example**:
  ```go
  for i, v := range []int{1, 2, 3} {
      fmt.Println(i, v) // Prints index and value
  }
  ```

#### 20. **`return`**

- **Purpose**: Exits a function and optionally returns a value.
- **Example**:
  ```go
  func add(a int, b int) int {
      return a + b // Returns the sum of a and b
  }
  ```

#### 21. **`select`**

### **Purpose**

The `select` statement provides another way to handle multiple channel operations. It allows a goroutine to wait on multiple channels and proceed with the operation that is ready first. This is particularly useful for handling timeouts and multiple communication paths.

### **Syntax**

```go
select {
case msg1 := <-ch1:
    // Handle message from channel 1
case msg2 := <-ch2:
    // Handle message from channel 2
case <-time.After(time.Second):
    // Handle timeout
}
```

### **Example**

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    ch1 := make(chan string)
    ch2 := make(chan string)

    go func() {
        time.Sleep(2 * time.Second)
        ch1 <- "Message from channel 1"
    }()

    go func() {
        time.Sleep(1 * time.Second)
        ch2 <- "Message from channel 2"
    }()

    select {
    case msg := <-ch1:
        fmt.Println(msg)
    case msg := <-ch2:
        fmt.Println(msg)
    case <-time.After(3 * time.Second):
        fmt.Println("Timeout!")
    }
}
```

### **Real-World Use Case**

`select` is invaluable in concurrent programming:

- **Timeouts**: Implementing timeouts for goroutines that might hang waiting for a channel.
- **Multiplexing**: Listening on multiple channels and acting upon whichever channel receives data first, commonly seen in server applications that handle multiple clients.

---

#### 22. **`struct`**

- **Purpose**: Defines a composite data type, combining multiple fields.
- **Example**:
  ```go
  type Person struct {
      Name string
      Age  int
  }
  ```

#### 23. **`switch`**

- **Purpose**: Executes one block of code among multiple options based on a value.
- **Example**:
  ```go
  switch day := 5; day {
  case 1:
      fmt.Println("Monday")
  case 5:
      fmt.Println("Friday")
  default:
      fmt.Println("Another day")
  }
  ```

#### 24. **`type`**

- **Purpose**: Defines a new type or

a type alias.

- **Example**:
  ```go
  type Celsius float64
  ```

#### 25. **`var`**

- **Purpose**: Declares a variable.
- **Example**:
  ```go
  var age int = 30 // Declares a variable age of type int
  ```

---
