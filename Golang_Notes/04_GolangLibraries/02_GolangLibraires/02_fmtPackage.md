## Basic Utilites

1. **fmt package :**
   The `fmt` package in Go provides formatted I/O with functions similar to C's `printf` and `scanf`. It's widely used for outputting formatted text and reading input from the console. Below is a comprehensive overview of the `fmt` package, including its key functions, formatting verbs, and examples.

### Key Functions

#### 1. **Printing Functions**

- **`Print`**

  - Outputs the arguments to standard output without a newline.

  ```go
  fmt.Print("Hello, World!")
  ```

- **`Println`**

  - Outputs the arguments to standard output with a newline.

  ```go
  fmt.Println("Hello, World!")  // Output: Hello, World!
  ```

- **`Printf`**

  - Outputs formatted text based on format specifiers.

  ```go
  fmt.Printf("Hello, %s!\n", "World")  // Output: Hello, World!
  ```

- **`Fprint`, `Fprintln`, `Fprintf`**
  - Print to a specified `io.Writer` (like a file).
  ```go
  file, _ := os.Create("output.txt")
  defer file.Close()
  fmt.Fprintln(file, "Hello, World!")  // Writes to the file
  ```

#### 2. **Scanning Functions**

- **`Scan`**

  - Reads input from standard input.

  ```go
  var name string
  fmt.Print("Enter your name: ")
  fmt.Scan(&name)  // User types input
  ```

- **`Scanln`**

  - Reads input from standard input until a newline.

  ```go
  var name string
  fmt.Print("Enter your name: ")
  fmt.Scanln(&name)  // Stops reading at newline
  ```

- **`Scanf`**
  - Reads formatted input.
  ```go
  var name string
  var age int
  fmt.Print("Enter your name and age: ")
  fmt.Scanf("%s %d", &name, &age)  // User types input like "Alice 30"
  ```

#### 3. **Formatting Functions**

- **`Sprintf`**

  - Returns a formatted string instead of printing it.

  ```go
  formatted := fmt.Sprintf("Hello, %s!", "World")
  fmt.Println(formatted)  // Output: Hello, World!
  ```

- **`Sprint`, `Sprintln`, `Sprintf`**
  - Similar to `Print`, but return the formatted string.
  ```go
  result := fmt.Sprintf("Number: %d", 42)
  fmt.Println(result)  // Output: Number: 42
  ```

### Format Specifiers

The `fmt` package uses format verbs to format data. Here are some commonly used verbs:

| Verb  | Description                  | Example                     |
| ----- | ---------------------------- | --------------------------- |
| `%v`  | Default format for the value | `fmt.Printf("%v", 42)`      |
| `%#v` | Go-syntax representation     | `fmt.Printf("%#v", 42)`     |
| `%T`  | Type of the value            | `fmt.Printf("%T", 42)`      |
| `%d`  | Decimal integer              | `fmt.Printf("%d", 42)`      |
| `%s`  | String                       | `fmt.Printf("%s", "Hello")` |
| `%f`  | Floating-point number        | `fmt.Printf("%f", 3.14)`    |
| `%t`  | Boolean                      | `fmt.Printf("%t", true)`    |
| `%x`  | Hexadecimal integer          | `fmt.Printf("%x", 255)`     |
| `%p`  | Pointer                      | `fmt.Printf("%p", &a)`      |

### Example Usage

Here are some complete examples demonstrating the use of various `fmt` functions:

#### Example 1: Printing with Formatting

```go
package main

import (
    "fmt"
)

func main() {
    name := "Alice"
    age := 30
    fmt.Printf("Name: %s, Age: %d\n", name, age)
}
```

#### Example 2: Reading Input

```go
package main

import (
    "fmt"
)

func main() {
    var name string
    var age int
    fmt.Print("Enter your name and age: ")
    fmt.Scanf("%s %d", &name, &age)
    fmt.Printf("Hello, %s! You are %d years old.\n", name, age)
}
```

#### Example 3: Writing to a File

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    file, err := os.Create("output.txt")
    if err != nil {
        fmt.Println("Error creating file:", err)
        return
    }
    defer file.Close()

    fmt.Fprintln(file, "Hello, World!")  // Writes to the file
}
```

#### Example 4: Formatting a Complex Structure

```go
package main

import (
    "fmt"
)

type Person struct {
    Name string
    Age  int
}

func main() {
    p := Person{"Bob", 25}
    fmt.Printf("Person: %+v\n", p)  // Output: Person: {Name:Bob Age:25}
}
```

### Summary

The `fmt` package is essential for handling formatted I/O in Go. It provides a rich set of functions for printing and reading formatted text, which is useful for both debugging and user interaction. You can leverage formatting verbs to control how data is presented, making it a powerful tool in any Go developer's toolkit.

If you have any specific questions or need further examples, feel free to ask!
