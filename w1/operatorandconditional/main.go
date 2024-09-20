package main

import "fmt"

const breakLine = "-------------------------"

func main() {
	fmt.Println("Print math sum:")
	mathSum() // Calls the mathSum function
	fmt.Println(breakLine)

	fmt.Println("Print math subtract:")
	mathSubtract() // Calls the mathSubtract function
	fmt.Println(breakLine)

	fmt.Println("Print math multiply:")
	mathMultiply() // Calls the mathMultiply function
	fmt.Println(breakLine)

	fmt.Println("Print math divide:")
	mathDivide() // Calls the mathDivide function
	fmt.Println(breakLine)

	fmt.Println("Print math modulo:")
	mathModulo() // Calls the mathModulo function
	fmt.Println(breakLine)

	fmt.Println("Print math sum augmented:")
	mathSumAugmented() // Calls the mathSumAugmented function
	fmt.Println(breakLine)

	fmt.Println("Print math subtract augmented:")
	mathSubtractAugmented() // Calls the mathSubtractAugmented function
	fmt.Println(breakLine)

	fmt.Println("Print math multiply augmented:")
	mathMultiplyAugmented() // Calls the mathMultiplyAugmented function
	fmt.Println(breakLine)

	fmt.Println("Print math divide augmented:")
	mathDivideAugmented() // Calls the mathDivideAugmented function
	fmt.Println(breakLine)

	fmt.Println("Print math modulo augmented:")
	mathModuloAugmented() // Calls the mathModuloAugmented function
	fmt.Println(breakLine)
}
