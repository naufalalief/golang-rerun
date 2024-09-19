package main

import "fmt"

// function that will convert data types using casting
func conversion() {
	var x float64 = float64(24) // Converts the integer value 24 to a float64 data type
	fmt.Println(x)              // Prints the value of the x variable to the console

	var y int32 = int32(31.00) // Converts the float64 value 31.00 to an int32 data type
	fmt.Println(y)             // Prints the value of the y variable to the console
}
