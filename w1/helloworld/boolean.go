package main

import "fmt"

// function that will return boolean values
func boolean() {
	var isTrue bool     // Declares a variable named isTrue of type bool with the zero value of false
	fmt.Println(isTrue) // Prints the value of the isTrue variable to the console

	isTrue = true       // Assigns the value of true to the isTrue variable
	fmt.Println(isTrue) // Prints the value of the isTrue variable to the console
}
