package main

import "fmt"

const breakLine = "-------------------------"

func main() {
	fmt.Println("Print math operator:")
	mathOperator() // Calls the mathOperator function
	fmt.Println(breakLine)

	fmt.Println("Print comparison operator:")
	comparisonOperator() // Calls the comparisonOperator function
	fmt.Println(breakLine)

	fmt.Println("Print logic operator:")
	logicOperator() // Calls the logicOperator function
	fmt.Println(breakLine)

	fmt.Println("Print conditional operator:")
	conditionals() // Calls the conditionals function
	fmt.Println(breakLine)

	fmt.Println("Print arrays:")
	Arrfunc() // Calls the Arrfunc function
	fmt.Println(breakLine)

	fmt.Println("Print slices:")
	Slicefunc() // Calls the Slicefunc function
	fmt.Println(breakLine)
}
