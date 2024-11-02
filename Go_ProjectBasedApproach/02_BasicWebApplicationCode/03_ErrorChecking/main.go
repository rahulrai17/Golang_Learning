package main

import (
	"fmt"
	"strconv" // Used for string to number conversion
)

// Function to divide two numbers with basic error handling (division by zero)
func divide(x, y float64) (float64, string) {
	if y == 0 {
		// Return an error message if dividing by zero
		return 0, "Error: Cannot divide by zero"
	}
	// If no error, return the result of division and an empty error string (no error)
	return x / y, ""
}

// Function to convert a string to an integer with basic error handling
func stringToInt(s string) (int, string) {
	// Convert the string to an integer
	value, err := strconv.Atoi(s)
	if err != nil {
		// Return a default value (0) and an error message if conversion fails
		return 0, "Error: Invalid input, cannot convert to integer"
	}
	// If conversion succeeds, return the value and an empty error string
	return value, ""
}

func main() {
	// Example 1: Division with basic error checking
	numerator := 10.0
	denominator := 0.0 // This will trigger a division by zero error
	result, err := divide(numerator, denominator)
	if err != "" {
		// Handle the error
		fmt.Println(err)
	} else {
		// If no error, print the result
		fmt.Println("Result of division:", result)
	}

	// Example 2: Valid division
	numerator = 10.0
	denominator = 2.0 // Valid input
	result, err = divide(numerator, denominator)
	if err != "" {
		// Handle the error
		fmt.Println(err)
	} else {
		// Print the result if no error
		fmt.Println("Result of division:", result)
	}

	// Example 3: String to integer conversion with basic error handling
	strNumber := "123abc" // This will cause a conversion error
	number, err := stringToInt(strNumber)
	if err != "" {
		// Handle the conversion error
		fmt.Println(err)
	} else {
		// Print the converted number if no error
		fmt.Println("Converted number:", number)
	}

	// Example 4: Successful string to integer conversion
	validStr := "42"
	validNumber, err := stringToInt(validStr)
	if err != "" {
		fmt.Println(err)
	} else {
		// Print the valid converted number
		fmt.Println("Successfully converted number:", validNumber)
	}
}
