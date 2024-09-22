package main

import "fmt"

func slices() {
	// Define a slice 'names' with 4 string elements
	var names = []string{"John", "Doe", "Frank", "Jack"}

	// Print each element of the 'names' slice
	fmt.Println(names[0], names[1], names[2], names[3])

	// Define a slice 'fruits' with 4 string elements
	var fruits = []string{
		"apple",
		"grape",
		"banana",
		"melon",
	}

	// Print the entire 'fruits' slice
	fmt.Println(fruits)
}

func diffWithArr() {
	// Array dengan ukuran tetap menggunakan notasi [...]
	var arr = [...]string{"John", "Doe", "Frank", "Jack"}
	fmt.Println("Array with [...] notation:", arr)

	// arr = append(arr, "Alice") // Error: cannot use append(arr, "Alice") (value of type []string) as [5]string value in assignment

	// Slice dengan ukuran dinamis
	var slice = []string{"John", "Doe", "Frank", "Jack"}
	fmt.Println("Slice:", slice)

	// Menambahkan elemen ke slice
	slice = append(slice, "Alice")
	fmt.Println("Updated Slice:", slice)
}

// newSlice is a function that creates a new slice from an existing array
func newSlice() {
	// Define an array 'arr' with 4 string elements
	arr := [4]string{"John", "Doe", "Frank", "Jack"}

	// Create a new slice 'slice' from the 'arr' array
	slice := arr[0:2]

	// Print the entire 'slice' slice
	fmt.Println(slice)
}

// sliceLen is a function that demonstrates the use of the len() function with slices
func sliceLen() {
	// Define a slice 'numbers' with 4 integer elements
	numbers := []int{1, 2, 3, 4}

	// Print the length of the 'numbers' slice
	fmt.Println("Length of slice:", len(numbers))
}

// sliceCap is a function that demonstrates the use of the cap() function with slices
func sliceCap() {
	// Define a slice 'numbers' with 4 integer elements
	numbers := []int{1, 2, 3, 4}

	aNumbers := numbers[0:3]

	// Print the capacity of the 'aNumbers' slice
	fmt.Println("Capacity of aNumbers:", cap(aNumbers))
	fmt.Println("Length of numbers:", len(aNumbers))

	// Print the capacity of the 'numbers' slice
	fmt.Println("Capacity of slice:", cap(numbers))
}

// sliceAppend is a function that demonstrates the use of the append() function with slices
func sliceAppend() {
	// Define a array 'numbers' with 4 integer elements
	numbers := []int{1, 2, 3, 4}
	// Define a new slice 'newNumbers' from the 'numbers' array
	newNumbers := numbers[0:3]

	fmt.Println("Array Numbers:", numbers)
	fmt.Println("Slice newNumbers:", newNumbers)

	// Append a new element to the 'newNumbers' slice
	newNumbers = append(newNumbers, 5)

	// Print the updated 'newNumbers' slice
	fmt.Println("Updated Slice newNumbers:", newNumbers)

}

// Slicefunc is the entry point for this snippet
func Slicefunc() {
	fmt.Println("Slice:")
	slices()

	fmt.Println("Difference between array and slice:")
	diffWithArr()

	fmt.Println("Create a new slice from an existing array:")
	newSlice()

	fmt.Println("Use of len() function with slices:")
	sliceLen()

	fmt.Println("Use of cap() function with slices:")
	sliceCap()

	fmt.Println("Use of append() function with slices:")
	sliceAppend()

}
