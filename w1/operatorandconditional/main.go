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
}
