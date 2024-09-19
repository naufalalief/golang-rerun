package main

import "fmt"

// function that will return the string "Hello, World!" and "Goodbye, World!" using a constant variable
func constantVariable() {
	const text = "Hello, World!" // Declares a constant variable named text with the value "Hello, World!"
	fmt.Println(text)            // Prints the value of the text constant variable to the console

	// text = "Goodbye, World!" // This will cause an error because constants cannot be reassigned
}
