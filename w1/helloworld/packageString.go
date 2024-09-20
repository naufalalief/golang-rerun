package main

import (
	"fmt"
	"strings"
)

// findIndex finds the index of a substring within a string using strings.Index.
// strings.Index returns the index of the first instance of a substring in a string.
// If the substring is not found, it returns -1.
func findIndex() {
	var message = "Hello, World!" // Declares a variable named message with the value "Hello, World!"
	var substring = "World"       // Declares a variable named substring with the value "World"

	var result = strings.Index(message, substring) // Finds the index of the substring variable in the message variable
	fmt.Println(result)                            // Prints the result to the console
}

// replaceString replaces instances of a substring within a string using strings.Replace.
// strings.Replace replaces all instances of a substring in a string with another substring.
// If the count parameter is -1, it replaces all instances.
// If the count parameter is greater than 0, it replaces that number of instances.
func replaceString() {
	var message = "ananas" // Declares a variable named message with the value "ananas"
	var find = "a"         // Declares a variable named find with the value "a"
	var replaceWith = "o"  // Declares a variable named replaceWith with the value "o"

	var result = strings.Replace(message, find, replaceWith, -1) // Replaces all instances of the find variable with the replaceWith variable in the message variable, resulting in "ononos"
	var result2 = strings.Replace(message, find, replaceWith, 1) // Replaces the first instance of the find variable with the replaceWith variable in the message variable, resulting in "onanas"
	var result3 = strings.Replace(message, find, replaceWith, 2) // Replaces the first two instances of the find variable with the replaceWith variable in the message variable, resulting in "ononas"

	fmt.Println(result, result2, result3) // Prints the results to the console
}
