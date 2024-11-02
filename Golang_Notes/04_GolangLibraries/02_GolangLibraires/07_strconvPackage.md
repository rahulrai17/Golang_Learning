## strconv package overview

The `strconv` package in Go provides essential functions for converting data between strings and various primitive data types such as integers, floating-point numbers, and booleans. This package is crucial when working with input/output, configurations, parsing command-line arguments, or handling data serialization.

### Overview of the `strconv` Package

Here is a summary of the functions provided by the `strconv` package in Go, along with a brief description for each:

| **Function**                                                             | **Description**                                                                               |
| ------------------------------------------------------------------------ | --------------------------------------------------------------------------------------------- |
| `Atoi(s string) (int, error)`                                            | Converts a string `s` to an integer.                                                          |
| `Itoa(i int) string`                                                     | Converts an integer `i` to a string.                                                          |
| `ParseInt(s string, base, bitSize int) (int64, error)`                   | Parses a string `s` to an integer in a specified `base` and `bitSize`.                        |
| `FormatInt(i int64, base int) string`                                    | Converts an `int64` to a string in the specified `base`.                                      |
| `ParseUint(s string, base, bitSize int) (uint64, error)`                 | Parses a string `s` to an unsigned integer.                                                   |
| `FormatUint(i uint64, base int) string`                                  | Converts a `uint64` to a string in the specified `base`.                                      |
| `ParseFloat(s string, bitSize int) (float64, error)`                     | Parses a string `s` to a floating-point number (float32 or float64).                          |
| `FormatFloat(f float64, fmt byte, prec, bitSize int) string`             | Converts a floating-point number to a string, with specified precision and format.            |
| `ParseBool(s string) (bool, error)`                                      | Parses a string `s` to a boolean value (`"true"`, `"false"`, `"1"`, `"0"`).                   |
| `FormatBool(b bool) string`                                              | Converts a boolean value `b` to a string (`"true"`, `"false"`).                               |
| `AppendInt(dst []byte, i int64, base int) []byte`                        | Appends the string form of an integer `i` in the specified `base` to a byte slice `dst`.      |
| `AppendUint(dst []byte, i uint64, base int) []byte`                      | Appends the string form of an unsigned integer to a byte slice `dst`.                         |
| `AppendFloat(dst []byte, f float64, fmt byte, prec, bitSize int) []byte` | Appends the string form of a float to a byte slice `dst` with specified format and precision. |
| `AppendBool(dst []byte, b bool) []byte`                                  | Appends the string form of a boolean to a byte slice `dst`.                                   |
| `Quote(s string) string`                                                 | Returns a double-quoted string literal representing the input string `s`.                     |
| `Unquote(s string) (string, error)`                                      | Converts a quoted string back to the unquoted form.                                           |
| `QuoteToASCII(s string) string`                                          | Returns a double-quoted ASCII string literal for the input string `s`.                        |
| `QuoteToGraphic(s string) string`                                        | Returns a double-quoted string literal with only graphic characters.                          |

This table provides a quick overview of the `strconv` package functions and their usage in converting between strings and primitive data types.

The main functionalities in `strconv` are:

- **String to primitive type conversions** (e.g., string to int, float, bool)
- **Primitive type to string conversions** (e.g., int, float, bool to string)
- **Handling complex base conversions** for numbers (e.g., binary, hexadecimal)
- **Dealing with errors** that occur during conversion

Let’s dive into the key functions provided by `strconv` with detailed explanations and examples.

---

### Key Functions in `strconv` Package

#### 1. **`strconv.Atoi(s string) (int, error)`**

- **Description:** Converts a string to an integer.
- **Use case:** Useful when converting numeric input (e.g., from a web form or CLI) into integers.
- **Example:**

  ```go
  package main

  import (
      "fmt"
      "strconv"
  )

  func main() {
      s := "123"
      i, err := strconv.Atoi(s)
      if err != nil {
          fmt.Println("Error:", err)
      } else {
          fmt.Println("Converted Integer:", i) // 123
      }
  }
  ```

#### 2. **`strconv.Itoa(i int) string`**

- **Description:** Converts an integer to a string.
- **Use case:** Useful for displaying numbers as strings or formatting them for text-based output.
- **Example:**

  ```go
  package main

  import (
      "fmt"
      "strconv"
  )

  func main() {
      i := 123
      s := strconv.Itoa(i)
      fmt.Println("Converted String:", s) // "123"
  }
  ```

#### 3. **`strconv.ParseInt(s string, base int, bitSize int) (int64, error)`**

- **Description:** Parses a string `s` to an integer of base `base` (e.g., base 2, 10, 16). The `bitSize` argument specifies the size of the integer (0, 8, 16, 32, or 64).
- **Use case:** Useful for converting strings with non-decimal bases like binary (base 2) or hexadecimal (base 16) to integers.
- **Example:**

  ```go
  package main

  import (
      "fmt"
      "strconv"
  )

  func main() {
      s := "1011"
      i, err := strconv.ParseInt(s, 2, 64) // Parse binary string
      if err != nil {
          fmt.Println("Error:", err)
      } else {
          fmt.Println("Converted Integer:", i) // 11
      }
  }
  ```

#### 4. **`strconv.FormatInt(i int64, base int) string`**

