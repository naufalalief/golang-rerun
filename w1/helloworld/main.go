package main

import "fmt"

const (
	breakLine = "-------------------------"
)

func main() {
	fmt.Println("Print helloworld:")
	helloWorld() // Calls the helloWorld function
	fmt.Println(breakLine)

	fmt.Println("Print variables:")
	variables() // Calls the variables function
	fmt.Println(breakLine)

	fmt.Println("Print variables with data type:")
	variablesDataType() // Calls the variablesDataType function
	fmt.Println(breakLine)

	fmt.Println("Print variables with shorthand:")
	variablesShorthand() // Calls the variablesShorthand function
	fmt.Println(breakLine)

	fmt.Println("Print constant variable:")
	constantVariable() // Calls the constantVariable function
	fmt.Println(breakLine)

	fmt.Println("Print numeric non-decimal:")
	numericNonDecimal() // Calls the numericNonDecimal function
	fmt.Println(breakLine)

	fmt.Println("Print numeric decimal:")
	numericDecimal() // Calls the numericDecimal function
	fmt.Println(breakLine)

	fmt.Println("Print boolean:")
	boolean() // Calls the boolean function
	fmt.Println(breakLine)

	fmt.Println("Print string:")
	stringDataType() // Calls the string function
	fmt.Println(breakLine)

	fmt.Println("Print null and zero values:")
	nullZeroValue() // Calls the nullZeroValue function
	fmt.Println(breakLine)

	fmt.Println("Print conversion data type:")
	conversion() // Calls the conversionDataType function
	fmt.Println(breakLine)

	fmt.Println(("Print specific index of string:"))
	findIndex() // Calls the findIndex function
	fmt.Println(breakLine)

	fmt.Println("Print replace string:")
	replaceString() // Calls the replaceString function
	fmt.Println(breakLine)

	fmt.Println("Print repeat string:")
	repeatString() // Calls the repeatString function
	fmt.Println(breakLine)

	fmt.Println("Print lower string:")
	lowerString() // Calls the lowerString function
	fmt.Println(breakLine)

	fmt.Println("Print upper string:")
	upperString() // Calls the upperString function
	fmt.Println(breakLine)

	fmt.Println("Print length of string:")
	lenString() // Calls the lenString function
	fmt.Println(breakLine)

}
