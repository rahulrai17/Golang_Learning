## Taking Inputs in Golang

In Go (Golang), there are multiple ways to take input depending on the use case, such as reading from the console, files, or external sources. Let's explore these methods along with examples and explanations.

### 1. **Reading Input from the Console using `fmt` package**

This is the most common way to take input in Go, primarily used for reading input from users in interactive programs.

#### 1) `fmt.Scan()`

- **Use case**: Suitable for reading simple input like integers, strings, or floats from the console.
- **Syntax**: `fmt.Scan(&var1, &var2, ...)`

**Example**:

```go
package main

import (
    "fmt"
)

func main() {
    var name string
    var age int

    fmt.Print("Enter your name: ")
    fmt.Scan(&name) // Scans one word

    fmt.Print("Enter your age: ")
    fmt.Scan(&age) // Scans an integer

    fmt.Printf("Name: %s, Age: %d\n", name, age)
}
```

- **Explanation**:
  - `fmt.Scan(&name)` takes one word of input and stores it in the `name` variable.
  - `fmt.Scan(&age)` takes an integer input and stores it in the `age` variable.

#### 2) `fmt.Scanln()`

- **Use case**: Reads a line of input until a newline is encountered. Useful when you want to ensure you read input until the end of the line.
- **Syntax**: `fmt.Scanln(&var1, &var2, ...)`

**Example**:

```go
package main

import (
    "fmt"
)

func main() {
    var name string

    fmt.Print("Enter your full name: ")
    fmt.Scanln(&name) // Scans input until a newline

    fmt.Printf("Full name: %s\n", name)
}
```

- **Explanation**:
  - `fmt.Scanln()` reads until the user presses Enter. If the input includes spaces, only the first word is captured in `name`.

#### 3) `fmt.Scanf()`

- **Use case**: Used for formatted input where you need to specify the format, similar to `scanf()` in C.
- **Syntax**: `fmt.Scanf(format, &var1, &var2, ...)`

**Example**:

```go
package main

import (
    "fmt"
)

func main() {
    var name string
    var age int

    fmt.Print("Enter your name and age: ")
    fmt.Scanf("%s %d", &name, &age) // Scans formatted input

    fmt.Printf("Name: %s, Age: %d\n", name, age)
}
```

- **Explanation**:
  - `fmt.Scanf("%s %d", &name, &age)` reads a string and an integer from the input, expecting the user to enter them in the specified format.

---

### 2. **Using `bufio` for Buffered Input**

The `bufio` package provides buffered I/O which allows efficient reading of large inputs or reading from multiple sources (stdin, files, etc.). It is especially useful when you need more control over input, like reading entire lines or handling large text inputs.

#### 1) `bufio.NewReader()`

- **Use case**: Reading an entire line of input including spaces or reading large text inputs efficiently.
- **Syntax**: `reader.ReadString('\n')`

**Example**:

```go
package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter a sentence: ")

    sentence, _ := reader.ReadString('\n') // Reads until newline
    fmt.Println("You entered:", sentence)
}
```

- **Explanation**:
  - `bufio.NewReader(os.Stdin)` creates a new reader for standard input.
  - `reader.ReadString('\n')` reads everything the user types until they press Enter (newline character).

#### 2) `bufio.Scanner()`

- **Use case**: For token-based input or reading line by line from stdin or a file.
- **Syntax**: `scanner.Scan()` followed by `scanner.Text()`

**Example**:

```go
package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    fmt.Print("Enter text: ")

    scanner.Scan() // Reads the next token
    text := scanner.Text() // Retrieves the text

    fmt.Println("You entered:", text)
}
```

- **Explanation**:
  - `scanner.Scan()` reads the input until a newline or whitespace by default.
  - `scanner.Text()` retrieves the scanned token (input).

#### Differences between `bufio.NewReader()` and `bufio.Scanner()`

- `bufio.NewReader()` is more flexible for reading specific delimiters or handling custom input requirements.
- `bufio.Scanner()` is more straightforward and commonly used for line-by-line or tokenized input but less flexible (it doesn’t handle arbitrary delimiters well).

---

### 3. **Using `os` Package for Command-Line Arguments**

Command-line arguments allow your program to accept inputs when it is executed. This is useful for automation and scripting purposes.

#### 1) `os.Args`

- **Use case**: When you want to capture inputs passed as arguments when running the Go program from the command line.
- **Syntax**: `os.Args`

**Example**:

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    args := os.Args // Captures command-line arguments

    if len(args) > 1 {
        fmt.Println("Arguments:", args[1:])
    } else {
        fmt.Println("No arguments provided")
    }
}
```

- **Explanation**:
  - `os.Args` returns a slice of strings representing command-line arguments. `os.Args[0]` is the program name, and the subsequent elements are the arguments.
  - This method is useful for writing command-line tools or scripts.

---

### 4. **Reading from Files using `os` or `ioutil`**

Input can also come from files. Reading from a file allows your program to process large datasets or structured inputs like JSON, CSV, or plain text.

#### 1) `os.Open()`

- **Use case**: When you need to read data from a file line by line or in chunks.
- **Syntax**: `os.Open(filename)`

**Example**:

```go
package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    file, err := os.Open("input.txt")
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        fmt.Println(scanner.Text()) // Print each line
    }
}
```

- **Explanation**:
  - `os.Open("input.txt")` opens the file for reading.
  - `scanner.Scan()` reads the file line by line, and `scanner.Text()` returns the current line as a string.

---

### 5. **Reading from Environment Variables**

Environment variables are another source of input for Go programs, commonly used in server configurations.

#### 1) `os.Getenv()`

- **Use case**: Reading configuration or sensitive data like API keys from environment variables.
- **Syntax**: `os.Getenv("VAR_NAME")`

**Example**:

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    key := os.Getenv("API_KEY") // Reads API_KEY environment variable

    if key != "" {
        fmt.Println("API Key:", key)
    } else {
        fmt.Println("API Key not set")
    }
}
```

- **Explanation**:
  - `os.Getenv("API_KEY")` retrieves the value of the `API_KEY` environment variable, commonly used in cloud or containerized applications.

---

### 6. **Using Third-party Libraries for Advanced Input**

There are third-party libraries in Go for specialized input handling, such as:

- **Cobra**: For building command-line applications with more complex flag handling and subcommands.
- **PromptUI**: For interactive command-line interfaces (CLI) with menus and prompts.

---

### Summary Table

| Input Method        | Use Case                                   | Example Use Case                                 |
| ------------------- | ------------------------------------------ | ------------------------------------------------ |
| `fmt.Scan()`        | Simple input of basic types                | Reading a single integer or word                 |
| `fmt.Scanln()`      | Input with space handling                  | Reading a full sentence                          |
| `bufio.NewReader()` | Efficient large input or custom delimiters | Reading multi-line text input                    |
| `bufio.Scanner()`   | Line-by-line or token-based input          | Reading input line by line (e.g., file or stdin) |
| `os.Args`           | Command-line arguments                     | Passing arguments to the program                 |
| `os.Open()`         | File input                                 | Processing file data                             |
| `os.Getenv()`       | Reading environment variables              | Configuration or sensitive data like API keys    |

Each input method is useful depending on the context of the program, such as interactive CLI tools, server configurations, file processing, or tokenized inputs.
