## Operators In Golang

In Go, operators are symbols that are used to perform operations on variables and values. Go has various types of operators: arithmetic, comparison, logical, bitwise, assignment, and more. Let's go through each category, explain the operators, their precedence, and how they work with examples.

---

## **Types of Operators in Go**

### 1. **Arithmetic Operators**

These operators are used to perform basic mathematical operations like addition, subtraction, multiplication, etc.

| Operator | Description         | Example |
| -------- | ------------------- | ------- |
| `+`      | Addition            | `x + y` |
| `-`      | Subtraction         | `x - y` |
| `*`      | Multiplication      | `x * y` |
| `/`      | Division            | `x / y` |
| `%`      | Modulus (remainder) | `x % y` |
| `++`     | Increment (by 1)    | `x++`   |
| `--`     | Decrement (by 1)    | `x--`   |

**Example**:

```go
package main

import "fmt"

func main() {
    a := 10
    b := 3

    fmt.Println("Addition:", a + b)        // 13
    fmt.Println("Subtraction:", a - b)     // 7
    fmt.Println("Multiplication:", a * b)  // 30
    fmt.Println("Division:", a / b)        // 3
    fmt.Println("Modulus:", a % b)         // 1
    a++
    fmt.Println("Increment:", a)           // 11
    b--
    fmt.Println("Decrement:", b)           // 2
}
```

### 2. **Comparison Operators**

Comparison operators are used to compare two values and return a boolean (`true` or `false`).

| Operator | Description           | Example  |
| -------- | --------------------- | -------- |
| `==`     | Equal to              | `x == y` |
| `!=`     | Not equal to          | `x != y` |
| `>`      | Greater than          | `x > y`  |
| `<`      | Less than             | `x < y`  |
| `>=`     | Greater than or equal | `x >= y` |
| `<=`     | Less than or equal    | `x <= y` |

**Example**:

```go
package main

import "fmt"

func main() {
    x := 5
    y := 8

    fmt.Println("x == y:", x == y)  // false
    fmt.Println("x != y:", x != y)  // true
    fmt.Println("x > y:", x > y)    // false
    fmt.Println("x < y:", x < y)    // true
    fmt.Println("x >= y:", x >= y)  // false
    fmt.Println("x <= y:", x <= y)  // true
}
```

### 3. **Logical Operators**

Logical operators are used to combine conditional statements.

| Operator | Description | Example  |
| -------- | ----------- | -------- |
| `&&`     | Logical AND | `x && y` |
| `        |             | `        |
| `!`      | Logical NOT | `!x`     |

**Example**:

```go
package main

import "fmt"

func main() {
    a := true
    b := false

    fmt.Println("a && b:", a && b) // false
    fmt.Println("a || b:", a || b) // true
    fmt.Println("!a:", !a)         // false
}
```

### 4. **Bitwise Operators**

Bitwise operators are used to perform operations at the binary level.

| Operator | Description | Example    |
| -------- | ----------- | ---------- |
| `&`      | Bitwise AND | `x & y`    |
| `        | `           | Bitwise OR |
| `^`      | Bitwise XOR | `x ^ y`    |
| `<<`     | Left shift  | `x << 1`   |
| `>>`     | Right shift | `x >> 1`   |

**Example**:

```go
package main

import "fmt"

func main() {
    x := 6  // 110 in binary
    y := 3  // 011 in binary

    fmt.Println("x & y:", x & y)  // Bitwise AND: 010 -> 2
    fmt.Println("x | y:", x | y)  // Bitwise OR:  111 -> 7
    fmt.Println("x ^ y:", x ^ y)  // Bitwise XOR: 101 -> 5
    fmt.Println("x << 1:", x << 1) // Left shift: 1100 -> 12
    fmt.Println("x >> 1:", x >> 1) // Right shift: 011 -> 3
}
```

### 5. **Assignment Operators**

These are used to assign values to variables.

| Operator | Description         | Example  |
| -------- | ------------------- | -------- |
| `=`      | Assign              | `x = y`  |
| `+=`     | Add and assign      | `x += y` |
| `-=`     | Subtract and assign | `x -= y` |
| `*=`     | Multiply and assign | `x *= y` |
| `/=`     | Divide and assign   | `x /= y` |
| `%=`     | Modulus and assign  | `x %= y` |

**Example**:

```go
package main

import "fmt"

func main() {
    x := 10

    x += 5  // x = x + 5
    fmt.Println("x += 5:", x)  // 15

    x -= 3  // x = x - 3
    fmt.Println("x -= 3:", x)  // 12

    x *= 2  // x = x * 2
    fmt.Println("x *= 2:", x)  // 24

    x /= 4  // x = x / 4
    fmt.Println("x /= 4:", x)  // 6

    x %= 3  // x = x % 3
    fmt.Println("x %= 3:", x)  // 0
}
```

### 6. **Other Operators**

#### **Address & Dereference Operators**

These are used to work with pointers.

- **`&`**: Returns the address of a variable.
- **`*`**: Dereferences a pointer (accesses the value at the address).

**Example**:

```go
package main

import "fmt"

func main() {
    x := 10
    p := &x  // Pointer to x

    fmt.Println("Address of x:", p)  // Prints the address of x
    fmt.Println("Value at address:", *p)  // Dereferences the pointer (prints 10)
}
```

---

## **Operator Precedence in Go**

Operator precedence defines the order in which operators are evaluated in expressions. Operators with higher precedence are evaluated before operators with lower precedence.

### **Precedence Table (from highest to lowest)**

| Precedence Level | Operators                            |
| ---------------- | ------------------------------------ |
| 1 (highest)      | `*`, `/`, `%`, `<<`, `>>`, `&`, `&^` |
| 2                | `+`, `-`, `                          |
| 3                | `==`, `!=`, `<`, `<=`, `>`, `>=`     |
| 4                | `&&`                                 |
| 5 (lowest)       | `                                    |

### **Example of Precedence**:

```go
package main

import "fmt"

func main() {
    x := 10 + 2 * 3 // Multiplication has higher precedence
    fmt.Println(x)  // Output: 16
}
```

In this example, multiplication (`*`) has higher precedence than addition (`+`), so `2 * 3` is evaluated first.

---

## **Program Demonstrating All Data Types and Operators**

```go
package main

import "fmt"

func main() {
    // Integer Types
    var a int = 10
    var b int = 3

    fmt.Println("Integer Operations:")
    fmt.Println("Addition:", a + b)          // 13
    fmt.Println("Subtraction:", a - b)       // 7
    fmt.Println("Multiplication:", a * b)    // 30
    fmt.Println("Division:", a / b)          // 3
    fmt.Println("Modulus:", a % b)           // 1

    // Float Types
    var f1 float64 = 5.5
    var f2 float64 = 2.2

    fmt.Println("\nFloat Operations:")
    fmt.Println("Addition:", f1 + f2)        // 7.7
    fmt.Println("Subtraction:", f1 - f2)     // 3.3
    fmt.Println("Multiplication:", f1 * f2)  // 12.1
    fmt.Println("Division:", f1 / f2)        //

2.5

    // String Type
    str1 := "Go"
    str2 := "Lang"

    fmt.Println("\nString Operations:")
    fmt.Println("Concatenation:", str1 + str2) // "GoLang"

    // Boolean Type
    var bool1 bool = true
    var bool2 bool = false

    fmt.Println("\nBoolean Operations:")
    fmt.Println("AND:", bool1 && bool2)       // false
    fmt.Println("OR:", bool1 || bool2)        // true
    fmt.Println("NOT:", !bool1)               // false

    // Bitwise Operations
    fmt.Println("\nBitwise Operations:")
    fmt.Println("Bitwise AND:", a & b)       // 2 (binary: 1010 & 0011 = 0010)
    fmt.Println("Bitwise OR:", a | b)        // 11 (binary: 1010 | 0011 = 1011)
    fmt.Println("Bitwise XOR:", a ^ b)       // 9 (binary: 1010 ^ 0011 = 1001)

    // Pointer Operations
    fmt.Println("\nPointer Operations:")
    p := &a  // Pointer to a
    fmt.Println("Address of a:", p)   // Memory address of a
    fmt.Println("Dereferencing p:", *p)  // 10 (value of a)

    // Comparison
    fmt.Println("\nComparison Operations:")
    fmt.Println("a == b:", a == b)  // false
    fmt.Println("a > b:", a > b)    // true
}
```

### **Summary of Operators in Go**

- **Arithmetic Operators**: Perform basic math operations.
- **Comparison Operators**: Compare two values.
- **Logical Operators**: Combine boolean expressions.
- **Bitwise Operators**: Work on binary representations of numbers.
- **Assignment Operators**: Assign or modify values.
- **Pointer Operators**: Work with memory addresses.
