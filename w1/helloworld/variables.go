package main

import (
	"fmt"
)

// Define constants for the repeated strings
const (
	helloWorldText   = "Hello, World!"
	goodbyeWorldText = "Goodbye, World!"
)

// function that will return the string "Hello, World!" and "Goodbye, World!" using a variable
func variables() {
	var text = helloWorldText // Declares a variable named text with the value "Hello, World!"
	fmt.Println(text)         // Prints the value of the text variable to the console

	text = goodbyeWorldText // Assigns a new value to the text variable
	fmt.Println(text)       // Prints the new value of the text variable to the console
}

// function that will return the string "Hello, World!" and "Goodbye, World!" using a variable and data type
func variablesDataType() {
	var text string
	text = helloWorldText // Declares a variable named text of type string and assigns the value "Hello, World!"
	fmt.Println(text)     // Prints the value of the text variable to the console

	text = goodbyeWorldText // Assigns a new value to the text variable
	fmt.Println(text)       // Prints the new value of the text variable to the console
}

// function that will return the string "Hello, World!" and "Goodbye, World!" using a variable and data type with shorthand
func variablesShorthand() {
	text := helloWorldText // Declares a variable named text of type string and assigns the value "Hello, World!"
	fmt.Println(text)      // Prints the value of the text variable to the console

	text = goodbyeWorldText // Assigns a new value to the text variable
	fmt.Println(text)       // Prints the new value of the text variable to the console
}
