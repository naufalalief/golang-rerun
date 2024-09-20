package main

import "fmt"

// mathSum adds two numbers together.
// The result is the sum of the two numbers.
func mathSum() {
	var a = 5
	var b = 3
	var result = a + b

	fmt.Println(result)
}

// mathSubtract subtracts one number from another.
// The result is the difference between the two numbers.
func mathSubtract() {
	var a = 5
	var b = 3
	var result = a - b

	fmt.Println(result)
}

// mathMultiply multiplies two numbers together.
// The result is the product of the two numbers.
func mathMultiply() {
	var a = 5
	var b = 3
	var result = a * b

	fmt.Println(result)
}

// mathDivide divides one number by another.
// The result is the quotient of the two numbers.
func mathDivide() {
	var a = 5
	var b = 3
	var result = a / b

	fmt.Println(result)
}

// mathModulo divides one number by another.
// The result is the remainder of the two numbers.
func mathModulo() {
	var a = 5
	var b = 3
	var result = a % b

	fmt.Println(result)
}

// mathSumAugmented adds two numbers together using the augmented assignment operator.
// The result is the sum of the two numbers.
func mathSumAugmented() {
	var a = 5
	var b = 3
	a += b

	fmt.Println(a)
}

// mathSubtractAugmented subtracts one number from another using the augmented assignment operator.
// The result is the difference between the two numbers.
func mathSubtractAugmented() {
	var a = 5
	var b = 3
	a -= b

	fmt.Println(a)
}

// mathMultiplyAugmented multiplies two numbers together using the augmented assignment operator.
// The result is the product of the two numbers.
func mathMultiplyAugmented() {
	var a = 5
	var b = 3
	a *= b

	fmt.Println(a)
}

// mathDivideAugmented divides one number by another using the augmented assignment operator.
// The result is the quotient of the two numbers.
func mathDivideAugmented() {
	var a = 5
	var b = 3
	a /= b

	fmt.Println(a)
}

// mathModuloAugmented divides one number by another using the augmented assignment operator.
// The result is the remainder of the two numbers.
func mathModuloAugmented() {
	var a = 5
	var b = 3
	a %= b

	fmt.Println(a)
}
