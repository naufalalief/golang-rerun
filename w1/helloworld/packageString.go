package main

import (
	"fmt"
	"strings"
)

// function that will find the index of a string with string.Index
// string.Index returns the index of the first instance of a substring in a string
// if the substring is not found, it returns -1
func findIndex() {
	var message = "Hello, World!" // Declares a variable named message with the value "Hello, World!"
	var index = "World"           // Declares a variable named index with the value "World"

	var result = strings.Index(message, index) // Finds the index of the index variable in the message variable
	fmt.Println(result)                        // Prints the result to the console
}
