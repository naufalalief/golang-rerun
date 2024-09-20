package main

import "fmt"

func arr() {
	// Define an array 'names' with 4 string elements
	var names = [4]string{"John", "Doe", "Frank", "Jack"}

	// Print each element of the 'names' array
	fmt.Println(names[0], names[1], names[2], names[3])

	// Define an array 'fruits' with 4 string elements
	var fruits = [4]string{
		"apple",
		"grape",
		"banana",
		"melon",
	}

	// Print the entire 'fruits' array
	fmt.Println(fruits)
}

// Array without number of elements
func arrWithoutNum() {
	// Define an array 'names' without specifying the number of elements
	var names = [...]string{"John", "Doe", "Frank", "Jack"}

	// Print each element of the 'names' array
	fmt.Println(names[0], names[1], names[2], names[3])
}

// Multi-dimensional array
func multiDimArr() {
	// Define a multi-dimensional array 'matrix' with 2 elements
	var matrix = [2][2]int{
		{1, 2},
		{3, 4},
	}

	// Print the entire 'matrix' array
	fmt.Println(matrix)
}

// Arrfunc is the entry point for this snippet
func Arrfunc() {
	fmt.Println("Array with number of elements:")
	arr()

	fmt.Println("Array without number of elements:")
	arrWithoutNum()

	fmt.Println("Multi-dimensional array:")
	multiDimArr()
}
