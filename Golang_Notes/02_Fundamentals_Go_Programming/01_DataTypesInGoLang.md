## Data Types in Golang

## **1. Boolean (`bool`)**

- **Definition:**
  - Represents a binary value: `true` or `false`.
- **Usage:**
  - Typically used in conditional logic (if-else statements, loops) and flags.

**Code Example:**

```go
var isActive bool = true

if isActive {
    fmt.Println("The system is active.")
} else {
    fmt.Println("The system is inactive.")
}
```

**Coding Question Example:**

- **Question:** Implement a function that returns `true` if a number is even, and `false` otherwise.

**Solution:**

```go
package main

import (
	"fmt"
)

func main() {
	var value int = 4;
	fmt.Println("Number is Even : ", isEven(value))
}

func isEven(n int) bool{
	return n%2 == 0
}

```

### **Best Practices:**

- Use boolean expressions directly for conditionals (avoid `if condition == true`, just use `if condition`).
- Make variable names expressive, like `isValid`, `isEnabled`, etc., for clarity.

---

## **2. Integer Types (`int`, `int8`, `int16`, `int32`, `int64`, `uint`)**

- **Definition:**
  - Signed (`int`) and unsigned (`uint`) integer types of varying sizes.
  - `int` is platform-dependent (32 or 64 bits), while `int8`, `int16`, `int32`, and `int64` have fixed bit sizes.

**Code Example:**

```go
var a int = 10
var b int64 = 9223372036854775807
var c uint = 100

fmt.Printf("a: %d, b: %d, c: %d\n", a, b, c)
```

**Coding Question Example:**

- **Question:** Write a function to compute the factorial of a non-negative integer using recursion.

**Solution:**

```go
func factorial(n int) int {
    if n == 0 {
        return 1
    }
    return n * factorial(n-1)
}
```

### **Best Practices:**

- Use `int` for general purposes unless a specific bit size or unsigned type is necessary.
- Be aware of integer overflow. Example: `int8` can only hold values between `-128` and `127`.
- Use `uint` cautiously, as arithmetic with negative numbers can cause unexpected results.

---

## **3. Floating Point Types (`float32`, `float64`)**

- **Definition:**
  - Represents real numbers with fractional components.
  - `float32` is single precision, and `float64` is double precision.

**Code Example:**

```go
var pi float64 = 3.14159
var smallNumber float32 = 0.000001

fmt.Printf("Pi: %f, Small Number: %f\n", pi, smallNumber)
```

**Coding Question Example:**

- **Question:** Write a function that calculates the area of a circle given its radius.

**Solution:**

```go
func circleArea(radius float64) float64 {
    return 3.14159 * radius * radius
}
```

### **Best Practices:**

- Use `float64` for better precision, especially in mathematical computations.
- Avoid direct equality checks with floating-point numbers due to precision issues; instead, check if the numbers are "close enough."

Example:

```go
func almostEqual(a, b, tolerance float64) bool {
    return math.Abs(a-b) < tolerance
}
```

---

## **4. String**

- **Definition:**
  - A sequence of bytes representing text (UTF-8 encoded by default).
  - Strings are immutable, meaning their value cannot be changed once set.

**Code Example:**

```go
var name string = "Gopher"
fmt.Println(name)
```

**Coding Question Example:**

- **Question:** Write a function to check if two strings are anagrams.

**Solution:**

```go
func areAnagrams(s1, s2 string) bool {
    if len(s1) != len(s2) {
        return false
    }

    count := make(map[rune]int)

    for _, char := range s1 {
        count[char]++
    }
    for _, char := range s2 {
        count[char]--
        if count[char] < 0 {
            return false
        }
    }
    return true
}
```

### **Best Practices:**

- Use string concatenation (`+`) carefully, as it creates new strings, which can be costly in loops.
- For performance-critical cases, consider using `strings.Builder` for efficient concatenation.

Example:

```go
var builder strings.Builder
builder.WriteString("Hello, ")
builder.WriteString("world!")
fmt.Println(builder.String())
```

---

## **5. Rune**

- **Definition:**
  - A `rune` is an alias for `int32` and represents a Unicode code point.

**Code Example:**

```go
var r rune = 'A'
fmt.Printf("Rune: %c, Unicode: %U\n", r, r)
```

**Coding Question Example:**

- **Question:** Write a function that reverses a string containing Unicode characters.

**Solution:**

```go
func reverseRunes(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}
```

### **Best Practices:**

- Use runes when working with Unicode characters to avoid errors with multi-byte characters (UTF-8).
- Avoid iterating over strings byte-by-byte if you need character-level processing.

---

## Composite Data Types

## **6. Array**

- **Definition:**
  - A fixed-size sequence of elements of the same type.

**Code Example:**

```go
var arr [5]int = [5]int{1, 2, 3, 4, 5}
fmt.Println(arr)
```

**Coding Question Example:**

- **Question:** Write a function to rotate an array left by `k` positions.

**Solution:**

```go
func rotateLeft(arr []int, k int) []int {
    n := len(arr)
    k = k % n
    return append(arr[k:], arr[:k]...)
}
```

### **Best Practices:**

