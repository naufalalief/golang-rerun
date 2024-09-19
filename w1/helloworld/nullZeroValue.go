package main

// function that will return null and zero values
// zero values are default values for variables
// null values are values that are not assigned to a variable
// for example, a variable that is declared but not assigned a value
// like var x int
// x is a null value because it has not been assigned a value
func nullZeroValue() {
	var x int  // Declares a variable named x of type int with the zero value of 0
	println(x) // Prints the value of the x variable to the console

	var y string // Declares a variable named y of type string with the zero value of ""
	println(y)   // Prints the value of the y variable to the console

	var z bool // Declares a variable named z of type bool with the zero value of false
	println(z) // Prints the value of the z variable to the console
}
