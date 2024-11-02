## Basic Utilites

1. **log package**
   The `flag` package in Go provides a simple way to parse command-line flags. It allows you to define flags and retrieve their values in your Go applications. Below is an in-depth look at the `flag` package, including its key functions, types, and examples.

### Key Types

1. **Flag Variables**
   - The `flag` package provides types for different kinds of command-line flags, such as `string`, `int`, `bool`, etc.

### Key Functions

#### 1. **Defining Flags**

- **`flag.StringVar(p *string, name string, value string, usage string)`**

  - Defines a string flag with the specified name, default value, and usage information, storing the result in the `p` variable.

  ```go
  var name string
  flag.StringVar(&name, "name", "Guest", "Name of the user")
  ```

- **`flag.IntVar(p *int, name string, value int, usage string)`**

  - Defines an integer flag.

  ```go
  var age int
  flag.IntVar(&age, "age", 18, "Age of the user")
  ```

- **`flag.BoolVar(p *bool, name string, value bool, usage string)`**

  - Defines a boolean flag.

  ```go
  var verbose bool
  flag.BoolVar(&verbose, "verbose", false, "Enable verbose output")
  ```

- **`flag.Float64Var(p *float64, name string, value float64, usage string)`**
  - Defines a float flag.
  ```go
  var threshold float64
  flag.Float64Var(&threshold, "threshold", 0.5, "Threshold value")
  ```

#### 2. **Using Shortcuts**

- **`flag.String(name string, value string, usage string)`**

  - Defines a string flag and returns a pointer to the string variable.

  ```go
  name := flag.String("name", "Guest", "Name of the user")
  ```

- **`flag.Int(name string, value int, usage string)`**

  - Defines an integer flag and returns a pointer to the integer variable.

  ```go
  age := flag.Int("age", 18, "Age of the user")
  ```

- **`flag.Bool(name string, value bool, usage string)`**

  - Defines a boolean flag and returns a pointer to the boolean variable.

  ```go
  verbose := flag.Bool("verbose", false, "Enable verbose output")
  ```

- **`flag.Float64(name string, value float64, usage string)`**
  - Defines a float flag and returns a pointer to the float variable.
  ```go
  threshold := flag.Float64("threshold", 0.5, "Threshold value")
  ```

#### 3. **Parsing Flags**

- **`flag.Parse()`**
  - Parses the command-line flags from `os.Args[1:]`. This must be called after all flags are defined and before using the values.
  ```go
  flag.Parse()
  ```

#### 4. **Displaying Usage Information**

- **`flag.Usage`**
  - A variable of type `func()` that can be set to customize usage output.
  ```go
  flag.Usage = func() {
      fmt.Println("Custom usage message")
      flag.PrintDefaults()
  }
  ```

#### 5. **Retrieving Flag Values**

- The values of the flags can be accessed through the pointers returned when defining the flags.

### Example Usage

Here’s a complete example demonstrating how to use the `flag` package:

#### Example: Simple Command-Line Application

```go
package main

import (
    "flag"
    "fmt"
)

func main() {
    // Define flags
    var name string
    var age int
    var verbose bool

    flag.StringVar(&name, "name", "Guest", "Name of the user")
    flag.IntVar(&age, "age", 18, "Age of the user")
    flag.BoolVar(&verbose, "verbose", false, "Enable verbose output")

    // Parse the flags
    flag.Parse()

    // Output based on flags
    if verbose {
        fmt.Println("Verbose mode enabled.")
    }
    fmt.Printf("Hello, %s! You are %d years old.\n", name, age)
}
```

### Running the Example

You can run the above program from the command line as follows:

```sh
go run main.go -name="Alice" -age=30 -verbose
```

Output:

```
Verbose mode enabled.
Hello, Alice! You are 30 years old.
```

### Example: Custom Usage Message

You can customize the usage message as follows:

```go
package main

import (
    "flag"
    "fmt"
)

func main() {
    // Define flags
    var name string
    flag.StringVar(&name, "name", "Guest", "Name of the user")

    // Custom usage
    flag.Usage = func() {
        fmt.Println("Usage: myapp [options]")
        flag.PrintDefaults() // Print default usage
    }

    // Parse the flags
    flag.Parse()

    // Print the name
    fmt.Printf("Hello, %s!\n", name)
}
```

### Summary

The `flag` package in Go is a powerful and flexible tool for parsing command-line flags. It provides a simple API for defining flags, parsing them, and accessing their values. This is useful for creating command-line applications and utilities. The ability to customize usage messages adds to its versatility.

If you have any specific questions or need further examples on certain functions, feel free to ask!
