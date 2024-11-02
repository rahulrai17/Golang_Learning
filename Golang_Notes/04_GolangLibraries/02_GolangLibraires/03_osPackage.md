## Basic Utilites

1. **os package :**
   The `os` package in Go provides a platform-independent interface to operating system functionality, such as working with files and directories, handling environment variables, and managing processes. Below is a detailed overview of the `os` package, including its key functions, constants, and practical examples.

### Key Functions and Types

#### 1. **File and Directory Manipulation**

- **`os.Open(name string) (*os.File, error)`**

  - Opens a file for reading.

  ```go
  file, err := os.Open("file.txt")
  if err != nil {
      log.Fatal(err)
  }
  defer file.Close()
  ```

- **`os.Create(name string) (*os.File, error)`**

  - Creates a new file or truncates an existing one.

  ```go
  file, err := os.Create("newfile.txt")
  if err != nil {
      log.Fatal(err)
  }
  defer file.Close()
  ```

- **`os.Remove(name string) error`**

  - Deletes a named file or empty directory.

  ```go
  err := os.Remove("newfile.txt")
  if err != nil {
      log.Fatal(err)
  }
  ```

- **`os.Rename(oldpath, newpath string) error`**

  - Renames (moves) a file or directory.

  ```go
  err := os.Rename("oldname.txt", "newname.txt")
  if err != nil {
      log.Fatal(err)
  }
  ```

- **`os.Mkdir(name string, perm os.FileMode) error`**

  - Creates a new directory with the specified permissions.

  ```go
  err := os.Mkdir("newdir", 0755)
  if err != nil {
      log.Fatal(err)
  }
  ```

- **`os.MkdirAll(path string, perm os.FileMode) error`**

  - Creates a directory and any necessary parents.

  ```go
  err := os.MkdirAll("parent/child", 0755)
  if err != nil {
      log.Fatal(err)
  }
  ```

- **`os.Stat(name string) (os.FileInfo, error)`**
  - Returns a description of the named file.
  ```go
  info, err := os.Stat("file.txt")
  if err != nil {
      log.Fatal(err)
  }
  fmt.Println(info.Size())  // Outputs the size of the file
  ```

#### 2. **Working with Environment Variables**

- **`os.Getenv(key string) string`**

  - Gets the value of the environment variable `key`.

  ```go
  home := os.Getenv("HOME")
  fmt.Println("Home directory:", home)
  ```

- **`os.Setenv(key, value string) error`**

  - Sets the value of the environment variable `key`.

  ```go
  err := os.Setenv("MY_VAR", "some_value")
  if err != nil {
      log.Fatal(err)
  }
  ```

- **`os.Unsetenv(key string) error`**

  - Unsets the environment variable `key`.

  ```go
  err := os.Unsetenv("MY_VAR")
  if err != nil {
      log.Fatal(err)
  }
  ```

- **`os.Environ() []string`**
  - Returns a slice of strings representing the environment, in the form `key=value`.
  ```go
  for _, e := range os.Environ() {
      fmt.Println(e)
  }
  ```

#### 3. **File and Directory Operations**

- **`os.ReadDir(name string) ([]os.DirEntry, error)`**

  - Reads the contents of a directory.

  ```go
  entries, err := os.ReadDir(".")
  if err != nil {
      log.Fatal(err)
  }
  for _, entry := range entries {
      fmt.Println(entry.Name())
  }
  ```

- **`os.Chdir(dir string) error`**

  - Changes the current working directory.

  ```go
  err := os.Chdir("newdir")
  if err != nil {
      log.Fatal(err)
  }
  ```

- **`os.Getwd() (string, error)`**
  - Returns the current working directory.
  ```go
  dir, err := os.Getwd()
  if err != nil {
      log.Fatal(err)
  }
  fmt.Println("Current Directory:", dir)
  ```

#### 4. **Process Management**

- **`os.Exit(code int)`**

  - Exits the program with the given status code.

  ```go
  os.Exit(1)  // Exits with an error status
  ```

- **`os.StartProcess(name string, argv []string, attr *os.ProcAttr) (*os.Process, error)`**
  - Starts a new process.
  ```go
  proc, err := os.StartProcess("mycommand", []string{"mycommand", "arg1"}, nil)
  if err != nil {
      log.Fatal(err)
  }
  ```

#### 5. **File Mode and File Info**

- **`os.FileMode`**

  - Represents file permissions.

  ```go
  perm := os.FileMode(0755)  // Read, write, execute for owner; read and execute for group and others
  ```

- **`os.FileInfo`**
  - Interface that provides file information (size, permissions, etc.).
  ```go
  info, _ := os.Stat("file.txt")
  fmt.Println(info.Name())    // File name
  fmt.Println(info.IsDir())   // Is it a directory?
  ```

### Example Usage

Here are some complete examples demonstrating various `os` package functions:

#### Example 1: Creating and Reading a File

```go
package main

import (
    "fmt"
    "log"
    "os"
)

func main() {
    // Create a new file
    file, err := os.Create("example.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    // Write to the file
    _, err = file.WriteString("Hello, World!\n")
    if err != nil {
        log.Fatal(err)
    }

    // Read the file
    content, err := os.ReadFile("example.txt")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(string(content))  // Output: Hello, World!
}
```

#### Example 2: Working with Environment Variables

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    // Set an environment variable
    os.Setenv("MY_VAR", "Hello, Environment!")

    // Get the environment variable
    value := os.Getenv("MY_VAR")
    fmt.Println("MY_VAR:", value)  // Output: MY_VAR: Hello, Environment!

    // List all environment variables
    for _, e := range os.Environ() {
        fmt.Println(e)
    }
}
```

#### Example 3: Directory Operations

```go
package main

import (
    "fmt"
    "log"
    "os"
)

func main() {
    // Create a directory
    err := os.Mkdir("mydir", 0755)
    if err != nil {
        log.Fatal(err)
    }

    // Change current directory
    err = os.Chdir("mydir")
    if err != nil {
        log.Fatal(err)
    }

    // Print current directory
    dir, _ := os.Getwd()
    fmt.Println("Current Directory:", dir)  // Output: Current Directory: /path/to/mydir

    // List contents of the parent directory
    entries, _ := os.ReadDir("..")
    for _, entry := range entries {
        fmt.Println(entry.Name())
    }
}
```

#### Example 4: Process Management

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    fmt.Println("Starting the process...")
    // Exit the program with a status code
    os.Exit(0)  // This line will exit the program immediately
}
```

### Summary

The `os` package is essential for performing operating system-level operations in Go. It provides a wide range of functionality, from file and directory manipulation to process management and environment variable handling. These features are crucial for building robust applications that need to interact with the underlying operating system.

If you have any specific questions or need further examples on certain functions, feel free to ask!
