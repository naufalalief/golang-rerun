package main

import (
	"fmt"
	"strconv"
)

// stringToInt converts a string to an integer using strconv.Atoi.
// strconv.Atoi converts a string to an integer and returns the result.
func stringToInt() {
	var str = "42"                 // Declares a variable named str with the value "42"
	var num, _ = strconv.Atoi(str) // Converts the str variable to an integer and assigns it to the num variable
	fmt.Println(num)               // Prints the num variable to the console
}

// intToString converts an integer to a string using strconv.Itoa.
// strconv.Itoa converts an integer to a string and returns the result.
func intToString() {
	var num = 42                // Declares a variable named num with the value 42
	var str = strconv.Itoa(num) // Converts the num variable to a string and assigns it to the str variable
	fmt.Println(str)            // Prints the str variable to the console
}

// parseInteger parses an integer from a string using strconv.ParseInt.
// strconv.ParseInt parses an integer from a string and returns the result.
func parseInteger() {
	var str = "42"                             // Declares a variable named str with the value "42"
	var num, _ = strconv.ParseInt(str, 10, 64) // Parses an integer from the str variable and assigns it to the num variable
	fmt.Println(num)                           // Prints the num variable to the console
}

// formatInteger formats an integer as a string using strconv.FormatInt.
// strconv.FormatInt formats an integer as a string and returns the result.
func formatInteger() {
	var num int64 = 42                   // Declares a variable named num with the value 42
	var str = strconv.FormatInt(num, 10) // Formats the num variable as a string and assigns it to the str variable
	fmt.Println(str)                     // Prints the str variable to the console
}

// parseFloat parses a floating-point number from a string using strconv.ParseFloat.
// strconv.ParseFloat parses a floating-point number from a string and returns the result.
func parseFloat() {
	var str = "3.14"                         // Declares a variable named str with the value "3.14"
	var num, _ = strconv.ParseFloat(str, 64) // Parses a floating-point number from the str variable and assigns it to the num variable
	fmt.Println(num)                         // Prints the num variable to the console
}

// formatFloat formats a floating-point number as a string using strconv.FormatFloat.
// strconv.FormatFloat formats a floating-point number as a string and returns the result.
func formatFloat() {
	var num = 3.14                                  // Declares a variable named num with the value 3.14
	var str = strconv.FormatFloat(num, 'f', -1, 64) // Formats the num variable as a string and assigns it to the str variable
	fmt.Println(str)                                // Prints the str variable to the console
}

// parseBool parses a boolean from a string using strconv.ParseBool.
// strconv.ParseBool parses a boolean from a string and returns the result.
func parseBool() {
	var str = "true"                  // Declares a variable named str with the value "true"
	var b, _ = strconv.ParseBool(str) // Parses a boolean from the str variable and assigns it to the b variable
	fmt.Println(b)                    // Prints the b variable to the console
}

// formatBool formats a boolean as a string using strconv.FormatBool.
// strconv.FormatBool formats a boolean as a string and returns the result.
func formatBool() {
	var b = true                    // Declares a variable named b with the value true
	var str = strconv.FormatBool(b) // Formats the b variable as a string and assigns it to the str variable
	fmt.Println(str)                // Prints the str variable to the console
}

// packageStrconv is the entry point for this snippet.
func packageStrconv() {
	fmt.Println("Convert string to integer:")
	stringToInt() // Calls the stringToInt function

	fmt.Println("Convert integer to string:")
	intToString() // Calls the intToString function

	fmt.Println("Parse integer from string:")
	parseInteger() // Calls the parseInteger function

	fmt.Println("Format integer as string:")
	formatInteger() // Calls the formatInteger function

	fmt.Println("Parse float from string:")
	parseFloat() // Calls the parseFloat function

	fmt.Println("Format float as string:")
	formatFloat() // Calls the formatFloat function

	fmt.Println("Parse bool from string:")
	parseBool() // Calls the parseBool function

	fmt.Println("Format bool as string:")
	formatBool() // Calls the formatBool function

}
