package main

import "fmt"

// function with data types numeric decimal
// the data types for decimal numbers are float32 and float64, which represent single-precision and double-precision floating-point numbers
// respectively
// floating-point numbers can represent a wider range of values than integers, but they are subject to rounding errors due to the limited precision of the data type
func numericDecimal() {
	var pi float32 = 3.14159265358979323846264338327950288419716939937510582097494459
	var e = 2.71828182845904523536028747135266249775724709369995957496696763

	fmt.Printf("Pi: %f\n", pi)            // Prints the value of the pi variable to the console
	fmt.Printf("Euler's Number: %f\n", e) // Prints the value of the e variable to the console
}
