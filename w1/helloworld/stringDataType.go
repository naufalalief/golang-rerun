package main

import "strconv"

// function that will return string values

func stringDataType() {
	var message = "Hi, My name is John Doe" // Declares a variable named message with the value "Hi, My name is John Doe"
	println(message)                        // Prints the value of the message variable to the console

	var name = "Jane Doe" // Declares a variable named name with the value "Jane Doe"
	var age = 25          // Declares a variable named age with the value 25
	var greeting = "Hi, My name is " + name + " and I am " + strconv.Itoa(age) + " years old"

	println(greeting) // Prints the value of the greeting variable to the console
}
