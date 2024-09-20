package main

import "fmt"

// andOperator compares two boolean values to see if they are both true.
// The result is true if both values are true, otherwise it is false.
func andOperator() {
	var a = true
	var b = false
	var result = a && b

	fmt.Println(result)
}

// orOperator compares two boolean values to see if at least one of them is true.
// The result is true if at least one value is true, otherwise it is false.
func orOperator() {
	var a = true
	var b = false
	var result = a || b

	fmt.Println(result)
}

// notOperator negates a boolean value.
// The result is true if the value is false, and false if the value is true.
func notOperator() {
	var a = true
	var result = !a

	fmt.Println(result)
}

// logicOperator is the entry point for this snippet.
func logicOperator() {
	fmt.Println("andOperator:")
	andOperator()

	fmt.Println("orOperator:")
	orOperator()

	fmt.Println("notOperator:")
	notOperator()
}
