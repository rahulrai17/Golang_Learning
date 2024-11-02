package main

import (
	"errors"  // Import the errors package for creating error values
	"fmt"     // Import the fmt package for formatted I/O
	"strconv" // Import strconv for converting strings to integers
)

// Function to divide two numbers with error handling for division by zero
func divide(x, y float64) (float64, error) {
	if y == 0 {
		// Return an error value if division by zero is attempted
		return 0, errors.New("cannot divide by zero")
	}
	// If no error, return the division result and nil (no error)
	return x / y, nil
}

// Function to convert a string to an integer with error handling
func stringToInt(s string) (int, error) {
	// Convert the string to an integer
	value, err := strconv.Atoi(s)
	if err != nil {
		// Return 0 and an error if the string cannot be converted to an integer
		return 0, errors.New("invalid input: cannot convert to integer")
	}
	// If conversion succeeds, return the integer and nil (no error)
	return value, nil
}

func main() {
	// Example 1: Division with error checking
	numerator := 10.0
	denominator := 0.0 // This will trigger a division by zero error
	result, err := divide(numerator, denominator)
	if err != nil {
		// Handle the division error by printing the error message
		fmt.Println("Division Error:", err)
	} else {
		// If no error, print the result
		fmt.Println("Result of division:", result)
	}

	// Example 2: Valid division
	numerator = 10.0
	denominator = 2.0 // Valid input
	result, err = divide(numerator, denominator)
	if err != nil {
		// Handle the error
		fmt.Println("Division Error:", err)
	} else {
		// Print the result if no error
		fmt.Println("Result of division:", result)
	}

	// Example 3: String to integer conversion with error handling
	strNumber := "123abc" // This will cause a conversion error
	number, err := stringToInt(strNumber)
	if err != nil {
		// Handle the conversion error
		fmt.Println("String Conversion Error:", err)
	} else {
		// Print the converted number if no error
		fmt.Println("Converted number:", number)
	}

	// Example 4: Successful string to integer conversion
	validStr := "42"
	validNumber, err := stringToInt(validStr)
	if err != nil {
		// Handle the error if any
		fmt.Println("String Conversion Error:", err)
	} else {
		// Print the valid converted number
		fmt.Println("Successfully converted number:", validNumber)
	}
}