- Use slices instead of arrays for most cases where dynamic resizing is required.
- Use arrays for fixed-size structures where the size is known at compile time.

---

## **7. Slice**

- **Definition:**
  - A dynamically-sized, flexible view into the elements of an array.

**Code Example:**

```go
slice := []int{1, 2, 3}
slice = append(slice, 4, 5)
fmt.Println(slice)
```

**Coding Question Example:**

- **Question:** Write a function that removes all duplicate values from a slice of integers.

**Solution:**

```go
func removeDuplicates(nums []int) []int {
    seen := make(map[int]bool)
    result := []int{}

    for _, num := range nums {
        if !seen[num] {
            result = append(result, num)
            seen[num] = true
        }
    }
    return result
}
```

### **Best Practices:**

- Be mindful of slice growth during `append()` operations, as this may involve reallocation of the underlying array.
- If performance is critical, consider pre-allocating the slice capacity with `make()`.

Example:

```go
slice := make([]int, 0, 10)  // Capacity of 10
```

---

## **8. Map**

- **Definition:**
  - A collection of key-value pairs. Keys must be unique, and keys can be of any type that supports equality comparisons.

**Code Example:**

```go
myMap := map[string]int{"Alice": 25, "Bob": 30}
fmt.Println(myMap)
```

**Coding Question Example:**

- **Question:** Write a function that counts the frequency of words in a given string.

**Solution:**

```go
func wordCount(s string) map[string]int {
    words := strings.Fields(s)
    countMap := make(map[string]int)

    for _, word := range words {
        countMap[word]++
    }
    return countMap
}
```

### **Best Practices:**

- Always check if a key exists in the map before accessing or modifying it (use the "comma ok" idiom).
- Use `make()` to initialize maps before adding elements.

---

## **9. Struct**

- **Definition:**
  - A composite data type that groups together fields (which can be of different types).

**Code Example:**

```go
type Person struct {
    Name string
    Age  int
}

person := Person{Name: "Alice", Age: 30}
fmt.Println(person)
```

**Coding Question Example:**

- **Question:** Implement a function that returns the name and age of the oldest person from a list of people.

**Solution:**

```go
type Person struct {
    Name string
    Age  int
}

func oldestPerson(people []Person) Person {
    var oldest Person
    for _, p := range people {
        if p.Age > oldest.Age {
            oldest = p
        }
    }
    return oldest
}
```

### **Best Practices:**

- Use structs to group related data logically.
- Take advantage of Go’s zero values; fields in a struct will automatically be set to their zero values if not initialized.

---

## Reference Types

## **10. Pointer**

- **Definition:**
  - A pointer stores the memory address of a variable.

**Code Example:**

```go
var x int = 10
var ptr *int = &x

fmt.Println("Value of x:", *ptr)
```

**Coding Question Example:**

- **Question:** Write a function that swaps the values of two integers using pointers.

**Solution:**

```go
func swap(a, b *int) {
    *a, *b = *b, *a
}
```

### **Best Practices:**

- Use pointers to avoid copying large structs or to allow functions to modify values.
- Be cautious of nil pointers; always check for `nil` before dereferencing.

---

## Advance Types

#### **a. Interfaces**

- **Developer's Perspective:**
  - Defines a set of methods a type must implement.
  - Essential for polymorphism and decoupling.
  - Example:
    ```go
    type Animal interface {
        Speak() string
    }
    ```
- **Interview Tip:**
  - Expect interview questions on the purpose of interfaces in Go’s type system.
  - You may be asked to implement polymorphic behavior or work with empty interfaces (`interface{}`).

#### **b. Channels (Concurrency)**

- **Developer's Perspective:**
  - Used for communication between goroutines.
  - Synchronizes concurrent operations safely.
  - Example:
    ```go
    ch := make(chan int)
    go func() { ch <- 42 }()
    value := <-ch
    ```
- **Interview Tip:**
  - Interviewers might test your concurrency skills with channel problems, especially in producer-consumer scenarios.

---

### Custom Data Types

#### **a. Type Aliases**

- **Developer's Perspective:**
  - Allows creating a new name for an existing type.
  - Example:
    ```go
    type Age int
    var age Age = 30
    ```
- **Interview Tip:**
  - Be ready to explain the difference between type aliases and type definitions.

#### **b. Type Definitions**

- **Developer's Perspective:**
  - A new type entirely distinct from the original, even though it might share the same underlying structure.
  - Example:
    ```go
    type MyInt int
    ```

---

### Best Practices and Interview Considerations

- **Zero Values:** Every variable in Go has a default zero value (e.g., `0` for integers, `""` for strings, `nil` for pointers, maps, slices, and channels). Interviewers may ask about these defaults, especially with respect to initialization patterns.
- **Pass by Value vs. Pass by Reference:** Go passes arguments by value, but you can use pointers to achieve reference-like behavior. Be ready to explain the difference and the benefits of each.

- **Memory Management:** While Go has garbage collection, you might be asked about how memory is allocated on the heap vs. stack, especially with larger structs or slices.

- **Concurrency with Goroutines and Channels:** Mastery over Go’s concurrency model is a huge plus for interviews. Expect questions on synchronization, race conditions, and deadlocks when using channels or `sync.Mutex`.

---
