package main

import "fmt"

// function with data types numeric non-decimal
// the data types for non-decimal numbers are uint8, uint16, uint32, uint64, uint, byte, int8, int16, int32, int64, int, and rune
// which represent signed and unsigned integers of various sizes
// the range of values that can be stored in these data types depends on the number of bits used to represent the value
// for example uint8 can store values from 0 to 255, while int8 can store values from -128 to 127
func numericNonDecimal() {
	var positiveNumber uint8 = 100  // Declares a variable named positiveNumber of type uint8 and assigns the value 100
	var negativeNumber = -123456789 // Declares a variable named negativeNumber and lets Go infer the data type based on the value

	fmt.Printf("Positive Number: %d\n", positiveNumber) // Prints the value of the positiveNumber variable to the console
	fmt.Printf("Negative Number: %d\n", negativeNumber) // Prints the value of the negativeNumber variable to the console
}
