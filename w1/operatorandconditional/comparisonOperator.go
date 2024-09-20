package main

import "fmt"

// comparisonEqual compares two numbers to see if they are equal.
// The result is true if the two numbers are equal, false otherwise.
func comparisonEqual() {
	var a = 5
	var b = 3
	var result = a == b

	fmt.Println(result)
}

// comparisonNotEqual compares two numbers to see if they are not equal.
// The result is true if the two numbers are not equal, false otherwise.
func comparisonNotEqual() {
	var a = 5
	var b = 3
	var result = a != b

	fmt.Println(result)
}

// comparisonGreaterThan compares two numbers to see if the first number is greater than the second number.
// The result is true if the first number is greater than the second number, false otherwise.
func comparisonGreaterThan() {
	var a = 5
	var b = 3
	var result = a > b

	fmt.Println(result)
}

// comparisonLessThan compares two numbers to see if the first number is less than the second number.
// The result is true if the first number is less than the second number, false otherwise.
func comparisonLessThan() {
	var a = 5
	var b = 3
	var result = a < b

	fmt.Println(result)
}

// comparisonGreaterThanOrEqual compares two numbers to see if the first number is greater than or equal to the second number.
// The result is true if the first number is greater than or equal to the second number, false otherwise.
func comparisonGreaterThanOrEqual() {
	var a = 5
	var b = 3
	var result = a >= b

	fmt.Println(result)
}

// comparisonLessThanOrEqual compares two numbers to see if the first number is less than or equal to the second number.
// The result is true if the first number is less than or equal to the second number, false otherwise.
func comparisonLessThanOrEqual() {
	var a = 5
	var b = 3
	var result = a <= b

	fmt.Println(result)
}

// comparisonOperator is the entry point for this snippet.
func comparisonOperator() {
	fmt.Println("Equal:")
	comparisonEqual()

	fmt.Println("Not Equal:")
	comparisonNotEqual()

	fmt.Println("Greater Than:")
	comparisonGreaterThan()

	fmt.Println("Less Than:")
	comparisonLessThan()

	fmt.Println("Greater Than or Equal:")
	comparisonGreaterThanOrEqual()

	fmt.Println("Less Than or Equal:")
	comparisonLessThanOrEqual()
}