- **Description:** Converts an `int64` into a string representation in the specified `base` (2, 10, 16, etc.).
- **Use case:** Useful for formatting numbers in different bases, such as binary or hexadecimal.
- **Example:**

  ```go
  package main

  import (
      "fmt"
      "strconv"
  )

  func main() {
      i := int64(11)
      s := strconv.FormatInt(i, 2) // Convert to binary string
      fmt.Println("Binary String:", s) // "1011"
  }
  ```

#### 5. **`strconv.ParseUint(s string, base int, bitSize int) (uint64, error)`**

- **Description:** Parses a string `s` to an unsigned integer. Handles non-negative integers in the specified base.
- **Use case:** Parsing unsigned integer values, commonly found in IDs or address spaces.
- **Example:**

  ```go
  package main

  import (
      "fmt"
      "strconv"
  )

  func main() {
      s := "42"
      u, err := strconv.ParseUint(s, 10, 64) // Parse string to unsigned integer
      if err != nil {
          fmt.Println("Error:", err)
      } else {
          fmt.Println("Converted Unsigned Integer:", u) // 42
      }
  }
  ```

#### 6. **`strconv.FormatUint(i uint64, base int) string`**

- **Description:** Converts an unsigned integer `uint64` into a string representation in the specified `base`.
- **Use case:** Useful for displaying unsigned integers in binary, octal, or hexadecimal formats.
- **Example:**

  ```go
  package main

  import (
      "fmt"
      "strconv"
  )

  func main() {
      u := uint64(42)
      s := strconv.FormatUint(u, 16) // Convert to hexadecimal string
      fmt.Println("Hexadecimal String:", s) // "2a"
  }
  ```

#### 7. **`strconv.ParseFloat(s string, bitSize int) (float64, error)`**

- **Description:** Parses a string `s` to a floating-point number. The `bitSize` argument defines whether the result is `float32` or `float64`.
- **Use case:** Convert floating-point numbers (usually in decimal format) from strings.
- **Example:**

  ```go
  package main

  import (
      "fmt"
      "strconv"
  )

  func main() {
      s := "3.14159"
      f, err := strconv.ParseFloat(s, 64) // Parse string to float64
      if err != nil {
          fmt.Println("Error:", err)
      } else {
          fmt.Println("Converted Float:", f) // 3.14159
      }
  }
  ```

#### 8. **`strconv.FormatFloat(f float64, fmt byte, prec int, bitSize int) string`**

- **Description:** Converts a floating-point number to a string, formatted according to `fmt` ('e' for exponential, 'f' for decimal, etc.) and `prec` (precision).
- **Use case:** Useful for formatting floats for display in scientific or decimal notation.
- **Example:**

  ```go
  package main

  import (
      "fmt"
      "strconv"
  )

  func main() {
      f := 3.14159
      s := strconv.FormatFloat(f, 'f', 2, 64) // Convert to string with 2 decimal places
      fmt.Println("Formatted Float String:", s) // "3.14"
  }
  ```

#### 9. **`strconv.ParseBool(s string) (bool, error)`**

- **Description:** Parses a string `s` to a boolean value. Acceptable values are "true", "false", "1", "0".
- **Use case:** Converts string representations of boolean values from input sources.
- **Example:**

  ```go
  package main

  import (
      "fmt"
      "strconv"
  )

  func main() {
      s := "true"
      b, err := strconv.ParseBool(s)
      if err != nil {
          fmt.Println("Error:", err)
      } else {
          fmt.Println("Converted Boolean:", b) // true
      }
  }
  ```

#### 10. **`strconv.FormatBool(b bool) string`**

- **Description:** Converts a boolean `b` to its string representation ("true" or "false").
- **Use case:** When you need to convert boolean values to strings for output.
- **Example:**

  ```go
  package main

  import (
      "fmt"
      "strconv"
  )

  func main() {
      b := true
      s := strconv.FormatBool(b)
      fmt.Println("Formatted Boolean String:", s) // "true"
  }
  ```

#### 11. **`strconv.AppendInt(dst []byte, i int64, base int) []byte`**

- **Description:** Appends the string form of the integer `i` (in the given `base`) to the byte slice `dst`.
- **Use case:** Efficiently building strings by appending integer representations.
- **Example:**

  ```go
  package main

  import (
      "fmt"
      "strconv"
  )

  func main() {
      b := []byte("Number: ")
      b = strconv.AppendInt(b, 42, 10)
      fmt.Println(string(b)) // "Number: 42"
  }
  ```

#### 12. **`strconv.AppendFloat(dst []byte, f float64, fmt byte, prec int, bitSize int) []byte`**

- **Description:** Appends the string form of the float `f` to the byte slice `dst`.
- **Use case:** Used for efficiently constructing strings that include floating-point values.
- **Example:**

  ```go
  package main

  import (
      "fmt"
      "strconv"
  )

  func main() {
      b := []byte("Pi: ")
      b = strconv.AppendFloat(b, 3.14159, 'f', 2, 64)
      fmt.Println(string(b)) // "Pi: 3.14"
  }
  ```

---

### Other Notable Functions

- **`strconv.Quote(s string) string`**: Returns a double-quoted string literal representing `s`.

- **`strconv.Unquote(s string) (string, error)`**: Interprets a quoted string literal and returns the corresponding string value.

---

### Conclusion

The `strconv` package is essential for converting between strings and Go's primitive types (integers, floats, booleans), providing robust error handling and support for formatting. This package is highly useful in various scenarios such as reading user input, parsing data, formatting output, and dealing with numerical representations across different bases. Understanding when and how to use these functions is crucial in Go programming for managing conversions efficiently.
