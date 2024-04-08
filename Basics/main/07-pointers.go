package main

import "fmt"

// Work with pointers
func pointers() {
	fmt.Println("Pointers")

	// Assign free memory to a pointer and assign a value
	var p *int32 = new(int32)
	fmt.Printf("Pointer P points to %v address that has the value %v\n", p, *p)
	*p = 10
	fmt.Printf("Pointer P points to %v address that has the value %v\n", p, *p)
	
	// Modify a variable value through a pointer
	var a int32 = 20
	p = &a
	fmt.Printf("The value of A is %v, and its address is %v\n", a, &a)
	*p = 15
	fmt.Printf("The value of A is %v, and its address is %v\n", a, &a)
	fmt.Println()

	// Pass a pointer to a function
	fmt.Println("Pass a pointer to a function")
	var array = [5]float64{1, 2, 3, 4, 5}
	fmt.Printf("The initial value of 'array' is: %v and its memory location is: %p\n", array, &array)
	fmt.Println("--Without Pointers:")
	var result = squareArray(array)
	fmt.Printf("The result is: %v\n", result)
	fmt.Println("--With Pointers:")
	result = squareArrayPointer(&array)
	fmt.Printf("The result is: %v\n", result)
	fmt.Println()
}

// Square every value of a given array
func squareArray(array [5]float64) [5]float64 {
	fmt.Printf("The value of 'array' inside the function is: %v and its memory location is: %p\n", array, &array)
	for i := range array {
		array[i] = array[i] * array[i]
	}
	return array
}

// Square every value of a given array, but using a pointer
func squareArrayPointer(array *[5]float64) [5]float64 {
	fmt.Printf("The value of 'array' inside the function is: %v and its memory location is: %p\n", *array, array)
	for i := range array {
		array[i] = array[i] * array[i]
	}
	return *array
}
